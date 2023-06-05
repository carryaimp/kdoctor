// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package netreachhealthy

import (
	"context"
	"fmt"
	k8sObjManager "github.com/kdoctor-io/kdoctor/pkg/k8ObjManager"
	crd "github.com/kdoctor-io/kdoctor/pkg/k8s/apis/kdoctor.io/v1beta1"
	"github.com/kdoctor-io/kdoctor/pkg/pluginManager/tools"
	"github.com/kdoctor-io/kdoctor/pkg/types"
	"go.uber.org/zap"
	networkingv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
)

func (s *PluginNetReachHealthy) WebhookMutating(logger *zap.Logger, ctx context.Context, obj runtime.Object) error {
	req, ok := obj.(*crd.NetReachHealthy)
	if !ok {
		s := "failed to get NetReachHealthy obj"
		logger.Error(s)
		return apierrors.NewBadRequest(s)
	}

	if req.DeletionTimestamp != nil {
		return nil
	}

	if req.Spec.Target == nil {
		var agentV4Url, agentV6Url *k8sObjManager.ServiceAccessUrl
		var e error

		testIngress := false
		var agentIngress *networkingv1.Ingress
		agentIngress, e = k8sObjManager.GetK8sObjManager().GetIngress(ctx, types.ControllerConfig.Configmap.AgentIngressName, types.ControllerConfig.PodNamespace)
		if e != nil {
			logger.Sugar().Errorf("failed to get ingress , error=%v", e)
		}
		if agentIngress != nil && len(agentIngress.Status.LoadBalancer.Ingress) > 0 {
			testIngress = true
		}

		serviceAccessPortName := "http"
		testLoadBalancer := false
		if types.ControllerConfig.Configmap.EnableIPv4 {
			agentV4Url, e = k8sObjManager.GetK8sObjManager().GetServiceAccessUrl(ctx, types.ControllerConfig.Configmap.AgentSerivceIpv4Name, types.ControllerConfig.PodNamespace, serviceAccessPortName)
			if e != nil {
				logger.Sugar().Errorf("failed to get agent ipv4 service url , error=%v", e)
			}
			if len(agentV4Url.LoadBalancerUrl) > 0 {
				testLoadBalancer = true
			}
		}
		if types.ControllerConfig.Configmap.EnableIPv6 {
			agentV6Url, e = k8sObjManager.GetK8sObjManager().GetServiceAccessUrl(ctx, types.ControllerConfig.Configmap.AgentSerivceIpv4Name, types.ControllerConfig.PodNamespace, serviceAccessPortName)
			if e != nil {
				logger.Sugar().Errorf("failed to get agent ipv6 service url , error=%v", e)
			}
			if len(agentV6Url.LoadBalancerUrl) > 0 {
				testLoadBalancer = true
			}
		}

		enableIpv4 := types.ControllerConfig.Configmap.EnableIPv4
		enableIpv6 := types.ControllerConfig.Configmap.EnableIPv6
		m := &crd.NetReachHealthyTarget{
			Endpoint:        true,
			MultusInterface: false,
			ClusterIP:       true,
			NodePort:        true,
			LoadBalancer:    testLoadBalancer,
			Ingress:         testIngress,
			IPv6:            &enableIpv6,
			IPv4:            &enableIpv4,
		}
		req.Spec.Target = m
		logger.Sugar().Debugf("set default target for NetReachHealthy %v", req.Name)
	}

	if req.Spec.Schedule == nil {
		req.Spec.Schedule = tools.GetDefaultSchedule()
		logger.Sugar().Debugf("set default SchedulePlan for NetReachHealthy %v", req.Name)
	}

	if req.Spec.Request == nil {
		m := &crd.NetHttpRequest{
			DurationInSecond:      types.ControllerConfig.Configmap.NethttpDefaultRequestDurationInSecond,
			QPS:                   types.ControllerConfig.Configmap.NethttpDefaultRequestQps,
			PerRequestTimeoutInMS: types.ControllerConfig.Configmap.NethttpDefaultRequestPerRequestTimeoutInMS,
		}
		req.Spec.Request = m
		logger.Sugar().Debugf("set default Request for NetReachHealthy %v", req.Name)
	}

	if req.Spec.SuccessCondition == nil {
		req.Spec.SuccessCondition = tools.GetDefaultNetSuccessCondition()
		logger.Sugar().Debugf("set default SuccessCondition for NetReachHealthy %v", req.Name)
	}

	return nil
}

func (s *PluginNetReachHealthy) WebhookValidateCreate(logger *zap.Logger, ctx context.Context, obj runtime.Object) error {
	r, ok := obj.(*crd.NetReachHealthy)
	if !ok {
		s := "failed to get NetReachHealthy obj"
		logger.Error(s)
		return apierrors.NewBadRequest(s)
	}
	logger.Sugar().Debugf("NetReachHealthy: %+v", r)

	// validate Schedule
	if true {
		if err := tools.ValidataCrdSchedule(r.Spec.Schedule); err != nil {
			s := fmt.Sprintf("NetReachHealthy %v : %v", r.Name, err)
			logger.Error(s)
			return apierrors.NewBadRequest(s)
		}
	}

	// validate request
	if true {
		if r.Spec.Request.QPS >= types.ControllerConfig.Configmap.NethttpDefaultRequestMaxQps {
			s := fmt.Sprintf("NetReachHealthy %v requires qps %v bigger than maximum %v", r.Name, r.Spec.Request.QPS, types.ControllerConfig.Configmap.NethttpDefaultRequestMaxQps)
			logger.Error(s)
			return apierrors.NewBadRequest(s)
		}
		if r.Spec.Request.PerRequestTimeoutInMS > int(r.Spec.Schedule.RoundTimeoutMinute*60*1000) {
			s := fmt.Sprintf("NetReachHealthy %v requires PerRequestTimeoutInMS %v ms smaller than Schedule.RoundTimeoutMinute %vm ", r.Name, r.Spec.Request.PerRequestTimeoutInMS, r.Spec.Schedule.RoundTimeoutMinute)
			logger.Error(s)
			return apierrors.NewBadRequest(s)
		}
		if r.Spec.Request.DurationInSecond > int(r.Spec.Schedule.RoundTimeoutMinute*60) {
			s := fmt.Sprintf("NetReachHealthy %v requires request.DurationInSecond %vs smaller than Schedule.RoundTimeoutMinute %vm ", r.Name, r.Spec.Request.DurationInSecond, r.Spec.Schedule.RoundTimeoutMinute)
			logger.Error(s)
			return apierrors.NewBadRequest(s)
		}
	}

	// validate target
	if true {
		if r.Spec.Target != nil {
			// validate target
			if r.Spec.Target.IPv4 != nil && *(r.Spec.Target.IPv4) && !types.ControllerConfig.Configmap.EnableIPv4 {
				s := fmt.Sprintf("NetReachHealthy %v TestIPv4, but kdoctor ipv4 feature is disabled", r.Name)
				logger.Error(s)
				return apierrors.NewBadRequest(s)
			}
			if r.Spec.Target.IPv6 != nil && *(r.Spec.Target.IPv6) && !types.ControllerConfig.Configmap.EnableIPv6 {
				s := fmt.Sprintf("NetReachHealthy %v TestIPv6, but kdoctor ipv6 feature is disabled", r.Name)
				logger.Error(s)
				return apierrors.NewBadRequest(s)
			}
		}
	}

	// validate SuccessCondition
	if true {
		if r.Spec.SuccessCondition.SuccessRate == nil && r.Spec.SuccessCondition.MeanAccessDelayInMs == nil {
			s := fmt.Sprintf("NetReachHealthy %v, no SuccessCondition specified in the spec", r.Name)
			logger.Error(s)
			return apierrors.NewBadRequest(s)
		}
		if r.Spec.SuccessCondition.SuccessRate != nil && (*(r.Spec.SuccessCondition.SuccessRate) > 1) {
			s := fmt.Sprintf("NetReachHealthy %v, SuccessCondition.SuccessRate %v must not be bigger than 1", r.Name, *(r.Spec.SuccessCondition.SuccessRate))
			logger.Error(s)
			return apierrors.NewBadRequest(s)
		}
		if r.Spec.SuccessCondition.SuccessRate != nil && (*(r.Spec.SuccessCondition.SuccessRate) < 0) {
			s := fmt.Sprintf("NetReachHealthy %v, SuccessCondition.SuccessRate %v must not be smaller than 0 ", r.Name, *(r.Spec.SuccessCondition.SuccessRate))
			logger.Error(s)
			return apierrors.NewBadRequest(s)
		}
	}

	return nil
}

// this will not be called, it is not allowed to modify crd
func (s *PluginNetReachHealthy) WebhookValidateUpdate(logger *zap.Logger, ctx context.Context, oldObj, newObj runtime.Object) error {

	return nil
}

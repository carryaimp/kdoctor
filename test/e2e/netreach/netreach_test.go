// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package netreach_test

import (
	"github.com/kdoctor-io/kdoctor/pkg/k8s/apis/kdoctor.io/v1beta1"
	"github.com/kdoctor-io/kdoctor/pkg/pluginManager"
	"github.com/kdoctor-io/kdoctor/test/e2e/common"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spidernet-io/e2eframework/tools"
)

var _ = Describe("testing netReach ", Label("netReach"), func() {
	var termMin = int64(1)
	// 1000ms is not stable on GitHub ci, so increased to 3000ms
	var requestTimeout = 3000
	It("success testing netReach", Label("B00001", "C00004", "E00001"), func() {
		var e error
		successRate := float64(1)
		successMean := int64(1500)
		crontab := "0 1"
		netReachName := "netreach-" + tools.RandomName()

		netReach := new(v1beta1.NetReach)
		netReach.Name = netReachName

		// agentSpec
		agentSpec := new(v1beta1.AgentSpec)
		agentSpec.TerminationGracePeriodMinutes = &termMin
		netReach.Spec.AgentSpec = agentSpec

		// successCondition
		successCondition := new(v1beta1.NetSuccessCondition)
		successCondition.SuccessRate = &successRate
		successCondition.MeanAccessDelayInMs = &successMean
		netReach.Spec.SuccessCondition = successCondition
		enable := true
		disable := false
		// target
		target := new(v1beta1.NetReachTarget)
		if !common.TestIPv4 && common.TestIPv6 {
			target.Ingress = &disable
		} else {
			target.Ingress = &enable
		}
		target.LoadBalancer = &enable
		target.ClusterIP = &enable
		target.Endpoint = &enable
		target.NodePort = &enable
		target.MultusInterface = &disable
		target.IPv4 = &common.TestIPv4
		target.IPv6 = &common.TestIPv6
		netReach.Spec.Target = target

		// request
		request := new(v1beta1.NetHttpRequest)
		request.PerRequestTimeoutInMS = requestTimeout
		request.QPS = 10
		request.DurationInSecond = 10
		netReach.Spec.Request = request

		// Schedule
		Schedule := new(v1beta1.SchedulePlan)
		Schedule.Schedule = &crontab
		Schedule.RoundNumber = 1
		Schedule.RoundTimeoutMinute = 1
		netReach.Spec.Schedule = Schedule

		e = frame.CreateResource(netReach)
		Expect(e).NotTo(HaveOccurred(), "create netReach resource")

		e = common.CheckRuntime(frame, netReach, pluginManager.KindNameNetReach, 60)
		Expect(e).NotTo(HaveOccurred(), "check task runtime spec")

		e = common.WaitKdoctorTaskDone(frame, netReach, pluginManager.KindNameNetReach, 120)
		Expect(e).NotTo(HaveOccurred(), "wait netReach task finish")

		success, e := common.CompareResult(frame, netReachName, pluginManager.KindNameNetReach, []string{}, reportNum, netReach)
		Expect(e).NotTo(HaveOccurred(), "compare report and task")
		Expect(success).To(BeTrue(), "compare report and task result")

		e = common.CheckRuntimeDeadLine(frame, netReachName, pluginManager.KindNameNetReach, 120)
		Expect(e).NotTo(HaveOccurred(), "check task runtime resource delete")
	})

	It("Successfully testing using default daemonSet  as workload with Task NetReach", Label("E00013"), func() {
		var e error
		successRate := float64(1)
		successMean := int64(1500)
		crontab := "0 1"
		netReachName := "netreach-" + tools.RandomName()

		netReach := new(v1beta1.NetReach)
		netReach.Name = netReachName

		// successCondition
		successCondition := new(v1beta1.NetSuccessCondition)
		successCondition.SuccessRate = &successRate
		successCondition.MeanAccessDelayInMs = &successMean
		netReach.Spec.SuccessCondition = successCondition
		enable := true
		disable := false
		// target
		target := new(v1beta1.NetReachTarget)
		if !common.TestIPv4 && common.TestIPv6 {
			target.Ingress = &disable
		} else {
			target.Ingress = &enable
		}
		target.LoadBalancer = &enable
		target.ClusterIP = &enable
		target.Endpoint = &enable
		target.NodePort = &enable
		target.MultusInterface = &disable
		target.IPv4 = &common.TestIPv4
		target.IPv6 = &common.TestIPv6
		netReach.Spec.Target = target

		// request
		request := new(v1beta1.NetHttpRequest)
		request.PerRequestTimeoutInMS = requestTimeout
		request.QPS = 5
		request.DurationInSecond = 5
		netReach.Spec.Request = request

		// Schedule
		Schedule := new(v1beta1.SchedulePlan)
		Schedule.Schedule = &crontab
		Schedule.RoundNumber = 1
		Schedule.RoundTimeoutMinute = 1
		netReach.Spec.Schedule = Schedule

		e = frame.CreateResource(netReach)
		Expect(e).NotTo(HaveOccurred(), "create netReach resource")

		e = common.CheckRuntime(frame, netReach, pluginManager.KindNameNetReach, 60)
		Expect(e).NotTo(HaveOccurred(), "check task runtime spec")

		e = common.WaitKdoctorTaskDone(frame, netReach, pluginManager.KindNameNetReach, 120)
		Expect(e).NotTo(HaveOccurred(), "wait netReach task finish")

		success, e := common.CompareResult(frame, netReachName, pluginManager.KindNameNetReach, []string{}, reportNum, netReach)
		Expect(e).NotTo(HaveOccurred(), "compare report and task")
		Expect(success).To(BeTrue(), "compare report and task result")

	})

	It("Successfully testing using default daemonSet  as workload with more Task NetReach", Label("E00016"), func() {
		var e error
		successRate := float64(1)
		successMean := int64(1500)
		crontab := "0 1"
		netReachName := "netreach-" + tools.RandomName()

		netReach := new(v1beta1.NetReach)
		netReach.Name = netReachName

		// successCondition
		successCondition := new(v1beta1.NetSuccessCondition)
		successCondition.SuccessRate = &successRate
		successCondition.MeanAccessDelayInMs = &successMean
		netReach.Spec.SuccessCondition = successCondition
		enable := true
		disable := false
		// target
		target := new(v1beta1.NetReachTarget)
		if !common.TestIPv4 && common.TestIPv6 {
			target.Ingress = &disable
		} else {
			target.Ingress = &enable
		}
		target.LoadBalancer = &enable
		target.ClusterIP = &enable
		target.Endpoint = &enable
		target.NodePort = &enable
		target.MultusInterface = &disable
		target.IPv4 = &common.TestIPv4
		target.IPv6 = &common.TestIPv6
		netReach.Spec.Target = target

		// request
		request := new(v1beta1.NetHttpRequest)
		request.PerRequestTimeoutInMS = requestTimeout
		request.QPS = 5
		request.DurationInSecond = 5
		netReach.Spec.Request = request

		// Schedule
		Schedule := new(v1beta1.SchedulePlan)
		Schedule.Schedule = &crontab
		Schedule.RoundNumber = 1
		Schedule.RoundTimeoutMinute = 1
		netReach.Spec.Schedule = Schedule

		e = frame.CreateResource(netReach)
		Expect(e).NotTo(HaveOccurred(), "create netReach resource")

		e = common.CheckRuntime(frame, netReach, pluginManager.KindNameNetReach, 60)
		Expect(e).NotTo(HaveOccurred(), "check task runtime spec")

		e = common.WaitKdoctorTaskDone(frame, netReach, pluginManager.KindNameNetReach, 120)
		Expect(e).NotTo(HaveOccurred(), "wait netReach task finish")

		success, e := common.CompareResult(frame, netReachName, pluginManager.KindNameNetReach, []string{}, reportNum, netReach)
		Expect(e).NotTo(HaveOccurred(), "compare report and task")
		Expect(success).To(BeTrue(), "compare report and task result")

	})
})

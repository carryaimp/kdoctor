// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package kdoctor

import (
	internalinterfaces "github.com/kdoctor-io/kdoctor/pkg/k8s/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/kdoctor-io/kdoctor/pkg/k8s/client/informers/externalversions/kdoctor.io/v1beta1"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V1beta1 provides access to shared informers for resources in V1beta1.
	V1beta1() v1beta1.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// V1beta1 returns a new v1beta1.Interface.
func (g *group) V1beta1() v1beta1.Interface {
	return v1beta1.New(g.factory, g.namespace, g.tweakListOptions)
}

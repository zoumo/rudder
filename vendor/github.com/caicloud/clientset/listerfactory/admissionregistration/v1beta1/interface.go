/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	informers "k8s.io/client-go/informers"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/admissionregistration/v1beta1"
)

// Interface provides access to all the listers in this group version.
type Interface interface { // MutatingWebhookConfigurations returns a MutatingWebhookConfigurationLister
	MutatingWebhookConfigurations() v1beta1.MutatingWebhookConfigurationLister
	// ValidatingWebhookConfigurations returns a ValidatingWebhookConfigurationLister
	ValidatingWebhookConfigurations() v1beta1.ValidatingWebhookConfigurationLister
}

type version struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

type infromerVersion struct {
	factory informers.SharedInformerFactory
}

// New returns a new Interface.
func New(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{client: client, tweakListOptions: tweakListOptions}
}

// NewFrom returns a new Interface.
func NewFrom(factory informers.SharedInformerFactory) Interface {
	return &infromerVersion{factory: factory}
}

// MutatingWebhookConfigurations returns a MutatingWebhookConfigurationLister.
func (v *version) MutatingWebhookConfigurations() v1beta1.MutatingWebhookConfigurationLister {
	return &mutatingWebhookConfigurationLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// MutatingWebhookConfigurations returns a MutatingWebhookConfigurationLister.
func (v *infromerVersion) MutatingWebhookConfigurations() v1beta1.MutatingWebhookConfigurationLister {
	return v.factory.Admissionregistration().V1beta1().MutatingWebhookConfigurations().Lister()
}

// ValidatingWebhookConfigurations returns a ValidatingWebhookConfigurationLister.
func (v *version) ValidatingWebhookConfigurations() v1beta1.ValidatingWebhookConfigurationLister {
	return &validatingWebhookConfigurationLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// ValidatingWebhookConfigurations returns a ValidatingWebhookConfigurationLister.
func (v *infromerVersion) ValidatingWebhookConfigurations() v1beta1.ValidatingWebhookConfigurationLister {
	return v.factory.Admissionregistration().V1beta1().ValidatingWebhookConfigurations().Lister()
}
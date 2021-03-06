/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	coordinationv1beta1 "k8s.io/api/coordination/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/coordination/v1beta1"
)

var _ v1beta1.LeaseLister = &leaseLister{}

var _ v1beta1.LeaseNamespaceLister = &leaseNamespaceLister{}

// leaseLister implements the LeaseLister interface.
type leaseLister struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewLeaseLister returns a new LeaseLister.
func NewLeaseLister(client kubernetes.Interface) v1beta1.LeaseLister {
	return NewFilteredLeaseLister(client, nil)
}

func NewFilteredLeaseLister(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1beta1.LeaseLister {
	return &leaseLister{
		client:           client,
		tweakListOptions: tweakListOptions,
	}
}

// List lists all Leases in the indexer.
func (s *leaseLister) List(selector labels.Selector) (ret []*coordinationv1beta1.Lease, err error) {
	listopt := v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.CoordinationV1beta1().Leases(v1.NamespaceAll).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Leases returns an object that can list and get Leases.
func (s *leaseLister) Leases(namespace string) v1beta1.LeaseNamespaceLister {
	return leaseNamespaceLister{client: s.client, tweakListOptions: s.tweakListOptions, namespace: namespace}
}

// leaseNamespaceLister implements the LeaseNamespaceLister
// interface.
type leaseNamespaceLister struct {
	client           kubernetes.Interface
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// List lists all Leases in the indexer for a given namespace.
func (s leaseNamespaceLister) List(selector labels.Selector) (ret []*coordinationv1beta1.Lease, err error) {
	listopt := v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.CoordinationV1beta1().Leases(s.namespace).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Get retrieves the Lease from the indexer for a given namespace and name.
func (s leaseNamespaceLister) Get(name string) (*coordinationv1beta1.Lease, error) {
	return s.client.CoordinationV1beta1().Leases(s.namespace).Get(name, v1.GetOptions{})
}

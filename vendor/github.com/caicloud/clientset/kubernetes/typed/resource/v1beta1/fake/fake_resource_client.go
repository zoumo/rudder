/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/caicloud/clientset/kubernetes/typed/resource/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeResourceV1beta1 struct {
	*testing.Fake
}

func (c *FakeResourceV1beta1) Clusters() v1beta1.ClusterInterface {
	return &FakeClusters{c}
}

func (c *FakeResourceV1beta1) Configs() v1beta1.ConfigInterface {
	return &FakeConfigs{c}
}

func (c *FakeResourceV1beta1) Machines() v1beta1.MachineInterface {
	return &FakeMachines{c}
}

func (c *FakeResourceV1beta1) StorageServices() v1beta1.StorageServiceInterface {
	return &FakeStorageServices{c}
}

func (c *FakeResourceV1beta1) StorageTypes() v1beta1.StorageTypeInterface {
	return &FakeStorageTypes{c}
}

func (c *FakeResourceV1beta1) Tags() v1beta1.TagInterface {
	return &FakeTags{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeResourceV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
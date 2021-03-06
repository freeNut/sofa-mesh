/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "istio.io/istio/pkg/api/rpccontroller.istio.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RpcServiceLister helps list RpcServices.
type RpcServiceLister interface {
	// List lists all RpcServices in the indexer.
	List(selector labels.Selector) (ret []*v1.RpcService, err error)
	// RpcServices returns an object that can list and get RpcServices.
	RpcServices(namespace string) RpcServiceNamespaceLister
	RpcServiceListerExpansion
}

// rpcServiceLister implements the RpcServiceLister interface.
type rpcServiceLister struct {
	indexer cache.Indexer
}

// NewRpcServiceLister returns a new RpcServiceLister.
func NewRpcServiceLister(indexer cache.Indexer) RpcServiceLister {
	return &rpcServiceLister{indexer: indexer}
}

// List lists all RpcServices in the indexer.
func (s *rpcServiceLister) List(selector labels.Selector) (ret []*v1.RpcService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.RpcService))
	})
	return ret, err
}

// RpcServices returns an object that can list and get RpcServices.
func (s *rpcServiceLister) RpcServices(namespace string) RpcServiceNamespaceLister {
	return rpcServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RpcServiceNamespaceLister helps list and get RpcServices.
type RpcServiceNamespaceLister interface {
	// List lists all RpcServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.RpcService, err error)
	// Get retrieves the RpcService from the indexer for a given namespace and name.
	Get(name string) (*v1.RpcService, error)
	RpcServiceNamespaceListerExpansion
}

// rpcServiceNamespaceLister implements the RpcServiceNamespaceLister
// interface.
type rpcServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RpcServices in the indexer for a given namespace.
func (s rpcServiceNamespaceLister) List(selector labels.Selector) (ret []*v1.RpcService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.RpcService))
	})
	return ret, err
}

// Get retrieves the RpcService from the indexer for a given namespace and name.
func (s rpcServiceNamespaceLister) Get(name string) (*v1.RpcService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("rpcservice"), name)
	}
	return obj.(*v1.RpcService), nil
}

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

package v1alpha

import (
	v1alpha "github.com/cruise-automation/rbacsync/pkg/apis/rbacsync/v1alpha"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RBACSyncConfigLister helps list RBACSyncConfigs.
type RBACSyncConfigLister interface {
	// List lists all RBACSyncConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha.RBACSyncConfig, err error)
	// RBACSyncConfigs returns an object that can list and get RBACSyncConfigs.
	RBACSyncConfigs(namespace string) RBACSyncConfigNamespaceLister
	RBACSyncConfigListerExpansion
}

// rBACSyncConfigLister implements the RBACSyncConfigLister interface.
type rBACSyncConfigLister struct {
	indexer cache.Indexer
}

// NewRBACSyncConfigLister returns a new RBACSyncConfigLister.
func NewRBACSyncConfigLister(indexer cache.Indexer) RBACSyncConfigLister {
	return &rBACSyncConfigLister{indexer: indexer}
}

// List lists all RBACSyncConfigs in the indexer.
func (s *rBACSyncConfigLister) List(selector labels.Selector) (ret []*v1alpha.RBACSyncConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha.RBACSyncConfig))
	})
	return ret, err
}

// RBACSyncConfigs returns an object that can list and get RBACSyncConfigs.
func (s *rBACSyncConfigLister) RBACSyncConfigs(namespace string) RBACSyncConfigNamespaceLister {
	return rBACSyncConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RBACSyncConfigNamespaceLister helps list and get RBACSyncConfigs.
type RBACSyncConfigNamespaceLister interface {
	// List lists all RBACSyncConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha.RBACSyncConfig, err error)
	// Get retrieves the RBACSyncConfig from the indexer for a given namespace and name.
	Get(name string) (*v1alpha.RBACSyncConfig, error)
	RBACSyncConfigNamespaceListerExpansion
}

// rBACSyncConfigNamespaceLister implements the RBACSyncConfigNamespaceLister
// interface.
type rBACSyncConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RBACSyncConfigs in the indexer for a given namespace.
func (s rBACSyncConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha.RBACSyncConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha.RBACSyncConfig))
	})
	return ret, err
}

// Get retrieves the RBACSyncConfig from the indexer for a given namespace and name.
func (s rBACSyncConfigNamespaceLister) Get(name string) (*v1alpha.RBACSyncConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha.Resource("rbacsyncconfig"), name)
	}
	return obj.(*v1alpha.RBACSyncConfig), nil
}

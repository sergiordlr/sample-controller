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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "k8s.io/sample-controller/pkg/apis/samplecontroller/v1alpha1"
)

// DeployCustomLister helps list DeployCustoms.
type DeployCustomLister interface {
	// List lists all DeployCustoms in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DeployCustom, err error)
	// DeployCustoms returns an object that can list and get DeployCustoms.
	DeployCustoms(namespace string) DeployCustomNamespaceLister
	DeployCustomListerExpansion
}

// deployCustomLister implements the DeployCustomLister interface.
type deployCustomLister struct {
	indexer cache.Indexer
}

// NewDeployCustomLister returns a new DeployCustomLister.
func NewDeployCustomLister(indexer cache.Indexer) DeployCustomLister {
	return &deployCustomLister{indexer: indexer}
}

// List lists all DeployCustoms in the indexer.
func (s *deployCustomLister) List(selector labels.Selector) (ret []*v1alpha1.DeployCustom, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DeployCustom))
	})
	return ret, err
}

// DeployCustoms returns an object that can list and get DeployCustoms.
func (s *deployCustomLister) DeployCustoms(namespace string) DeployCustomNamespaceLister {
	return deployCustomNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DeployCustomNamespaceLister helps list and get DeployCustoms.
type DeployCustomNamespaceLister interface {
	// List lists all DeployCustoms in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DeployCustom, err error)
	// Get retrieves the DeployCustom from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DeployCustom, error)
	DeployCustomNamespaceListerExpansion
}

// deployCustomNamespaceLister implements the DeployCustomNamespaceLister
// interface.
type deployCustomNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DeployCustoms in the indexer for a given namespace.
func (s deployCustomNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DeployCustom, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DeployCustom))
	})
	return ret, err
}

// Get retrieves the DeployCustom from the indexer for a given namespace and name.
func (s deployCustomNamespaceLister) Get(name string) (*v1alpha1.DeployCustom, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("deploycustom"), name)
	}
	return obj.(*v1alpha1.DeployCustom), nil
}

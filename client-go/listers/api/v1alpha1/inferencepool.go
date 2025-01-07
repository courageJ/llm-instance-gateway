/*
Copyright 2024 The Kubernetes Authors.

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
	v1alpha1 "inference.networking.x-k8s.io/gateway-api-inference-extension/api/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// InferencePoolLister helps list InferencePools.
// All objects returned here must be treated as read-only.
type InferencePoolLister interface {
	// List lists all InferencePools in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InferencePool, err error)
	// InferencePools returns an object that can list and get InferencePools.
	InferencePools(namespace string) InferencePoolNamespaceLister
	InferencePoolListerExpansion
}

// inferencePoolLister implements the InferencePoolLister interface.
type inferencePoolLister struct {
	listers.ResourceIndexer[*v1alpha1.InferencePool]
}

// NewInferencePoolLister returns a new InferencePoolLister.
func NewInferencePoolLister(indexer cache.Indexer) InferencePoolLister {
	return &inferencePoolLister{listers.New[*v1alpha1.InferencePool](indexer, v1alpha1.Resource("inferencepool"))}
}

// InferencePools returns an object that can list and get InferencePools.
func (s *inferencePoolLister) InferencePools(namespace string) InferencePoolNamespaceLister {
	return inferencePoolNamespaceLister{listers.NewNamespaced[*v1alpha1.InferencePool](s.ResourceIndexer, namespace)}
}

// InferencePoolNamespaceLister helps list and get InferencePools.
// All objects returned here must be treated as read-only.
type InferencePoolNamespaceLister interface {
	// List lists all InferencePools in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InferencePool, err error)
	// Get retrieves the InferencePool from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InferencePool, error)
	InferencePoolNamespaceListerExpansion
}

// inferencePoolNamespaceLister implements the InferencePoolNamespaceLister
// interface.
type inferencePoolNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.InferencePool]
}

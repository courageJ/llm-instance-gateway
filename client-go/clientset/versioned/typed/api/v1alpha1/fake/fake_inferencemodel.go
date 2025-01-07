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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "inference.networking.x-k8s.io/gateway-api-inference-extension/api/v1alpha1"
	apiv1alpha1 "inference.networking.x-k8s.io/gateway-api-inference-extension/client-go/applyconfiguration/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInferenceModels implements InferenceModelInterface
type FakeInferenceModels struct {
	Fake *FakeApiV1alpha1
	ns   string
}

var inferencemodelsResource = v1alpha1.SchemeGroupVersion.WithResource("inferencemodels")

var inferencemodelsKind = v1alpha1.SchemeGroupVersion.WithKind("InferenceModel")

// Get takes name of the inferenceModel, and returns the corresponding inferenceModel object, and an error if there is any.
func (c *FakeInferenceModels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InferenceModel, err error) {
	emptyResult := &v1alpha1.InferenceModel{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(inferencemodelsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.InferenceModel), err
}

// List takes label and field selectors, and returns the list of InferenceModels that match those selectors.
func (c *FakeInferenceModels) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InferenceModelList, err error) {
	emptyResult := &v1alpha1.InferenceModelList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(inferencemodelsResource, inferencemodelsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InferenceModelList{ListMeta: obj.(*v1alpha1.InferenceModelList).ListMeta}
	for _, item := range obj.(*v1alpha1.InferenceModelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested inferenceModels.
func (c *FakeInferenceModels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(inferencemodelsResource, c.ns, opts))

}

// Create takes the representation of a inferenceModel and creates it.  Returns the server's representation of the inferenceModel, and an error, if there is any.
func (c *FakeInferenceModels) Create(ctx context.Context, inferenceModel *v1alpha1.InferenceModel, opts v1.CreateOptions) (result *v1alpha1.InferenceModel, err error) {
	emptyResult := &v1alpha1.InferenceModel{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(inferencemodelsResource, c.ns, inferenceModel, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.InferenceModel), err
}

// Update takes the representation of a inferenceModel and updates it. Returns the server's representation of the inferenceModel, and an error, if there is any.
func (c *FakeInferenceModels) Update(ctx context.Context, inferenceModel *v1alpha1.InferenceModel, opts v1.UpdateOptions) (result *v1alpha1.InferenceModel, err error) {
	emptyResult := &v1alpha1.InferenceModel{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(inferencemodelsResource, c.ns, inferenceModel, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.InferenceModel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInferenceModels) UpdateStatus(ctx context.Context, inferenceModel *v1alpha1.InferenceModel, opts v1.UpdateOptions) (result *v1alpha1.InferenceModel, err error) {
	emptyResult := &v1alpha1.InferenceModel{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(inferencemodelsResource, "status", c.ns, inferenceModel, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.InferenceModel), err
}

// Delete takes name of the inferenceModel and deletes it. Returns an error if one occurs.
func (c *FakeInferenceModels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(inferencemodelsResource, c.ns, name, opts), &v1alpha1.InferenceModel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInferenceModels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(inferencemodelsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InferenceModelList{})
	return err
}

// Patch applies the patch and returns the patched inferenceModel.
func (c *FakeInferenceModels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InferenceModel, err error) {
	emptyResult := &v1alpha1.InferenceModel{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(inferencemodelsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.InferenceModel), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied inferenceModel.
func (c *FakeInferenceModels) Apply(ctx context.Context, inferenceModel *apiv1alpha1.InferenceModelApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.InferenceModel, err error) {
	if inferenceModel == nil {
		return nil, fmt.Errorf("inferenceModel provided to Apply must not be nil")
	}
	data, err := json.Marshal(inferenceModel)
	if err != nil {
		return nil, err
	}
	name := inferenceModel.Name
	if name == nil {
		return nil, fmt.Errorf("inferenceModel.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.InferenceModel{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(inferencemodelsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.InferenceModel), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeInferenceModels) ApplyStatus(ctx context.Context, inferenceModel *apiv1alpha1.InferenceModelApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.InferenceModel, err error) {
	if inferenceModel == nil {
		return nil, fmt.Errorf("inferenceModel provided to Apply must not be nil")
	}
	data, err := json.Marshal(inferenceModel)
	if err != nil {
		return nil, err
	}
	name := inferenceModel.Name
	if name == nil {
		return nil, fmt.Errorf("inferenceModel.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.InferenceModel{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(inferencemodelsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.InferenceModel), err
}

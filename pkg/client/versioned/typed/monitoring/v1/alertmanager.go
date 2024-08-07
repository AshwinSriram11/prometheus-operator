// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1"
	scheme "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/scheme"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AlertmanagersGetter has a method to return a AlertmanagerInterface.
// A group's client should implement this interface.
type AlertmanagersGetter interface {
	Alertmanagers(namespace string) AlertmanagerInterface
}

// AlertmanagerInterface has methods to work with Alertmanager resources.
type AlertmanagerInterface interface {
	Create(ctx context.Context, alertmanager *v1.Alertmanager, opts metav1.CreateOptions) (*v1.Alertmanager, error)
	Update(ctx context.Context, alertmanager *v1.Alertmanager, opts metav1.UpdateOptions) (*v1.Alertmanager, error)
	UpdateStatus(ctx context.Context, alertmanager *v1.Alertmanager, opts metav1.UpdateOptions) (*v1.Alertmanager, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Alertmanager, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AlertmanagerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Alertmanager, err error)
	Apply(ctx context.Context, alertmanager *monitoringv1.AlertmanagerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Alertmanager, err error)
	ApplyStatus(ctx context.Context, alertmanager *monitoringv1.AlertmanagerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Alertmanager, err error)
	GetScale(ctx context.Context, alertmanagerName string, options metav1.GetOptions) (*autoscalingv1.Scale, error)
	UpdateScale(ctx context.Context, alertmanagerName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (*autoscalingv1.Scale, error)

	AlertmanagerExpansion
}

// alertmanagers implements AlertmanagerInterface
type alertmanagers struct {
	client rest.Interface
	ns     string
}

// newAlertmanagers returns a Alertmanagers
func newAlertmanagers(c *MonitoringV1Client, namespace string) *alertmanagers {
	return &alertmanagers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the alertmanager, and returns the corresponding alertmanager object, and an error if there is any.
func (c *alertmanagers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Alertmanager, err error) {
	result = &v1.Alertmanager{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alertmanagers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Alertmanagers that match those selectors.
func (c *alertmanagers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AlertmanagerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AlertmanagerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alertmanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested alertmanagers.
func (c *alertmanagers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("alertmanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a alertmanager and creates it.  Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *alertmanagers) Create(ctx context.Context, alertmanager *v1.Alertmanager, opts metav1.CreateOptions) (result *v1.Alertmanager, err error) {
	result = &v1.Alertmanager{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("alertmanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(alertmanager).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a alertmanager and updates it. Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *alertmanagers) Update(ctx context.Context, alertmanager *v1.Alertmanager, opts metav1.UpdateOptions) (result *v1.Alertmanager, err error) {
	result = &v1.Alertmanager{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alertmanagers").
		Name(alertmanager.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(alertmanager).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *alertmanagers) UpdateStatus(ctx context.Context, alertmanager *v1.Alertmanager, opts metav1.UpdateOptions) (result *v1.Alertmanager, err error) {
	result = &v1.Alertmanager{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alertmanagers").
		Name(alertmanager.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(alertmanager).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the alertmanager and deletes it. Returns an error if one occurs.
func (c *alertmanagers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alertmanagers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *alertmanagers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alertmanagers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched alertmanager.
func (c *alertmanagers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Alertmanager, err error) {
	result = &v1.Alertmanager{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("alertmanagers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied alertmanager.
func (c *alertmanagers) Apply(ctx context.Context, alertmanager *monitoringv1.AlertmanagerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Alertmanager, err error) {
	if alertmanager == nil {
		return nil, fmt.Errorf("alertmanager provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(alertmanager)
	if err != nil {
		return nil, err
	}
	name := alertmanager.Name
	if name == nil {
		return nil, fmt.Errorf("alertmanager.Name must be provided to Apply")
	}
	result = &v1.Alertmanager{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("alertmanagers").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *alertmanagers) ApplyStatus(ctx context.Context, alertmanager *monitoringv1.AlertmanagerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Alertmanager, err error) {
	if alertmanager == nil {
		return nil, fmt.Errorf("alertmanager provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(alertmanager)
	if err != nil {
		return nil, err
	}

	name := alertmanager.Name
	if name == nil {
		return nil, fmt.Errorf("alertmanager.Name must be provided to Apply")
	}

	result = &v1.Alertmanager{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("alertmanagers").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// GetScale takes name of the alertmanager, and returns the corresponding autoscalingv1.Scale object, and an error if there is any.
func (c *alertmanagers) GetScale(ctx context.Context, alertmanagerName string, options metav1.GetOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alertmanagers").
		Name(alertmanagerName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *alertmanagers) UpdateScale(ctx context.Context, alertmanagerName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alertmanagers").
		Name(alertmanagerName).
		SubResource("scale").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scale).
		Do(ctx).
		Into(result)
	return
}

/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2019 Red Hat, Inc.
 *
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubevirt.io/machine-remediation-operator/pkg/apis/machineremediation/v1alpha1"
)

// FakeMachineHealthChecks implements MachineHealthCheckInterface
type FakeMachineHealthChecks struct {
	Fake *FakeMachineremediationV1alpha1
	ns   string
}

var machinehealthchecksResource = schema.GroupVersionResource{Group: "machineremediation.kubevirt.io", Version: "v1alpha1", Resource: "machinehealthchecks"}

var machinehealthchecksKind = schema.GroupVersionKind{Group: "machineremediation.kubevirt.io", Version: "v1alpha1", Kind: "MachineHealthCheck"}

// Get takes name of the machineHealthCheck, and returns the corresponding machineHealthCheck object, and an error if there is any.
func (c *FakeMachineHealthChecks) Get(name string, options v1.GetOptions) (result *v1alpha1.MachineHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(machinehealthchecksResource, c.ns, name), &v1alpha1.MachineHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineHealthCheck), err
}

// List takes label and field selectors, and returns the list of MachineHealthChecks that match those selectors.
func (c *FakeMachineHealthChecks) List(opts v1.ListOptions) (result *v1alpha1.MachineHealthCheckList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(machinehealthchecksResource, machinehealthchecksKind, c.ns, opts), &v1alpha1.MachineHealthCheckList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MachineHealthCheckList{ListMeta: obj.(*v1alpha1.MachineHealthCheckList).ListMeta}
	for _, item := range obj.(*v1alpha1.MachineHealthCheckList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineHealthChecks.
func (c *FakeMachineHealthChecks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(machinehealthchecksResource, c.ns, opts))

}

// Create takes the representation of a machineHealthCheck and creates it.  Returns the server's representation of the machineHealthCheck, and an error, if there is any.
func (c *FakeMachineHealthChecks) Create(machineHealthCheck *v1alpha1.MachineHealthCheck) (result *v1alpha1.MachineHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(machinehealthchecksResource, c.ns, machineHealthCheck), &v1alpha1.MachineHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineHealthCheck), err
}

// Update takes the representation of a machineHealthCheck and updates it. Returns the server's representation of the machineHealthCheck, and an error, if there is any.
func (c *FakeMachineHealthChecks) Update(machineHealthCheck *v1alpha1.MachineHealthCheck) (result *v1alpha1.MachineHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(machinehealthchecksResource, c.ns, machineHealthCheck), &v1alpha1.MachineHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineHealthCheck), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMachineHealthChecks) UpdateStatus(machineHealthCheck *v1alpha1.MachineHealthCheck) (*v1alpha1.MachineHealthCheck, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(machinehealthchecksResource, "status", c.ns, machineHealthCheck), &v1alpha1.MachineHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineHealthCheck), err
}

// Delete takes name of the machineHealthCheck and deletes it. Returns an error if one occurs.
func (c *FakeMachineHealthChecks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(machinehealthchecksResource, c.ns, name), &v1alpha1.MachineHealthCheck{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineHealthChecks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(machinehealthchecksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MachineHealthCheckList{})
	return err
}

// Patch applies the patch and returns the patched machineHealthCheck.
func (c *FakeMachineHealthChecks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MachineHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(machinehealthchecksResource, c.ns, name, pt, data, subresources...), &v1alpha1.MachineHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineHealthCheck), err
}

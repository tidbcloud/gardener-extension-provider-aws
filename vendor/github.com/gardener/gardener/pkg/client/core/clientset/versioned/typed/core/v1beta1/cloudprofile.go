/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package v1beta1

import (
	"time"

	v1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	scheme "github.com/gardener/gardener/pkg/client/core/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloudProfilesGetter has a method to return a CloudProfileInterface.
// A group's client should implement this interface.
type CloudProfilesGetter interface {
	CloudProfiles() CloudProfileInterface
}

// CloudProfileInterface has methods to work with CloudProfile resources.
type CloudProfileInterface interface {
	Create(*v1beta1.CloudProfile) (*v1beta1.CloudProfile, error)
	Update(*v1beta1.CloudProfile) (*v1beta1.CloudProfile, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.CloudProfile, error)
	List(opts v1.ListOptions) (*v1beta1.CloudProfileList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.CloudProfile, err error)
	CloudProfileExpansion
}

// cloudProfiles implements CloudProfileInterface
type cloudProfiles struct {
	client rest.Interface
}

// newCloudProfiles returns a CloudProfiles
func newCloudProfiles(c *CoreV1beta1Client) *cloudProfiles {
	return &cloudProfiles{
		client: c.RESTClient(),
	}
}

// Get takes name of the cloudProfile, and returns the corresponding cloudProfile object, and an error if there is any.
func (c *cloudProfiles) Get(name string, options v1.GetOptions) (result *v1beta1.CloudProfile, err error) {
	result = &v1beta1.CloudProfile{}
	err = c.client.Get().
		Resource("cloudprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudProfiles that match those selectors.
func (c *cloudProfiles) List(opts v1.ListOptions) (result *v1beta1.CloudProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.CloudProfileList{}
	err = c.client.Get().
		Resource("cloudprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudProfiles.
func (c *cloudProfiles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("cloudprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudProfile and creates it.  Returns the server's representation of the cloudProfile, and an error, if there is any.
func (c *cloudProfiles) Create(cloudProfile *v1beta1.CloudProfile) (result *v1beta1.CloudProfile, err error) {
	result = &v1beta1.CloudProfile{}
	err = c.client.Post().
		Resource("cloudprofiles").
		Body(cloudProfile).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudProfile and updates it. Returns the server's representation of the cloudProfile, and an error, if there is any.
func (c *cloudProfiles) Update(cloudProfile *v1beta1.CloudProfile) (result *v1beta1.CloudProfile, err error) {
	result = &v1beta1.CloudProfile{}
	err = c.client.Put().
		Resource("cloudprofiles").
		Name(cloudProfile.Name).
		Body(cloudProfile).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudProfile and deletes it. Returns an error if one occurs.
func (c *cloudProfiles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("cloudprofiles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudProfiles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("cloudprofiles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudProfile.
func (c *cloudProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.CloudProfile, err error) {
	result = &v1beta1.CloudProfile{}
	err = c.client.Patch(pt).
		Resource("cloudprofiles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

/*
Copyright 2017 SAP SE

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

package v1

import (
	v1 "github.com/sapcc/kubernetes-operators/openstack-seeder/pkg/apis/seeder/v1"
	scheme "github.com/sapcc/kubernetes-operators/openstack-seeder/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OpenstackSeedsGetter has a method to return a OpenstackSeedInterface.
// A group's client should implement this interface.
type OpenstackSeedsGetter interface {
	OpenstackSeeds(namespace string) OpenstackSeedInterface
}

// OpenstackSeedInterface has methods to work with OpenstackSeed resources.
type OpenstackSeedInterface interface {
	Create(*v1.OpenstackSeed) (*v1.OpenstackSeed, error)
	Update(*v1.OpenstackSeed) (*v1.OpenstackSeed, error)
	UpdateStatus(*v1.OpenstackSeed) (*v1.OpenstackSeed, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.OpenstackSeed, error)
	List(opts meta_v1.ListOptions) (*v1.OpenstackSeedList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.OpenstackSeed, err error)
	OpenstackSeedExpansion
}

// openstackSeeds implements OpenstackSeedInterface
type openstackSeeds struct {
	client rest.Interface
	ns     string
}

// newOpenstackSeeds returns a OpenstackSeeds
func newOpenstackSeeds(c *OpenstackV1Client, namespace string) *openstackSeeds {
	return &openstackSeeds{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the openstackSeed, and returns the corresponding openstackSeed object, and an error if there is any.
func (c *openstackSeeds) Get(name string, options meta_v1.GetOptions) (result *v1.OpenstackSeed, err error) {
	result = &v1.OpenstackSeed{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("openstackseeds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OpenstackSeeds that match those selectors.
func (c *openstackSeeds) List(opts meta_v1.ListOptions) (result *v1.OpenstackSeedList, err error) {
	result = &v1.OpenstackSeedList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("openstackseeds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested openstackSeeds.
func (c *openstackSeeds) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("openstackseeds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a openstackSeed and creates it.  Returns the server's representation of the openstackSeed, and an error, if there is any.
func (c *openstackSeeds) Create(openstackSeed *v1.OpenstackSeed) (result *v1.OpenstackSeed, err error) {
	result = &v1.OpenstackSeed{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("openstackseeds").
		Body(openstackSeed).
		Do().
		Into(result)
	return
}

// Update takes the representation of a openstackSeed and updates it. Returns the server's representation of the openstackSeed, and an error, if there is any.
func (c *openstackSeeds) Update(openstackSeed *v1.OpenstackSeed) (result *v1.OpenstackSeed, err error) {
	result = &v1.OpenstackSeed{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("openstackseeds").
		Name(openstackSeed.Name).
		Body(openstackSeed).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *openstackSeeds) UpdateStatus(openstackSeed *v1.OpenstackSeed) (result *v1.OpenstackSeed, err error) {
	result = &v1.OpenstackSeed{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("openstackseeds").
		Name(openstackSeed.Name).
		SubResource("status").
		Body(openstackSeed).
		Do().
		Into(result)
	return
}

// Delete takes name of the openstackSeed and deletes it. Returns an error if one occurs.
func (c *openstackSeeds) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("openstackseeds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *openstackSeeds) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("openstackseeds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched openstackSeed.
func (c *openstackSeeds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.OpenstackSeed, err error) {
	result = &v1.OpenstackSeed{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("openstackseeds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

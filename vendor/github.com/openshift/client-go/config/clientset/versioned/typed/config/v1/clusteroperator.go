// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/openshift/api/config/v1"
	scheme "github.com/openshift/client-go/config/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterOperatorsGetter has a method to return a ClusterOperatorInterface.
// A group's client should implement this interface.
type ClusterOperatorsGetter interface {
	ClusterOperators() ClusterOperatorInterface
}

// ClusterOperatorInterface has methods to work with ClusterOperator resources.
type ClusterOperatorInterface interface {
	Create(ctx context.Context, clusterOperator *v1.ClusterOperator, opts metav1.CreateOptions) (*v1.ClusterOperator, error)
	Update(ctx context.Context, clusterOperator *v1.ClusterOperator, opts metav1.UpdateOptions) (*v1.ClusterOperator, error)
	UpdateStatus(ctx context.Context, clusterOperator *v1.ClusterOperator, opts metav1.UpdateOptions) (*v1.ClusterOperator, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ClusterOperator, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ClusterOperatorList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterOperator, err error)
	ClusterOperatorExpansion
}

// clusterOperators implements ClusterOperatorInterface
type clusterOperators struct {
	client rest.Interface
}

// newClusterOperators returns a ClusterOperators
func newClusterOperators(c *ConfigV1Client) *clusterOperators {
	return &clusterOperators{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterOperator, and returns the corresponding clusterOperator object, and an error if there is any.
func (c *clusterOperators) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterOperator, err error) {
	result = &v1.ClusterOperator{}
	err = c.client.Get().
		Resource("clusteroperators").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterOperators that match those selectors.
func (c *clusterOperators) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterOperatorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterOperatorList{}
	err = c.client.Get().
		Resource("clusteroperators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterOperators.
func (c *clusterOperators) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusteroperators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterOperator and creates it.  Returns the server's representation of the clusterOperator, and an error, if there is any.
func (c *clusterOperators) Create(ctx context.Context, clusterOperator *v1.ClusterOperator, opts metav1.CreateOptions) (result *v1.ClusterOperator, err error) {
	result = &v1.ClusterOperator{}
	err = c.client.Post().
		Resource("clusteroperators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterOperator).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterOperator and updates it. Returns the server's representation of the clusterOperator, and an error, if there is any.
func (c *clusterOperators) Update(ctx context.Context, clusterOperator *v1.ClusterOperator, opts metav1.UpdateOptions) (result *v1.ClusterOperator, err error) {
	result = &v1.ClusterOperator{}
	err = c.client.Put().
		Resource("clusteroperators").
		Name(clusterOperator.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterOperator).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterOperators) UpdateStatus(ctx context.Context, clusterOperator *v1.ClusterOperator, opts metav1.UpdateOptions) (result *v1.ClusterOperator, err error) {
	result = &v1.ClusterOperator{}
	err = c.client.Put().
		Resource("clusteroperators").
		Name(clusterOperator.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterOperator).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterOperator and deletes it. Returns an error if one occurs.
func (c *clusterOperators) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusteroperators").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterOperators) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusteroperators").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterOperator.
func (c *clusterOperators) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterOperator, err error) {
	result = &v1.ClusterOperator{}
	err = c.client.Patch(pt).
		Resource("clusteroperators").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
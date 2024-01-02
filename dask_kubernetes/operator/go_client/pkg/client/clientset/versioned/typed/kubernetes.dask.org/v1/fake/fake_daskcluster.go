// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	kubernetesdaskorgv1 "github.com/dask/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/apis/kubernetes.dask.org/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDaskClusters implements DaskClusterInterface
type FakeDaskClusters struct {
	Fake *FakeKubernetesV1
	ns   string
}

var daskclustersResource = schema.GroupVersionResource{Group: "kubernetes.dask.org", Version: "v1", Resource: "daskclusters"}

var daskclustersKind = schema.GroupVersionKind{Group: "kubernetes.dask.org", Version: "v1", Kind: "DaskCluster"}

// Get takes name of the daskCluster, and returns the corresponding daskCluster object, and an error if there is any.
func (c *FakeDaskClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *kubernetesdaskorgv1.DaskCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(daskclustersResource, c.ns, name), &kubernetesdaskorgv1.DaskCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubernetesdaskorgv1.DaskCluster), err
}

// List takes label and field selectors, and returns the list of DaskClusters that match those selectors.
func (c *FakeDaskClusters) List(ctx context.Context, opts v1.ListOptions) (result *kubernetesdaskorgv1.DaskClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(daskclustersResource, daskclustersKind, c.ns, opts), &kubernetesdaskorgv1.DaskClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubernetesdaskorgv1.DaskClusterList{ListMeta: obj.(*kubernetesdaskorgv1.DaskClusterList).ListMeta}
	for _, item := range obj.(*kubernetesdaskorgv1.DaskClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested daskClusters.
func (c *FakeDaskClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(daskclustersResource, c.ns, opts))

}

// Create takes the representation of a daskCluster and creates it.  Returns the server's representation of the daskCluster, and an error, if there is any.
func (c *FakeDaskClusters) Create(ctx context.Context, daskCluster *kubernetesdaskorgv1.DaskCluster, opts v1.CreateOptions) (result *kubernetesdaskorgv1.DaskCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(daskclustersResource, c.ns, daskCluster), &kubernetesdaskorgv1.DaskCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubernetesdaskorgv1.DaskCluster), err
}

// Update takes the representation of a daskCluster and updates it. Returns the server's representation of the daskCluster, and an error, if there is any.
func (c *FakeDaskClusters) Update(ctx context.Context, daskCluster *kubernetesdaskorgv1.DaskCluster, opts v1.UpdateOptions) (result *kubernetesdaskorgv1.DaskCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(daskclustersResource, c.ns, daskCluster), &kubernetesdaskorgv1.DaskCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubernetesdaskorgv1.DaskCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDaskClusters) UpdateStatus(ctx context.Context, daskCluster *kubernetesdaskorgv1.DaskCluster, opts v1.UpdateOptions) (*kubernetesdaskorgv1.DaskCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(daskclustersResource, "status", c.ns, daskCluster), &kubernetesdaskorgv1.DaskCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubernetesdaskorgv1.DaskCluster), err
}

// Delete takes name of the daskCluster and deletes it. Returns an error if one occurs.
func (c *FakeDaskClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(daskclustersResource, c.ns, name, opts), &kubernetesdaskorgv1.DaskCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDaskClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(daskclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &kubernetesdaskorgv1.DaskClusterList{})
	return err
}

// Patch applies the patch and returns the patched daskCluster.
func (c *FakeDaskClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kubernetesdaskorgv1.DaskCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(daskclustersResource, c.ns, name, pt, data, subresources...), &kubernetesdaskorgv1.DaskCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubernetesdaskorgv1.DaskCluster), err
}

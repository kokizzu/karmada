// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/karmada-io/karmada/pkg/apis/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeResourceExploringWebhookConfigurations implements ResourceExploringWebhookConfigurationInterface
type FakeResourceExploringWebhookConfigurations struct {
	Fake *FakeConfigV1alpha1
}

var resourceexploringwebhookconfigurationsResource = schema.GroupVersionResource{Group: "config.karmada.io", Version: "v1alpha1", Resource: "resourceexploringwebhookconfigurations"}

var resourceexploringwebhookconfigurationsKind = schema.GroupVersionKind{Group: "config.karmada.io", Version: "v1alpha1", Kind: "ResourceInterpreterWebhookConfiguration"}

// Get takes name of the resourceExploringWebhookConfiguration, and returns the corresponding resourceExploringWebhookConfiguration object, and an error if there is any.
func (c *FakeResourceExploringWebhookConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResourceInterpreterWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(resourceexploringwebhookconfigurationsResource, name), &v1alpha1.ResourceInterpreterWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceInterpreterWebhookConfiguration), err
}

// List takes label and field selectors, and returns the list of ResourceExploringWebhookConfigurations that match those selectors.
func (c *FakeResourceExploringWebhookConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResourceInterpreterWebhookConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(resourceexploringwebhookconfigurationsResource, resourceexploringwebhookconfigurationsKind, opts), &v1alpha1.ResourceInterpreterWebhookConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResourceInterpreterWebhookConfigurationList{ListMeta: obj.(*v1alpha1.ResourceInterpreterWebhookConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResourceInterpreterWebhookConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceExploringWebhookConfigurations.
func (c *FakeResourceExploringWebhookConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(resourceexploringwebhookconfigurationsResource, opts))
}

// Create takes the representation of a resourceExploringWebhookConfiguration and creates it.  Returns the server's representation of the resourceExploringWebhookConfiguration, and an error, if there is any.
func (c *FakeResourceExploringWebhookConfigurations) Create(ctx context.Context, resourceExploringWebhookConfiguration *v1alpha1.ResourceInterpreterWebhookConfiguration, opts v1.CreateOptions) (result *v1alpha1.ResourceInterpreterWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(resourceexploringwebhookconfigurationsResource, resourceExploringWebhookConfiguration), &v1alpha1.ResourceInterpreterWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceInterpreterWebhookConfiguration), err
}

// Update takes the representation of a resourceExploringWebhookConfiguration and updates it. Returns the server's representation of the resourceExploringWebhookConfiguration, and an error, if there is any.
func (c *FakeResourceExploringWebhookConfigurations) Update(ctx context.Context, resourceExploringWebhookConfiguration *v1alpha1.ResourceInterpreterWebhookConfiguration, opts v1.UpdateOptions) (result *v1alpha1.ResourceInterpreterWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(resourceexploringwebhookconfigurationsResource, resourceExploringWebhookConfiguration), &v1alpha1.ResourceInterpreterWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceInterpreterWebhookConfiguration), err
}

// Delete takes name of the resourceExploringWebhookConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeResourceExploringWebhookConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(resourceexploringwebhookconfigurationsResource, name), &v1alpha1.ResourceInterpreterWebhookConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceExploringWebhookConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(resourceexploringwebhookconfigurationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResourceInterpreterWebhookConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched resourceExploringWebhookConfiguration.
func (c *FakeResourceExploringWebhookConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResourceInterpreterWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(resourceexploringwebhookconfigurationsResource, name, pt, data, subresources...), &v1alpha1.ResourceInterpreterWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceInterpreterWebhookConfiguration), err
}

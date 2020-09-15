package client

import (
	"context"

	"github.com/victoriavilasb/ydata-crd/domain"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type Client struct {
	rest rest.Interface
	ns   string
}

func (c *Client) List(ctx context.Context, opts metav1.ListOptions) (*domain.Ydatas, error) {
	result := domain.Ydatas{}

	err := c.rest.
		Get().
		Namespace(c.ns).
		Resource("ydatas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *Client) Get(ctx context.Context, name string, opts metav1.GetOptions) (*domain.Ydata, error) {
	result := domain.Ydata{}
	err := c.rest.
		Get().
		Namespace(c.ns).
		Resource("ydatas").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *Client) Create(ctx context.Context, ydata *domain.Ydata) (*domain.Ydata, error) {
	result := domain.Ydata{}
	err := c.rest.
		Post().
		Namespace(c.ns).
		Resource("ydatas").
		Body(ydata).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *Client) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.rest.
		Get().
		Namespace(c.ns).
		Resource("ydatas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(ctx)
}

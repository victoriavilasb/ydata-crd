package client

import (
	"github.com/victoriavilasb/ydata-crd/domain"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type ydataClient struct {
	restClient rest.Interface
	ns         string
}

func (c *ydataClient) List(opts metav1.ListOptions) (*domain.YdataList, error) {
	result := domain.YdataList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("ydatas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *ydataClient) Get(name string, opts metav1.GetOptions) (*domain.Ydata, error) {
	result := domain.Ydata{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("ydatas").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *ydataClient) Create(ydata *domain.Ydata) (*domain.Ydata, error) {
	result := domain.Ydata{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("ydatas").
		Body(ydata).
		Do().
		Into(&result)

	return &result, err
}

func (c *ydataClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource("ydatas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

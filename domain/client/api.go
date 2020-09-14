package client

import (
	"github.com/victoriavilasb/ydata-crd/domain"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type YClient struct {
	rest rest.Interface
}

func NewClient(c *rest.Config) (*YClient, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: domain.GroupName, Version: domain.GroupVersion}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &YClient{rest: client}, nil
}

func (c *YClient) Ydatas(namespace string) domain.YdataInterface {
	return &Client{
		rest: c.rest,
		ns:   namespace,
	}
}

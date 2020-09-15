package domain

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

// YdataInterface ..
type YdataInterface interface {
	List(ctx context.Context, opts metav1.ListOptions) (*Ydatas, error)
	Get(ctx context.Context, name string, options metav1.GetOptions) (*Ydata, error)
	Create(context.Context, *Ydata) (*Ydata, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type YdataExample interface {
	Ydatas(namespace string) YdataInterface
}

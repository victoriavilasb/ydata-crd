package domain

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

// YdataInterface ..
type YdataInterface interface {
	List(opts metav1.ListOptions) (*YdataList, error)
	Get(name string, options metav1.GetOptions) (*Ydata, error)
	Create(*Ydata) (*Ydata, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type YdataExample interface {
	Ydatas(namespace string) YdataInterface
}

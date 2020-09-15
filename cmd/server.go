package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/victoriavilasb/ydata-crd/domain"
	"github.com/victoriavilasb/ydata-crd/domain/client"
	"github.com/victoriavilasb/ydata-crd/infra/errors"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig string

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to Kubernetes config file")
	flag.Parse()
}

func main() {
	ctx := context.Background()
	var err error
	var config *rest.Config
	if kubeconfig == "" {
		log.Printf("using in-cluster configuration")
		config, err = rest.InClusterConfig()
	} else {
		log.Printf("using configuration from '%s'", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	if err != nil {
		panic(errors.New(err, errors.KindUnexpected))
	}

	domain.AddToScheme(scheme.Scheme)

	clientSet, err := client.NewClientConfig(config)
	if err != nil {
		panic(err)
	}

	// ydatas, err := clientSet.Ydatas("kubeflow").List(ctx, metav1.ListOptions{})
	// if err != nil {
	// 	panic(errors.New(err, errors.KindUnexpected))
	// }

	// fmt.Printf("ydatas found: %+v\n", ydatas)

	watch := WatchResources(ctx, clientSet)

	for {
		projectsFromStore := watch.List()
		fmt.Printf("project in store: %d\n", len(projectsFromStore))

		time.Sleep(2 * time.Second)
	}
}

func WatchResources(ctx context.Context, clientSet *client.YClient) cache.Store {
	projectStore, projectController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return clientSet.Ydatas("kubeflow").List(ctx, lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return clientSet.Ydatas("kubeflow").Watch(ctx, lo)
			},
		},
		&domain.Ydata{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{},
	)

	go projectController.Run(wait.NeverStop)
	return projectStore
}

package application

import (
	"flag"
	"log"

	"github.com/victoriavilasb/ydata-crd/domain"
	"github.com/victoriavilasb/ydata-crd/infra/errors"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig string

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to Kubernetes config file")
	flag.Parse()
}

func main() {

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
}

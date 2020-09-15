# Ydata CRD 

### Intro

This repo contains an example of how to build CRD controllers in kubernetes API. 

To understand CRD, here's a great [resource](https://www.youtube.com/watch?v=xGafiZEX0YA&t=756s)

### How to run

This controller was deployed at GCP, and its image is in Google Container Registry.

To run this locally you must have a kubernetes cluster configured in your machine, and this cluster must have a namespace named "kubeflow" (The reason is because this controller was built to run in kubeflow). 

After that you can follow the steps below:
```sh
git clone git@github.com:victoriavilasb/ydata-crd.git
cd ydata-crd
make crd
make server
```

The **make crd** command will create the custom resource definition in your namespace, and the **make server** command will start the controller server in your machine, and it will be watching for changes in ydata resources.

### How to see it working

After running the controller in your machine you can create a sample Ydata resource by running the following command:

```
make sample-resource
```

You will see in your controller that a Ydata resource will have been created in your cache store. 

> ydata in store: 1

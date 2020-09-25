crd:
	@ kubectl apply -f manifests/ydata-crd.yaml

sample-resource:
	@ kubectl apply -f manifests/ydata.yaml

server:
	@ echo
	@ echo "Starting the server..."
	@ echo
	@ go run ./cmd --kubeconfig  ~/.kube/config

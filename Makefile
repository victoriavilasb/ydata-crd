crd:
	@ kubectl apply -f artifacts/ydata-crd.yaml

sample-crd:
	@ kubectl apply -f artifacts/ydata.yaml

server:
	@ echo
	@ echo "Starting the server..."
	@ echo
	@ go run ./cmd --kubeconfig  ~/.kube/config

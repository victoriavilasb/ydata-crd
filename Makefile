crd:
	@ kubectl apply -f artifacts/ydata-crd.yaml

server:
	@ echo
	@ echo "Starting the server..."
	@ echo
	@ go run ./cmd --kubeconfig  ~/.kube/config

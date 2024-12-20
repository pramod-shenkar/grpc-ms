hello:
	echo \
	"world" 

protoc:

	mkdir ./proto/$(service) || true
	protoc -I ./proto \
	--go_out ./proto/$(service) \
	--go_opt paths=source_relative \
	--go_grpc_out ./proto/$(service) \
	--go_grpc_opt paths=source_relative \
	./proto/$(service).proto

docker:
	docker build --build-arg servicename=order -t order-service:latest .
	docker-compose -f ./deploy/docker/docker-compose.yaml up --force-recreate --renew-anon-volumes -d

log:
	docker logs order-service

cmd:
	docker exec -it order-service sh


k8s:
	# kind create cluster --name grpc-go

	docker build --build-arg servicename=order -t order-service:latest .
	docker build --build-arg servicename=payment -t payment-service:latest .

	kind load docker-image order-service:latest --name grpc-go
	kind load docker-image payment-service:latest --name grpc-go


	kubectl delete  -f "./deploy/kubernetes/infra/db.yaml"
	kubectl delete  -f "./deploy/kubernetes/app/order.deployment.yaml"

	kubectl apply  -f "./deploy/kubernetes/infra/db.yaml"
	kubectl apply  -f "./deploy/kubernetes/app/order.deployment.yaml"


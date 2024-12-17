hello:
	echo \
	"world" 

protoc:

	# mkdir ./golang/$(service)/internal/adapter/grpc/pb || true
	# protoc -I ./proto \
	# --go_out ./golang/$(service)/internal/adapter/grpc/pb \
	# --go_opt paths=source_relative \
	# --go_grpc_out ./golang/$(service)/internal/adapter/grpc/pb \
	# --go_grpc_opt paths=source_relative \
	# ./proto/$(service).proto

	mkdir ./proto/$(service) || true
	protoc -I ./proto \
	--go_out ./proto/$(service) \
	--go_opt paths=source_relative \
	--go_grpc_out ./proto/$(service) \
	--go_grpc_opt paths=source_relative \
	./proto/$(service).proto


infra:
	docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=verysecretpass -e MYSQL_DATABASE=order mysql

connect-db:
	# docker exec -it $(service)-mysql mysqld -uroot -pverysecretpass order

deploy-docker:
	docker build --build-arg servicename=order -t order-service:latest .
	# docker-compose -f ./deploy/docker-compose.yaml up --force-recreate --renew-anon-volumes -d

k8s:
	kind create cluster --name grpc-go

	docker build --build-arg servicename=order -t order-service:latest .
	docker build --build-arg servicename=payment -t payment-service:latest .

	kind load docker-image order-service:latest --name grpc-go
	kind load docker-image payment-service:latest --name grpc-go


	# kubectl delete  -f "./deploy/kubernetes/infra/db.yaml"
	# kubectl delete  -f "./deploy/kubernetes/app/order.deployment.yaml"

	kubectl apply  -f "./deploy/kubernetes/infra/db.yaml"
	kubectl apply  -f "./deploy/kubernetes/app/order.deployment.yaml"


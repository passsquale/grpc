.PHONY: serv-gen
serv-gen:
	protoc -I=api --go_out=./pkg api/random-service/random-service.proto --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative

.PHONY: run
run:
	go run cmd/random/main.go
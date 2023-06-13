proto: 
	protoc --proto_path=packages --go_out=. --go-grpc_out=. packages/**/proto/*.proto

run: ## Running server
	go run server/main.go
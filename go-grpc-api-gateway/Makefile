proto: 
	protoc --proto_path=packages --go_out=. --go-grpc_out=. packages/**/proto/*.proto

server:
	go run server/main.go
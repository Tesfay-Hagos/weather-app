proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/services/**/pb/*.proto

server:
	go run cmd/main.go
docs:
	swag fmt && swag init -g cmd/main.go
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/pb/*.proto

server:
	go run cmd/main.go
migrate-down:
	migrate -database cockroachdb://root@localhost:26257/auth?sslmode=disable -path internal/constant/db/schema  -verbose down
migrate-up:
	migrate -database cockroachdb://root@localhost:26257/auth?sslmode=disable -path internal/constant/db/schema  -verbose up
migrate-create:
	migrate create -ext sql -dir internal/constant/db/schema -tz "UTC" $(args)
sqlc:
	sqlc generate -f sqlc.yaml
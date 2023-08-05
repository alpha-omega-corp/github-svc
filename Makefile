server:
	go run cmd/main.go

protoc:
	protoc --go_out=. \
		--go_opt=paths=source_relative \
        --go-grpc_out=. \
        --go-grpc_opt=paths=source_relative \
        proto/*.proto

db_create:
	docker-compose up -d

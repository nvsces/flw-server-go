build:
	docker-compose build flw-server-go

run:
	docker-compose up flw-server-go

# test:
# 	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up

# swag:
# 	swag init -g cmd/main.go


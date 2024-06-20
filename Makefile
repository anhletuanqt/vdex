PROJECT_NAME := $(shell basename "$(PWD)")

vdex_abi:
	cd contract && abigen --pkg vdex --abi ./vdex.abi --out ./vdex.go

run_database:
	docker run -d \
	-e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=vchain_verifier \
	--hostname database \
	--restart always \
	--publish 1234:5432 \
	--name vchain_verifier-database postgres:14-alpine

stop_database:
	docker stop vchain_verifier-database

start_database:
	docker start vchain_verifier-database

rm_database:
	docker rm vchain_verifier-database

migrate_db:
	docker run \
	-v ./migrations:/migrations --network host migrate/migrate:4 \
    -path=/migrations/ \
    -database "postgres://postgres:postgres@localhost:1234/vdex?sslmode=disable" up

build_service:
	mkdir -p build
	go build -o build/${PROJECT_NAME} cmd/main.go

run_service:
	(export $(cat app.env | xargs) >/dev/null && ./build/$(PROJECT_NAME))

run:
	go run main.go

build_cross_platform:
	mkdir -p build
	GOOS=linux GOARCH=amd64 go build -o build/${PROJECT_NAME}-linux-amd64 cmd/main.go
	GOOS=linux GOARCH=arm64 go build -o build/${PROJECT_NAME}-linux-arm64 cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o build/${PROJECT_NAME}-windows-amd64 cmd/main.go
	GOOS=darwin GOARCH=amd64 go build -o build/${PROJECT_NAME}-darwin-amd64 cmd/main.go
vdex-abi:
	cd contract && abigen --pkg vdex --abi ./vdex.abi --out ./vdex.go
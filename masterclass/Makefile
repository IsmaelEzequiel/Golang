#!make
include .env

postgres:
	docker run --name psql16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123 -d postgres:16.3-alpine3.20

createdb:
	docker exec -it psql16 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it psql16 dropdb simple_bank

status:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${DATABASE_URL} goose status

init:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${DATABASE_URL} goose -dir="./db/migrate" -s create init_schema sql

migrateup:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${DATABASE_URL} goose -dir="./db/migrate" up

migratedown:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${DATABASE_URL} goose -dir="./db/migrate" down

sqlc:
	snap run sqlc generate

test:
	go test -v -cover ./...

.PHONY: test sqlc init status postgres createdb dropdb migrateup migratedown


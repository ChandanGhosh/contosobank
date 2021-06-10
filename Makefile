postgres:
	docker run --name postgres12 -p5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

rmpostgres:
	docker rm -f postgres12

createdb:
	docker exec -it postgres12 createdb -U root -O root contoso_bank
dropdb:
	docker exec -it postgres12 dropdb contoso_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/contoso_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/contoso_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
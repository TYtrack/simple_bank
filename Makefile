mypostgres:
	docker run --name mypostgres -p 5432:5432 -e POSTGRES_USER=zplus -e POSTGRES_PASSWORD=123456 -d postgres

execpostgres:
	docker exec -it mypostgres /bin/sh

clipostgres:
	docker exec -it mypostgres psql -U zplus -d simple_bank

stoppostgres:
	docker stop mypostgres

rmpostgres:
	docker rm mypostgres

createdb:
	docker exec -it mypostgres createdb --username=zplus --owner=zplus simple_bank

dropdb:
	docker exec -it mypostgres dropdb --username=zplus simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://zplus:123456@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://zplus:123456@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://zplus:123456@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://zplus:123456@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
	
sqlc:
	sqlc generate
	
test :
	go test -v -cover ./...

server :
	go run main.go

mock :
	mockgen -package mockdb -destination db/mock/store.go bank_project/db/sqlc Store

.PHONY:
	mypostgres createdb dropdb
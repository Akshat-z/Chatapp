postgresinit:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it postgres15 dropdb go-chat

migrateup:
	docker run -it --rm --network host --user 1000:1000 --volume ./db:/db migrate/migrate:v4.17.0 -path=/db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose up

migratedown:
	docker run -it --rm --network host --user 1000:1000 --volume ./db:/db migrate/migrate:v4.17.0 -path=/db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose up

migration:
	docker run -it --rm --network host --user 1000:1000 --volume "$(pwd)/db:/db" migrate/migrate:v4.17.0 create -ext sql -dir /db/migrations users_table

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown migration




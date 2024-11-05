postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it postgres12 dropdb simplebank

migrateup:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose down

redis:
	docker run -d --name redis-stack-server -p 6379:6379 redis/redis-stack-server:latest

mysql:
	docker run --name mysql -e MYSQL_ROOT_PASSWORD=1 -p 3306:3306 -d mysql

.PHONY: postgres createdb dropdb migrateup migratedown mysql
postgres:
	docker run --name pg \
	-p 5432:5432 \
	-e POSTGRES_USER=root \
	-e POSTGRES_PASSWORD=secretPass123 \
	-d postgres:12-alpine 

pgadmin:
	docker run -p 80:80 \
	--name pgadmin \
    -e 'PGADMIN_DEFAULT_EMAIL=user@domain.com' \
    -e 'PGADMIN_DEFAULT_PASSWORD=SuperSecret' \
    -d dpage/pgadmin4 \

createdb:
	docker exec -it pg createdb --username=root --owner=root go-bank

dropdb:
	docker exec -it pg dropdb go-bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secretPass123@localhost:5432/go-bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secretPass123@localhost:5432/go-bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres pgadmin createdb dropdb migrateup migratedown sqlc test
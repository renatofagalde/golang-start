postgres:
	docker run --name CUSTOM  -e GIN_MODE=releasd --net=CUSTOM-network -p 5432:5432 -e POSTGRES_USER=user_api -e POSTGRES_PASSWORD=user_password -d postgres:12-alpine

createdb:
	docker exec -it CUSTOM createdb --username=user_api --owner=user_api CUSTOM

dropdb:
	docker exec -it CUSTOM dropdb  CUSTOM

migrationup:
	migrate -path db/migration -database "postgresql://user_api:user_password@localhost:5432/CUSTOM?sslmode=disable" -verbose up

migrationup1:
	migrate -path db/migration -database "postgresql://user_api:user_password@localhost:5432/CUSTOM?sslmode=disable" -verbose up 1

migrationdown:
	yes | migrate -path db/migration -database "postgresql://user_api:user_password@localhost:5432/CUSTOM?sslmode=disable" -verbose down

migrationdown1:
	yes | migrate -path db/migration -database "postgresql://user_api:user_password@localhost:5432/CUSTOM?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination ./db/mock/store.go ./bank/db/sqlc Store

init-db:
	docker-compose -f docker-compose-local.yaml up -d postgres
	sleep 10
	docker-compose -f docker-compose-local.yaml exec postgres psql -U custom_api -d CUSTOM -f /docker-entrypoint-initdb.d/init.sql

up: init-db
	docker-compose -f docker-compose-local.yaml up -d api


.PHONEY: postgres createdb dropdb migrationup migrationdown migrationup1 migrationdown1 sqlc test server mock init-db

customPostgres:
	docker run --name customPostgres -e GIN_MODE=release --net=CUSTOM-network -p 5432:5432 -e POSTGRES_USER=api -e POSTGRES_PASSWORD=password -d postgres:12-alpine

test:
	go test -v -cover ./...

server:
	go run main.go
#
# init-db:
# 	docker compose -f docker-compose-local.yaml up -d customPostgres
# 	sleep 10
# 	docker compose -f docker-compose-local.yaml exec customPostgres psql -U custom_api -d CUSTOM -f /docker-entrypoint-initdb.d/init.sql
#
# up: init-db
# 	docker compose -f docker-compose-local.yaml up -d customPostgres
#
.PHONEY: customPostgres test server init-db up

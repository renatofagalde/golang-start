#!/bin/sh

set -e

source /app/app.env
echo " ->DB_SOURCE = ${DB_SOURCE}"
echo " ->SERVER_ADDRES = ${SERVER_ADDRESS}"

ls -lh /app/migration/

/app/migrate.linux-amd64 -path /app/migration/ -database ${DB_SOURCE} -verbose up

echo "start the app"
exec "$@"

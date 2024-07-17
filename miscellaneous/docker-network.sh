#!/bin/bash

docker network create api-network

docker network connect custom-postgres-1

echo "Configuração network concluídas com sucesso!"


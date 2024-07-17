#!/bin/bash

# Parar todos os containers
echo "Parando todos os containers..."
docker stop $(docker ps -a -q)

# Remover todos os containers
echo "Removendo todos os containers..."
docker rm $(docker ps -a -q)

# Remover todas as imagens
echo "Removendo todas as imagens..."
docker rmi $(docker images -q)

# Remover todos os volumes
echo "Removendo todos os volumes..."
docker volume rm $(docker volume ls -q)

# Remover todas as networks não utilizadas
echo "Removendo todas as networks não utilizadas..."
docker network prune -f

# Limpar espaço de disco usado por Docker
echo "Limpando espaço de disco usado por Docker..."
docker system prune -a -f --volumes

echo "Limpeza completa do Docker concluída!"


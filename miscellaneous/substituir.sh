#!/bin/bash

# Defina o diretório que deseja procurar e substituir
DIRECTORY="."

# Substitua todas as ocorrências de "main/app/" por "main/app/" em todos os arquivos
grep -rl "main/app/" "$DIRECTORY" | while read -r file ; do
    sed -i 's|main/app/|main/app/|g' "$file"
done

echo "Substituição concluída."


#!/bin/bash

# Função para validar o domínio
validate_domain() {
    if [[ ! "$1" =~ ^[a-zA-Z]+$ ]]; then
        echo "Domínio inválido. O domínio deve conter apenas letras e não pode ter números, espaços ou caracteres especiais."
        exit 1
    fi
}

# Solicitar o domínio
read -p "Digite o domínio: " DOMAIN

# Validar o domínio
validate_domain "$DOMAIN"

# Converter o domínio para minúsculas
DOMAIN_LOWER=$(echo "$DOMAIN" | tr '[:upper:]' '[:lower:]')
# Converter a primeira letra para maiúscula
DOMAIN_CAPITALIZED=$(echo "$DOMAIN_LOWER" | sed 's/.*/\u&/')

# Função para renomear arquivos e diretórios
rename_files_and_directories() {
    find . -type f -name '*custom*' | while read -r file; do
        newfile=$(echo "$file" | sed "s/custom/$DOMAIN_LOWER/g")
        mv "$file" "$newfile"
    done

    find . -type d -name '*custom*' | while read -r dir; do
        newdir=$(echo "$dir" | sed "s/custom/$DOMAIN_LOWER/g")
        mv "$dir" "$newdir"
    done
}

# Função para substituir ocorrências dentro dos arquivos
replace_within_files() {
    find . -type f -not -name "setup.sh" -print0 | while IFS= read -r -d '' file; do
        # Substituir custom, Custom, CUSTOM respeitando a capitalização
        sed -i "s/\bcustom\b/$DOMAIN_LOWER/g" "$file"
        sed -i "s/\bCustom\b/$DOMAIN_CAPITALIZED/g" "$file"
        sed -i "s/\bCUSTOM\b/$DOMAIN/g" "$file"

        # Substituir ocorrências em camelCase e PascalCase respeitando a capitalização
        sed -i "s/\bCustom\([A-Z]\)/$DOMAIN_CAPITALIZED\1/g" "$file"
        sed -i "s/\bcustom\([A-Z]\)/$DOMAIN_LOWER\1/g" "$file"

        # Substituir ocorrências que tenham letras antes e depois
        sed -i "s/\([a-zA-Z]\)Custom\([a-zA-Z]\)/\1$DOMAIN_CAPITALIZED\2/g" "$file"
        sed -i "s/\([a-zA-Z]\)custom\([a-zA-Z]\)/\1$DOMAIN_LOWER\2/g" "$file"

        # Substituir rotas
        sed -i "s|/customs\b|/${DOMAIN_LOWER}s|g" "$file"
        sed -i "s|/custom\b|/$DOMAIN_LOWER|g" "$file"
    done
}

# Função para substituir ocorrências nos arquivos docker-compose.yml e docker-compose-local.yaml
replace_in_docker_compose_files() {
    for compose_file in docker-compose.yml docker-compose-local.yaml; do
        if [ -f "$compose_file" ]; then
            sed -i "s/\bCustom\b/$DOMAIN_CAPITALIZED/g" "$compose_file"
            sed -i "s/\bcustom\b/$DOMAIN_LOWER/g" "$compose_file"
            sed -i "s/\bCUSTOM\b/$DOMAIN/g" "$compose_file"
        fi
    done
}

# Renomear arquivos e diretórios
rename_files_and_directories

# Substituir ocorrências dentro dos arquivos
replace_within_files

# Substituir ocorrências nos arquivos docker-compose.yml e docker-compose-local.yaml
replace_in_docker_compose_files

# Inicializar o módulo Go
#go mod init main
# Baixar dependências do módulo Go
go mod tidy
# Executar testes com cobertura
# go test -v -cover ./...

go build -o .

# Subir o docker-compose usando docker-compose-local.yaml
docker compose -f docker-compose-local.yaml up -d

./miscellaneous/docker-network.sh
echo "Configuração e execução do Docker concluídas com sucesso!"


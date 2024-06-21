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

# Função para renomear arquivos e diretórios
rename_files_and_directories() {
    find . -name '*CUSTOM*' | while read file; do
        newfile=$(echo "$file" | sed "s/CUSTOM/$DOMAIN_LOWER/g")
        mv "$file" "$newfile"
    done

    # Renomear novamente para capturar subdiretórios e arquivos já renomeados
    find . -name '*CUSTOM*' | while read file; do
        newfile=$(echo "$file" | sed "s/CUSTOM/$DOMAIN_LOWER/g")
        mv "$file" "$newfile"
    done
}

# Função para substituir ocorrências dentro dos arquivos
replace_within_files() {
    find . -type f -not -name "setup.sh" -print0 | while IFS= read -r -d '' file; do
        sed -i "s/CUSTOM/$DOMAIN/g" "$file"
        sed -i "s/Custom/${DOMAIN^}/g" "$file"
        sed -i "s/custom/$DOMAIN_LOWER/g" "$file"
    done
}

# Renomear arquivos e diretórios
rename_files_and_directories

# Substituir ocorrências dentro dos arquivos
replace_within_files

echo "Configuração concluída com sucesso!"

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
    find . -name '*CUSTOM*' | while read -r file; do
        newfile=$(echo "$file" | sed "s/CUSTOM/$DOMAIN_LOWER/g")
        mv "$file" "$newfile"
    done

    # Renomear novamente para capturar subdiretórios e arquivos já renomeados
    find . -name '*CUSTOM*' | while read -r file; do
        newfile=$(echo "$file" | sed "s/CUSTOM/$DOMAIN_LOWER/g")
        mv "$file" "$newfile"
    done
}

# Função para substituir ocorrências dentro dos arquivos
replace_within_files() {
    find . -type f -not -name "setup.sh" -print0 | while IFS= read -r -d '' file; do
        # Substituir CUSTOM, Custom, custom respeitando a capitalização
#        sed -i "s/\bCUSTOM\b/$DOMAIN/g" "$file"
#        sed -i "s/\bCustom\b/$DOMAIN_CAPITALIZED/g" "$file"
#        sed -i "s/\bcustom\b/$DOMAIN_LOWER/g" "$file"

        # Substituir ocorrências em camelCase e PascalCase respeitando a capitalização
        sed -i "s/Custom\([A-Z]\)/$DOMAIN_CAPITALIZED\1/g" "$file"
        sed -i "s/custom\([A-Z]\)/$DOMAIN_LOWER\1/g" "$file"

        # Substituir rotas
        sed -i "s|/customs\b|/${DOMAIN_LOWER}s|g" "$file"
        sed -i "s|/custom\b|/$DOMAIN_LOWER|g" "$file"

        # Substituir todas as ocorrências de custom, Custom, CUSTOM independentemente do contexto
        sed -i "s/custom/$DOMAIN_LOWER/gI" "$file"
        sed -i "s/Custom/$DOMAIN_CAPITALIZED/gI" "$file"
        sed -i "s/CUSTOM/$DOMAIN/gI" "$file"
    done
}

# Renomear arquivos e diretórios
rename_files_and_directories

# Substituir ocorrências dentro dos arquivos
replace_within_files

# Inicializar o módulo Go
go mod init main
# Baixar dependências do módulo Go
go mod tidy
# Executar testes com cobertura
go test -v -cover ./...

echo "Configuração concluída com sucesso!"


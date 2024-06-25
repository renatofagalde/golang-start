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
        # Substituir "CUSTOM", "Custom", "custom"
        sed -i "s/\([*&]\?\)\bCUSTOM\b/\1$DOMAIN/g" "$file"
        sed -i "s/\([*&]\?\)\bCustom\b/\1$DOMAIN_CAPITALIZED/g" "$file"
        sed -i "s/\([*&]\?\)\bcustom\b/\1$DOMAIN_LOWER/g" "$file"

        # Substituir rotas que contenham "custom" no plural
        sed -i "s|/customs\b|/${DOMAIN_LOWER}s|g" "$file"
        
        # Substituir rotas que contenham "custom"
        sed -i "s|/custom\b|/$DOMAIN_LOWER|g" "$file"

        # Substituir variáveis e assinaturas de funções
        sed -i "s/\b\([*&]\?\)customController\b/\1${DOMAIN_LOWER}Controller/g" "$file"
        sed -i "s/\b\([*&]\?\)CustomController\b/\1${DOMAIN_CAPITALIZED}Controller/g" "$file"
        sed -i "s/\b\([*&]\?\)customRequest\b/\1${DOMAIN_LOWER}Request/g" "$file"
        sed -i "s/\b\([*&]\?\)CustomRequest\b/\1${DOMAIN_CAPITALIZED}Request/g" "$file"
        sed -i "s/\b\([*&]\?\)customControllerInterface\b/\1${DOMAIN_LOWER}ControllerInterface/g" "$file"
        sed -i "s/\b\([*&]\?\)CustomControllerInterface\b/\1${DOMAIN_CAPITALIZED}ControllerInterface/g" "$file"

        # Substituir NewCustomDomain e customDomain, incluindo casos entre parênteses
        sed -i "s/\bNewCustomDomain\b/New${DOMAIN_CAPITALIZED}Domain/g" "$file"
        sed -i "s/\bnewCustomDomain\b/new${DOMAIN_CAPITALIZED}Domain/g" "$file"
        sed -i "s/\bnewcustomdomain\b/new${DOMAIN_LOWER}domain/g" "$file"
        sed -i "s/\bcustomDomain\b/${DOMAIN_LOWER}Domain/g" "$file"
        sed -i "s/\bCustomDomain\b/${DOMAIN_CAPITALIZED}Domain/g" "$file"
        sed -i "s/(\s*customDomain\s*)/(${DOMAIN_LOWER}Domain)/g" "$file"
        sed -i "s/(\s*CustomDomain\s*)/(${DOMAIN_CAPITALIZED}Domain)/g" "$file"

        # Substituir CustomDomainService
        sed -i "s/\bCustomDomainService\b/${DOMAIN_CAPITALIZED}DomainService/g" "$file"
        sed -i "s/\bcustomDomainService\b/${DOMAIN_LOWER}DomainService/g" "$file"

        # Substituir NewCustomControllerInterface
        sed -i "s/\bNewCustomControllerInterface\b/New${DOMAIN_CAPITALIZED}ControllerInterface/g" "$file"
        sed -i "s/\bnewCustomControllerInterface\b/new${DOMAIN_CAPITALIZED}ControllerInterface/g" "$file"
        sed -i "s/\bnewcustomcontrollerinterface\b/new${DOMAIN_LOWER}controllerinterface/g" "$file"

        # Substituir CustomDomainInterface
        sed -i "s/\bCustomDomainInterface\b/${DOMAIN_CAPITALIZED}DomainInterface/g" "$file"
        sed -i "s/\bcustomDomainInterface\b/${DOMAIN_LOWER}DomainInterface/g" "$file"

        # Substituir GetCustom e variações
        sed -i "s/\bGetCustom\b/Get${DOMAIN_CAPITALIZED}/g" "$file"
        sed -i "s/\bgetCustom\b/get${DOMAIN_CAPITALIZED}/g" "$file"
        sed -i "s/\bgetcustom\b/get${DOMAIN_LOWER}/g" "$file"
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
:

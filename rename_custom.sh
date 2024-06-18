
#!/bin/bash

# Solicita ao usuário o novo valor para substituir "CUSTOM"
read -p "Digite o novo valor para substituir 'CUSTOM': " novo_valor

# Converte o valor fornecido pelo usuário para minúsculas
novo_valor=$(echo "$novo_valor" | tr '[:upper:]' '[:lower:]')

# Verifica se o usuário não forneceu um valor vazio
if [ -z "$novo_valor" ]; then
    echo "Erro: Você deve fornecer um valor válido."
    exit 1
fi

# Encontra e renomeia a palavra "CUSTOM" nos nomes de arquivos e diretórios
find . -depth -name '*CUSTOM*' | while read file; do
    novo_nome=$(echo "$file" | sed "s/CUSTOM/$novo_valor/g")
    mv "$file" "$novo_nome"
done

# Encontra e substitui a palavra "CUSTOM" dentro dos arquivos
grep -rl 'CUSTOM' . | while read file; do
    sed -i "s/CUSTOM/$novo_valor/g" "$file"
done

echo "Substituição concluída. Todas as ocorrências de 'CUSTOM' foram renomeadas para '$novo_valor'."

# Exclui o próprio script
rm -- "$0"

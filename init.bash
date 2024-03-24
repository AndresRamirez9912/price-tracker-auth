#!/bin/bash

# Leer el archivo .env línea por línea
while IFS= read -r line; do
  # Ignorar líneas en blanco o comentarios
  if [[ "$line" =~ ^[[:space:]]*$ || "$line" =~ ^# ]]; then
    continue
  fi

  # Dividir la línea en nombre y valor de la variable
  key=$(echo "$line" | cut -d '=' -f 1 | tr -d '[:space:]')
  value=$(echo "$line" | cut -d '=' -f 2- | sed 's/^"\|"$//g')  # Remove leading and trailing double quotes

  # Establecer la variable de entorno
  export $key=$value
done < .env

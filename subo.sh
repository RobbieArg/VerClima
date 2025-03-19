#!/bin/bash

# Si el usuario proporciona un mensaje, lo usa; si no, usa el mensaje por defecto
mensaje="${1:-Siguiente Commit}"

git add .
git commit -m "$mensaje"
git push

# Mostrar el estado del repositorio después del push
echo -e "\nEstado del repositorio después del push:"
git status

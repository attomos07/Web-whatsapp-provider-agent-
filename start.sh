#!/bin/bash

# Script de inicio rápido para WhatsApp Bot Go
# Este script verifica las dependencias y ejecuta el bot

set -e

echo "🤖 WhatsApp Bot - Inicio Rápido"
echo "================================"
echo ""

# Verificar si Go está instalado
if ! command -v go &> /dev/null; then
    echo "❌ Go no está instalado."
    echo "Por favor instala Go desde: https://golang.org/dl/"
    exit 1
fi

GO_VERSION=$(go version | awk '{print $3}')
echo "✅ Go detectado: $GO_VERSION"
echo ""

# Verificar si existe go.mod
if [ ! -f "go.mod" ]; then
    echo "❌ go.mod no encontrado."
    echo "Asegúrate de estar en el directorio correcto."
    exit 1
fi

# Instalar dependencias si es necesario
echo "📦 Verificando dependencias..."
go mod download
go mod tidy
echo "✅ Dependencias verificadas"
echo ""

# Verificar si existe la base de datos
if [ -f "whatsapp.db" ]; then
    echo "📱 Sesión existente detectada"
    echo "El bot se conectará automáticamente"
else
    echo "📱 Primera ejecución detectada"
    echo "Se mostrará un código QR para escanear con WhatsApp"
fi
echo ""

# Iniciar el bot
echo "🚀 Iniciando bot..."
echo "================================"
echo ""

go run .

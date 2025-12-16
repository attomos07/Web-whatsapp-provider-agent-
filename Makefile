.PHONY: help run build clean install docker-build docker-run test

# Variables
BINARY_NAME=whatsapp-bot
DOCKER_IMAGE=whatsapp-bot-go

help: ## Mostrar esta ayuda
	@echo "Comandos disponibles:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

install: ## Instalar dependencias
	@echo "📦 Instalando dependencias..."
	go mod download
	go mod tidy
	@echo "✅ Dependencias instaladas"

run: ## Ejecutar el bot
	@echo "🚀 Iniciando bot..."
	go run .

build: ## Compilar el binario
	@echo "🔨 Compilando..."
	go build -o $(BINARY_NAME) .
	@echo "✅ Binario creado: $(BINARY_NAME)"

build-linux: ## Compilar para Linux
	@echo "🔨 Compilando para Linux..."
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux .
	@echo "✅ Binario Linux creado: $(BINARY_NAME)-linux"

build-windows: ## Compilar para Windows
	@echo "🔨 Compilando para Windows..."
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe .
	@echo "✅ Binario Windows creado: $(BINARY_NAME).exe"

build-mac: ## Compilar para macOS
	@echo "🔨 Compilando para macOS..."
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-mac .
	@echo "✅ Binario macOS creado: $(BINARY_NAME)-mac"

build-all: build-linux build-windows build-mac ## Compilar para todas las plataformas
	@echo "✅ Todos los binarios creados"

clean: ## Limpiar archivos compilados
	@echo "🧹 Limpiando..."
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)-linux
	rm -f $(BINARY_NAME).exe
	rm -f $(BINARY_NAME)-mac
	@echo "✅ Archivos limpiados"

clean-db: ## Eliminar la base de datos (desconectará el bot)
	@echo "⚠️  ¿Estás seguro? Esto eliminará la sesión de WhatsApp."
	@echo "Presiona Ctrl+C para cancelar, Enter para continuar..."
	@read
	rm -f whatsapp.db*
	@echo "✅ Base de datos eliminada"

docker-build: ## Construir imagen Docker
	@echo "🐳 Construyendo imagen Docker..."
	docker build -t $(DOCKER_IMAGE) .
	@echo "✅ Imagen Docker creada: $(DOCKER_IMAGE)"

docker-run: ## Ejecutar en Docker
	@echo "🐳 Ejecutando en Docker..."
	docker run -it -v $(PWD):/root $(DOCKER_IMAGE)

test: ## Ejecutar tests
	@echo "🧪 Ejecutando tests..."
	go test -v ./...

format: ## Formatear código
	@echo "📝 Formateando código..."
	go fmt ./...
	@echo "✅ Código formateado"

lint: ## Ejecutar linter
	@echo "🔍 Ejecutando linter..."
	golangci-lint run
	@echo "✅ Linter completado"

deps-update: ## Actualizar dependencias
	@echo "🔄 Actualizando dependencias..."
	go get -u ./...
	go mod tidy
	@echo "✅ Dependencias actualizadas"

dev: ## Modo desarrollo con hot-reload (requiere air)
	@echo "🔥 Iniciando en modo desarrollo..."
	air

prod: build ## Ejecutar en producción
	@echo "🚀 Ejecutando en producción..."
	./$(BINARY_NAME)

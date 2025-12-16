# 🔧 Solución al Error de Instalación

## Error que estás viendo:
```
go: go.mau.fi/whatsmeow@v0.0.0-20240124191856-c311808d6e2e: invalid version: unknown revision
```

## ✅ Solución:

### Opción 1: Reinstalar dependencias (RECOMENDADO)

```bash
# 1. Eliminar go.mod y go.sum existentes
rm go.mod go.sum

# 2. Inicializar módulo nuevo
go mod init whatsapp-bot-go

# 3. Agregar dependencias con versiones actuales
go get go.mau.fi/whatsmeow@latest
go get github.com/mattn/go-sqlite3@latest
go get github.com/mdp/qrterminal/v3@latest
go get golang.org/x/text@latest
go get google.golang.org/protobuf@latest

# 4. Limpiar y verificar
go mod tidy

# 5. Ejecutar
go run .
```

### Opción 2: Actualización rápida

```bash
# Actualizar todas las dependencias
go get -u ./...
go mod tidy
go run .
```

### Opción 3: Script automático

Copia y pega todo esto en tu terminal:

```bash
#!/bin/bash
cd whatsapp-bot-go
rm -f go.mod go.sum
go mod init whatsapp-bot-go
go get go.mau.fi/whatsmeow@latest
go get github.com/mattn/go-sqlite3@latest  
go get github.com/mdp/qrterminal/v3@latest
go get golang.org/x/text@latest
go get google.golang.org/protobuf@latest
go mod tidy
echo "✅ Dependencias instaladas correctamente"
go run .
```

## 📝 Verificar instalación

Después de ejecutar los comandos, deberías ver:

```bash
go list -m all
```

Esto debe mostrar algo como:
```
whatsapp-bot-go
github.com/mattn/go-sqlite3 v1.14.x
github.com/mdp/qrterminal/v3 v3.2.0
go.mau.fi/whatsmeow v0.0.0-20241XXX...
...
```

## 🎯 go.mod correcto (ejemplo)

Tu `go.mod` debería verse así después:

```go
module whatsapp-bot-go

go 1.21

require (
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/mdp/qrterminal/v3 v3.2.0
	go.mau.fi/whatsmeow v0.0.0-20241215xxxxx-xxxxxxxxxxxx
	golang.org/x/text v0.21.0
	google.golang.org/protobuf v1.35.2
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	go.mau.fi/libsignal v0.1.1 // indirect
	go.mau.fi/util v0.8.2 // indirect
	golang.org/x/crypto v0.29.0 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	rsc.io/qr v0.2.0 // indirect
)
```

(Los números de versión exactos pueden variar, eso es normal)

## 🚀 Después de instalar correctamente

```bash
# Ejecutar el bot
go run .

# Deberías ver:
# 🔐 Escanea este código QR con tu WhatsApp:
# [QR CODE aparece aquí]
```

## 💡 Nota importante

El archivo `go.mod` que incluí tiene una versión específica de `whatsmeow` que puede haber cambiado. Go necesita descargar la versión más reciente directamente desde el repositorio de GitHub.

Por eso es mejor:
1. Eliminar el `go.mod` incluido
2. Dejar que Go genere uno nuevo con las versiones actuales
3. Esto asegura compatibilidad con la última versión de whatsmeow

## ❓ Si aún tienes problemas

Verifica que tienes:
- Go 1.21 o superior: `go version`
- Conexión a internet (para descargar dependencias)
- Permisos de escritura en el directorio

Si el problema persiste:
```bash
# Limpiar todo el cache de Go
go clean -modcache

# Luego volver a intentar
rm go.mod go.sum
go mod init whatsapp-bot-go
go get go.mau.fi/whatsmeow@latest
# ... resto de comandos
```

## ✅ Una vez funcionando

Guarda tu `go.mod` y `go.sum` actualizados. Estos archivos ahora tendrán las versiones correctas y funcionarán sin problemas.

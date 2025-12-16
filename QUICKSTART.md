# 🚀 Instalación Rápida - WhatsApp Bot Go

Guía paso a paso para poner en marcha tu bot en menos de 5 minutos.

## ✅ Pre-requisitos

1. **Go instalado** (versión 1.21 o superior)
   - Descarga: https://golang.org/dl/
   - Verifica: `go version`

2. **Git instalado** (opcional, para clonar)
   - Verifica: `git --version`

3. **WhatsApp en tu teléfono**
   - Con número activo
   - Conexión a internet

## 📦 Opción 1: Usando el archivo comprimido

```bash
# 1. Extraer el archivo
tar -xzf whatsapp-bot-go.tar.gz

# 2. Entrar al directorio
cd whatsapp-bot-go

# 3. Instalar dependencias
go mod download

# 4. Ejecutar el bot
go run .
```

## 🔧 Opción 2: Instalación Manual

```bash
# 1. Crear directorio del proyecto
mkdir whatsapp-bot-go
cd whatsapp-bot-go

# 2. Copiar todos los archivos .go al directorio

# 3. Inicializar módulo Go
go mod init whatsapp-bot-go

# 4. Agregar dependencias
go get go.mau.fi/whatsmeow
go get github.com/mattn/go-sqlite3
go get github.com/mdp/qrterminal/v3
go get golang.org/x/text
go get google.golang.org/protobuf

# 5. Limpiar dependencias
go mod tidy

# 6. Ejecutar
go run .
```

## 📱 Primera Ejecución

1. **Ejecuta el bot:**
   ```bash
   go run .
   ```

2. **Verás un código QR en la terminal:**
   ```
   🔐 Escanea este código QR con tu WhatsApp:
   
   ███████████████████████
   ███████████████████████
   ███████████████████████
   ```

3. **Abre WhatsApp en tu teléfono:**
   - **Android**: Menú (⋮) → Dispositivos vinculados → Vincular un dispositivo
   - **iPhone**: Configuración → Dispositivos vinculados → Vincular un dispositivo

4. **Escanea el QR**

5. **¡Listo!** Verás:
   ```
   ✅ Bot conectado exitosamente!
   📱 Esperando mensajes...
   ```

## 🎯 Probar el Bot

Envía un mensaje desde otro teléfono al número vinculado:

```
Tú: Hola
Bot: ¡Hola [Tu nombre]! ✂️ Soy el asistente virtual...
```

## 🛑 Detener el Bot

Presiona `Ctrl + C` en la terminal.

## 🔄 Ejecuciones Posteriores

```bash
# Ya no necesitas escanear QR
go run .
```

La sesión queda guardada en `whatsapp.db`.

## 🏗️ Compilar Binario (Opcional)

Para crear un ejecutable:

```bash
# Compilar
go build -o bot .

# Ejecutar
./bot
```

O usa el Makefile:

```bash
make build    # Compilar
make run      # Ejecutar sin compilar
make clean    # Limpiar archivos
```

## 🐳 Docker (Opcional)

```bash
# Construir imagen
docker build -t whatsapp-bot .

# Ejecutar
docker run -it -v $(pwd):/root whatsapp-bot
```

## ⚙️ Configuración

Edita `config.go` para cambiar:
- Nombre del negocio
- Servicios y precios
- Promociones
- Horarios
- Ubicación

## 📝 Estructura de Archivos

```
whatsapp-bot-go/
├── main.go          # Punto de entrada
├── flows.go         # Flujos del bot
├── utils.go         # Funciones auxiliares
├── config.go        # Configuración
├── go.mod           # Dependencias
├── README.md        # Documentación
├── EXAMPLES.md      # Ejemplos de uso
├── API.md           # Referencia técnica
├── Makefile         # Comandos útiles
├── Dockerfile       # Para Docker
├── start.sh         # Script de inicio
└── .gitignore       # Archivos ignorados
```

## 🔐 Seguridad

**IMPORTANTE:** 

- ✅ Agrega `whatsapp.db` a `.gitignore`
- ❌ NUNCA compartas `whatsapp.db` (contiene tu sesión)
- ❌ NUNCA subas `whatsapp.db` a GitHub

## 🐛 Solución de Problemas

### Error: "no Go files"
```bash
# Asegúrate de estar en el directorio correcto
cd whatsapp-bot-go
ls *.go  # Debe mostrar: main.go flows.go utils.go config.go
```

### Error: "package not found"
```bash
# Reinstalar dependencias
go mod download
go mod tidy
```

### El QR no aparece
```bash
# Verifica que tienes conexión a internet
ping google.com

# Verifica la versión de Go
go version  # Debe ser 1.21+
```

### Error: "database locked"
```bash
# Elimina la base de datos y vuelve a escanear
rm whatsapp.db*
go run .
```

### El bot no responde
- Verifica que WhatsApp Web esté activo
- Revisa los logs en la terminal
- Asegúrate de que el teléfono tenga internet

## 📚 Comandos Útiles

```bash
# Ejecutar
go run .

# Compilar
go build -o bot .

# Ver dependencias
go list -m all

# Actualizar dependencias
go get -u ./...

# Limpiar cache
go clean -cache

# Formatear código
go fmt ./...

# Hacer tests
go test ./...
```

## 🚀 Modo Producción

Para usar en servidor:

```bash
# 1. Compilar para Linux
GOOS=linux GOARCH=amd64 go build -o bot .

# 2. Copiar al servidor (vía SCP)
scp bot usuario@servidor:/ruta/

# 3. En el servidor, ejecutar
./bot

# 4. Para mantener corriendo (con screen o tmux)
screen -S whatsapp-bot
./bot
# Ctrl+A, D para separar
```

## 🎓 Próximos Pasos

1. Lee `README.md` para documentación completa
2. Revisa `EXAMPLES.md` para ver ejemplos de conversaciones
3. Lee `API.md` para entender la arquitectura
4. Modifica `config.go` para personalizar
5. Edita `flows.go` para agregar nuevos flujos

## 💡 Personalización Rápida

### Cambiar nombre del negocio
Edita `config.go`:
```go
const BUSINESS_NAME = "Tu Barbería"
```

### Agregar nuevo servicio
En `config.go`:
```go
var SERVICES = map[string]int{
    "Tu Nuevo Servicio": 150,
}
```

### Modificar mensaje de bienvenida
En `flows.go`, función `processMessageWithFlows()`:
```go
return fmt.Sprintf("¡Hola %s! Tu mensaje personalizado...", name)
```

## 📞 Soporte

- Issues: GitHub Issues
- Documentación: README.md, API.md
- Ejemplos: EXAMPLES.md

## ✨ Características Destacadas

- ✅ Sin costo (usa WhatsApp Web)
- ✅ Conexión permanente (no requiere rescanear)
- ✅ Múltiples usuarios simultáneos
- ✅ Sistema de flujos conversacionales
- ✅ Memoria por usuario
- ✅ Respuestas inteligentes
- ✅ Fácil de personalizar
- ✅ Alta performance (Go)
- ✅ Bajo consumo de recursos

## 🎉 ¡Felicitaciones!

Tu bot está listo. Ahora puedes:
- Recibir consultas 24/7
- Agendar citas automáticamente
- Proporcionar información de servicios
- Responder preguntas frecuentes

---

**¿Necesitas ayuda?** Revisa la documentación completa en README.md

**¿Encontraste un bug?** Abre un issue en GitHub

**¿Tienes ideas?** Las pull requests son bienvenidas 🚀

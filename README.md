# 🤖 WhatsApp Bot en Go - Barbería

Bot de WhatsApp creado con Go usando `whatsmeow` (equivalente a Baileys en Node.js). Es un proveedor **gratuito** que se conecta mediante WhatsApp Web escaneando un código QR.

## 📋 Características

- ✅ Conexión gratuita vía WhatsApp Web (QR Code)
- ✅ Sistema de flujos conversacionales
- ✅ Gestión de estados por usuario
- ✅ Agendamiento de citas paso a paso
- ✅ Consulta de servicios y precios
- ✅ Sistema de promociones
- ✅ Persistencia con SQLite
- ✅ Manejo de múltiples usuarios simultáneos

## 🚀 Características del Bot

### Flujos Disponibles

1. **Bienvenida**: Saludo personalizado al usuario
2. **Servicios y Precios**: Lista completa de servicios
3. **Promociones**: Ofertas especiales (2x1, descuentos estudiantes)
4. **Agendamiento**: Sistema paso a paso para agendar citas
   - Captura de nombre
   - Selección de servicio
   - Elección de barbero
   - Fecha deseada
   - Hora preferida
   - Confirmación automática
5. **Horarios**: Información de horarios de atención
6. **Ubicación**: Dirección y estacionamiento

## 📦 Requisitos

- Go 1.21 o superior
- SQLite3

## 🔧 Instalación

### 1. Clonar el repositorio

```bash
git clone <tu-repositorio>
cd whatsapp-bot-go
```

### 2. Instalar dependencias

```bash
go mod download
```

### 3. Ejecutar el bot

```bash
go run .
```

## 🎯 Uso

### Primera vez (Escanear QR)

1. Ejecuta el bot:
```bash
go run .
```

2. Se mostrará un código QR en la terminal

3. Abre WhatsApp en tu teléfono:
   - Android: Menú (⋮) → Dispositivos vinculados → Vincular un dispositivo
   - iPhone: Configuración → Dispositivos vinculados → Vincular un dispositivo

4. Escanea el código QR

5. ¡Listo! El bot está conectado

### Ejecuciones posteriores

El bot recordará la sesión. Solo ejecuta:
```bash
go run .
```

Y se conectará automáticamente sin necesidad de escanear el QR nuevamente.

## 📁 Estructura del Proyecto

```
whatsapp-bot-go/
│
├── main.go          # Punto de entrada, manejo de eventos
├── flows.go         # Flujos conversacionales y estados de usuario
├── utils.go         # Funciones auxiliares
├── go.mod           # Dependencias del proyecto
├── go.sum           # Checksums de dependencias
├── README.md        # Este archivo
└── whatsapp.db      # Base de datos SQLite (se crea automáticamente)
```

## 🔐 Seguridad

- La sesión se guarda localmente en `whatsapp.db`
- **NO compartas este archivo** con nadie
- Añade `whatsapp.db` a tu `.gitignore`

## 📝 Ejemplo de Conversación

```
Usuario: Hola
Bot: ¡Hola Juan! ✂️ Soy el asistente virtual de la Barbería.

Puedo ayudarte a:
📅 Agendar tu cita
💰 Consultar servicios y precios
🎁 Ver promociones

¿En qué puedo asistirte hoy?

Usuario: Quiero agendar una cita
Bot: ¡Perfecto! Vamos a agendar tu cita. 📅

Por favor, dime tu **nombre completo**:

Usuario: Juan Pérez
Bot: Gracias, Juan Pérez. 😊

¿Qué **servicio** deseas?

Algunos ejemplos:
• Corte Tradicional
• Afeitado Tradicional
• Arreglo de Barba
• Combo (especifica cuál)

Usuario: Corte tradicional
Bot: Perfecto. 💈

¿Tienes preferencia de **barbero**?

Opciones:
• Brandon
• Cualquiera

[... y así continúa el flujo]
```

## 🛠️ Personalización

### Modificar Servicios y Precios

Edita la función `getServicesAndPrices()` en `utils.go`

### Agregar Nuevos Flujos

1. Abre `flows.go`
2. Agrega tu lógica en la función `processMessageWithFlows()`
3. Crea nuevas funciones para manejar estados complejos

### Cambiar Horarios

Edita el texto en la sección de horarios en `flows.go`

## 🐛 Solución de Problemas

### El QR no aparece

- Asegúrate de que tienes conexión a internet
- Verifica que el puerto no esté en uso
- Intenta eliminar `whatsapp.db` y volver a escanear

### El bot no responde

- Verifica que WhatsApp Web esté activo
- Revisa los logs en la terminal
- Asegúrate de que el teléfono tenga conexión

### Error de dependencias

```bash
go mod tidy
go mod download
```

## 📚 Dependencias Principales

- `go.mau.fi/whatsmeow` - Cliente de WhatsApp Web
- `github.com/mattn/go-sqlite3` - Base de datos SQLite
- `github.com/mdp/qrterminal/v3` - Generación de QR en terminal
- `golang.org/x/text` - Normalización de texto

## 🚀 Compilar para Producción

### Linux

```bash
GOOS=linux GOARCH=amd64 go build -o whatsapp-bot
./whatsapp-bot
```

### Windows

```bash
GOOS=windows GOARCH=amd64 go build -o whatsapp-bot.exe
```

### Docker (Opcional)

```bash
docker build -t whatsapp-bot .
docker run -it whatsapp-bot
```

## 🔄 Comparación con BuilderBot (Node.js)

| Característica | BuilderBot (Node.js) | Este Bot (Go) |
|---------------|----------------------|---------------|
| Proveedor | Meta (Oficial/Pago) | WhatsApp Web (Gratuito) |
| Autenticación | JWT Token | QR Code |
| Performance | Buena | Excelente |
| Concurrencia | Event Loop | Goroutines |
| Memoria | ~50MB | ~10MB |
| Setup | Complejo | Simple |

## 📈 Próximas Mejoras

- [ ] Integración con Google Calendar
- [ ] Integración con Google Sheets
- [ ] Panel web de administración
- [ ] Notificaciones automáticas
- [ ] Estadísticas de uso
- [ ] Respuestas con botones interactivos
- [ ] Envío de imágenes y documentos
- [ ] Multi-idioma

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Por favor:

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

MIT License - puedes usar este código libremente en tus proyectos.

## 💡 Inspiración

Este proyecto está inspirado en BuilderBot pero implementado desde cero en Go para mayor rendimiento y simplicidad.

## 📞 Soporte

Si tienes preguntas o necesitas ayuda:
- Abre un issue en GitHub
- Revisa la documentación de whatsmeow: https://pkg.go.dev/go.mau.fi/whatsmeow

---

Hecho con ❤️ y ☕ para la comunidad

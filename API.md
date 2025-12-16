# 📡 API Reference

Este documento describe la estructura y funciones principales del bot.

## 📚 Estructura de Archivos

### main.go
Archivo principal que maneja:
- Conexión a WhatsApp Web
- Generación de QR Code
- Manejo de eventos (mensajes, conexión, desconexión)
- Inicialización del bot

**Funciones principales:**
- `main()`: Punto de entrada del programa
- `handleEvents()`: Manejador central de eventos
- `handleMessage()`: Procesa mensajes entrantes

### flows.go
Sistema de flujos conversacionales:
- Gestión de estados por usuario
- Flujo de agendamiento paso a paso
- Almacenamiento en memoria de conversaciones

**Estructuras:**
```go
type UserState struct {
    IsScheduling        bool              // Si está en proceso de agendamiento
    Step                int               // Paso actual del flujo
    Data                map[string]string // Datos recopilados
    ConversationHistory []string          // Historial de mensajes
    LastMessageTime     int64             // Timestamp último mensaje
    AppointmentSaved    bool              // Si la cita fue guardada
}
```

**Funciones principales:**
- `getUserState(userID)`: Obtiene o crea estado de usuario
- `clearUserState(userID)`: Limpia estado de usuario
- `processMessageWithFlows()`: Procesa mensaje según flujo activo

**Tipo AppointmentFlow:**
- `Start()`: Inicia el flujo de agendamiento
- `Continue()`: Continúa con el siguiente paso
- `ConfirmAppointment()`: Confirma y guarda la cita

### utils.go
Funciones auxiliares y utilidades:
- Normalización de texto
- Detección de palabras clave
- Envío de mensajes
- Servicios y promociones

**Funciones principales:**
- `normalizeText(text)`: Quita acentos y convierte a minúsculas
- `isGreeting(message)`: Detecta saludos
- `containsKeywords(message, keywords)`: Busca palabras clave
- `sendMessage(jid, text)`: Envía mensaje de texto
- `getServicesAndPrices()`: Retorna servicios y precios
- `getPromotions()`: Retorna promociones

## 🔄 Flujo de Mensajes

```
Usuario envía mensaje
        ↓
handleMessage() recibe el evento
        ↓
Extrae texto del mensaje
        ↓
processMessageWithFlows() analiza el mensaje
        ↓
┌───────────────────────────────────┐
│ ¿Usuario está agendando?          │
├───────────────────────────────────┤
│ SÍ → AppointmentFlow.Continue()   │
│ NO → Detectar intención           │
└───────────────────────────────────┘
        ↓
Generar respuesta
        ↓
sendMessage() envía respuesta
```

## 🎯 Palabras Clave por Flujo

### Agendamiento
- cita, agendar, turno, reservar, corte, quiero

### Servicios y Precios
- servicio, precio, costo, cuanto

### Promociones
- promocion, descuento, oferta, 2x1

### Horarios
- horario, hora, disponibilidad, cuando

### Ubicación
- ubicacion, direccion, donde, como llegar

## 💾 Persistencia

### Base de Datos SQLite
Archivo: `whatsapp.db`

**Tablas principales:**
- `whatsmeow_device`: Información del dispositivo
- `whatsmeow_identity_keys`: Claves de encriptación
- `whatsmeow_pre_keys`: Claves pre-compartidas
- `whatsmeow_sessions`: Sesiones activas

### Estados en Memoria
Los estados de usuario se almacenan en memoria usando un `map[string]*UserState`

**Ventajas:**
- Rápido acceso
- No requiere base de datos adicional

**Desventajas:**
- Se pierde al reiniciar el bot
- No escala a múltiples instancias

**Mejora futura:** Implementar Redis o base de datos persistente.

## 🔌 Integración con WhatsApp

### Biblioteca: whatsmeow
GitHub: https://github.com/tulir/whatsmeow

**Características:**
- Cliente completo de WhatsApp Web
- Soporte Multi-Device
- Encriptación E2E
- Manejo de mensajes multimedia
- Grupos y broadcasts

### Autenticación
1. Primera vez: Genera QR Code
2. Usuario escanea con WhatsApp
3. Sesión se guarda en `whatsapp.db`
4. Próximas veces: Conexión automática

### Tipos de Mensajes Soportados
- ✅ Texto simple
- ✅ Texto extendido
- ✅ Botones (limitado)
- ⚠️ Imágenes (por implementar)
- ⚠️ Audio (por implementar)
- ⚠️ Documentos (por implementar)

## 🚀 Extender Funcionalidad

### Agregar un Nuevo Flujo

1. **En flows.go**, agregar detección:
```go
if containsKeywords(message, []string{"palabra", "clave"}) {
    return "Tu respuesta aquí"
}
```

2. **Si requiere estado**, crear nueva estructura:
```go
type MiNuevoFlow struct{}

func (f *MiNuevoFlow) Start(userID, message string) string {
    // Lógica de inicio
}
```

3. **Integrar en processMessageWithFlows()**

### Agregar Persistencia a Base de Datos

Ejemplo con SQLite:

```go
import "database/sql"
import _ "github.com/mattn/go-sqlite3"

func saveAppointment(data map[string]string) error {
    db, err := sql.Open("sqlite3", "./appointments.db")
    if err != nil {
        return err
    }
    defer db.Close()
    
    _, err = db.Exec(`
        INSERT INTO appointments (name, service, barber, date, time)
        VALUES (?, ?, ?, ?, ?)
    `, data["nombre"], data["servicio"], data["barbero"], 
       data["fecha"], data["hora"])
    
    return err
}
```

### Agregar Servidor Web/API

```go
import "net/http"

func main() {
    // ... código existente ...
    
    // Agregar servidor HTTP
    go func() {
        http.HandleFunc("/health", healthHandler)
        http.HandleFunc("/appointments", appointmentsHandler)
        http.ListenAndServe(":8080", nil)
    }()
    
    // ... resto del código ...
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status": "healthy"}`))
}
```

## 🔧 Variables de Configuración

Actualmente no se usan variables de entorno, pero puedes agregar:

```go
import "os"

var (
    PORT          = os.Getenv("PORT")          // Puerto del servidor
    BUSINESS_NAME = os.Getenv("BUSINESS_NAME") // Nombre del negocio
    LOCATION      = os.Getenv("LOCATION")      // Ubicación
)
```

## 📊 Monitoreo

### Logs
Todos los eventos se registran en stdout:
- 📨 Mensajes recibidos
- ✅ Respuestas enviadas
- 🟢 Conexiones
- 🔴 Desconexiones
- ❌ Errores

### Métricas Sugeridas
Para producción, considera agregar:
- Número de usuarios activos
- Mensajes procesados por minuto
- Tasa de respuesta
- Errores por tipo
- Tiempo de respuesta promedio

## 🔐 Seguridad

### Buenas Prácticas

1. **Nunca compartas `whatsapp.db`**
   - Contiene tus claves de sesión
   - Acceso completo a tu WhatsApp

2. **Valida entrada de usuarios**
   - Sanitiza textos antes de procesarlos
   - Previene injection attacks

3. **Rate Limiting**
   - Implementa límites de mensajes por usuario
   - Previene spam

4. **Datos Sensibles**
   - No almacenes contraseñas
   - Encripta información personal

## 🎓 Recursos Adicionales

- [Documentación whatsmeow](https://pkg.go.dev/go.mau.fi/whatsmeow)
- [Protocolo WhatsApp](https://github.com/tulir/whatsmeow/blob/main/PROTOCOL.md)
- [BuilderBot (inspiración)](https://builderbot.vercel.app/)

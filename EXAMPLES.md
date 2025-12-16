# 💬 Ejemplos de Conversaciones

Este documento muestra ejemplos reales de cómo interactuar con el bot.

## 📱 Ejemplo 1: Consulta de Servicios

```
Usuario: Hola
Bot: ¡Hola Juan! ✂️ Soy el asistente virtual de la Barbería.

Puedo ayudarte a:
📅 Agendar tu cita
💰 Consultar servicios y precios
🎁 Ver promociones

¿En qué puedo asistirte hoy?

Usuario: ¿Cuánto cuesta un corte?
Bot: 💈 *SERVICIOS Y PRECIOS*

*SERVICIOS INDIVIDUALES:*
✂️ Corte Tradicional - $300
  (Cualquier tipo de corte a tu gusto)

🪒 Afeitado Tradicional - $270
  (Con toallas calientes, máquina y navaja, masaje relajante)

[... lista completa de servicios ...]

¿Te gustaría agendar una cita? 📅
```

## 📅 Ejemplo 2: Agendamiento Completo

```
Usuario: Quiero agendar una cita
Bot: ¡Perfecto! Vamos a agendar tu cita. 📅

Por favor, dime tu **nombre completo**:

Usuario: Carlos Rodríguez
Bot: Gracias, Carlos Rodríguez. 😊

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

Usuario: Brandon
Bot: Excelente. 📅

¿Para qué **fecha** quieres tu cita?

Puedes decirme:
• El día (Ej: Lunes, Martes)
• Una fecha específica (Ej: 20/12/2025)
• Mañana

Usuario: mañana
Bot: Perfecto. ⏰

¿A qué **hora** prefieres?

Nuestro horario es de 9:00 AM a 7:00 PM
(Ejemplo: 10:00 AM, 3:00 PM, 5 de la tarde)

Usuario: 3 de la tarde
Bot: ¡Perfecto! 🎉 Tu cita ha sido agendada exitosamente.

📋 **Resumen de tu cita:**

👤 Nombre: Carlos Rodríguez
✂️ Servicio: Corte tradicional
💈 Barbero: Brandon
📅 Fecha: mañana
⏰ Hora: 3 de la tarde

Te esperamos en la fecha y hora acordada. ¡Gracias por confiar en nosotros! 😊
```

## 🎁 Ejemplo 3: Consulta de Promociones

```
Usuario: ¿Tienen promociones?
Bot: 🎁 *PROMOCIONES ESPECIALES*

📚 *MARTES DE ESTUDIANTES* - $250
   Con credencial vigente

🎉 *MIÉRCOLES 2X1* - $350
   Opciones:
   • Corte + Barba
   • Corte + Mascarilla
   • Barba + Mascarilla

👩 *CORTE MUJERES* - $250
   Todos los días

💡 ¡Aprovecha nuestras promociones y luce increíble!

¿Quieres agendar tu cita? 📅
```

## 🕐 Ejemplo 4: Consulta de Horarios

```
Usuario: ¿A qué hora abren?
Bot: 🕐 Nuestros horarios son:

📅 Lunes a Sábado: 9:00 AM - 7:00 PM
📅 Domingo: 10:00 AM - 5:00 PM

¿Deseas agendar una cita?
```

## 📍 Ejemplo 5: Ubicación

```
Usuario: ¿Dónde están ubicados?
Bot: 📍 Estamos ubicados en:

Calle Principal #123
Colonia Centro
Ciudad, Estado

🅿️ Contamos con estacionamiento exclusivo para clientes.
```

## 🔄 Ejemplo 6: Flujo Interrumpido

```
Usuario: Quiero agendar
Bot: ¡Perfecto! Vamos a agendar tu cita. 📅

Por favor, dime tu **nombre completo**:

Usuario: ¿Cuánto cuesta?
Bot: 💈 *SERVICIOS Y PRECIOS*
[... muestra precios ...]

Usuario: Pedro García
Bot: Gracias, Pedro García. 😊

¿Qué **servicio** deseas?
[... continúa con el flujo ...]
```

## ❌ Ejemplo 7: Mensaje No Entendido

```
Usuario: asdfghjkl
Bot: Lo siento, no entendí tu mensaje. 😅

Puedes preguntarme sobre:
• Servicios y precios 💰
• Promociones 🎁
• Agendar una cita 📅
• Horarios 🕐
• Ubicación 📍
```

## 🎯 Ejemplo 8: Agendamiento Rápido (Todo en un mensaje)

```
Usuario: Hola, quiero agendar un corte con Brandon para mañana a las 10 am, soy Miguel López
Bot: ¡Perfecto! Vamos a agendar tu cita. 📅

Ya tengo algunos datos:
👤 Nombre: Miguel López
💈 Barbero: Brandon
📅 Fecha: mañana
⏰ Hora: 10 am

Solo necesito que me confirmes:
¿Qué **servicio** deseas?

Algunos ejemplos:
• Corte Tradicional
• Afeitado Tradicional
• Arreglo de Barba
• Combo (especifica cuál)

Usuario: Corte tradicional
Bot: ¡Perfecto! 🎉 Tu cita ha sido agendada exitosamente.

[... confirmación completa ...]
```

## 💡 Tips para Usuarios

### Para Agendar Más Rápido
Puedes proporcionar toda la información en un solo mensaje:
```
"Hola, quiero un corte tradicional con Brandon para el viernes a las 4 pm, soy Juan Pérez"
```

### Formas de Expresar Fecha
- "hoy"
- "mañana"
- "pasado mañana"
- "lunes", "martes", etc.
- "20/12/2024"
- "el viernes"

### Formas de Expresar Hora
- "10:00 AM"
- "3:00 PM"
- "10 am"
- "3 pm"
- "10 de la mañana"
- "3 de la tarde"
- "5 de la tarde"

### Servicios Más Solicitados
1. Corte Tradicional ($300)
2. Combo Corte + Afeitado Express ($450)
3. Arreglo de Barba ($220)
4. Combo Corte + Afeitado Tradicional ($500)

### Promociones del Día
- **Martes**: Estudiantes pagan solo $250 (con credencial)
- **Miércoles**: 2x1 en combos por $350
- **Todos los días**: Cortes para mujeres a $250

## 🤖 Comportamiento del Bot

### Qué Hace el Bot
✅ Responde inmediatamente
✅ Recuerda el contexto durante la conversación
✅ Extrae información de mensajes largos
✅ Normaliza formatos de fecha y hora
✅ Confirma citas automáticamente
✅ Proporciona información detallada

### Qué NO Hace el Bot
❌ No responde en grupos (configurable)
❌ No procesa imágenes aún
❌ No cancela citas (por implementar)
❌ No envía recordatorios (por implementar)
❌ No procesa pagos

## 🔧 Personalizaciones Disponibles

### Para Desarrolladores
Puedes modificar:
- Servicios y precios en `config.go`
- Mensajes del bot en `utils.go`
- Flujo de agendamiento en `flows.go`
- Palabras clave detectadas en `flows.go`

### Agregar Nuevos Flujos
Ejemplo para agregar "Consulta de Disponibilidad":

```go
// En flows.go, dentro de processMessageWithFlows()

if containsKeywords(message, []string{"disponible", "horarios libres"}) {
    return getAvailableSlots()
}

// Nueva función
func getAvailableSlots() string {
    return "🗓️ Horarios disponibles hoy:\n\n" +
           "• 10:00 AM\n" +
           "• 2:00 PM\n" +
           "• 5:00 PM\n\n" +
           "¿Quieres agendar?"
}
```

## 📊 Estadísticas de Uso

Con el bot puedes rastrear:
- Mensajes recibidos por día
- Citas agendadas exitosamente
- Servicios más solicitados
- Horarios más populares
- Tasa de conversión (mensajes → citas)

## 🚀 Próximas Funcionalidades

Planeadas para implementar:
- [ ] Cancelación de citas
- [ ] Reprogramación de citas
- [ ] Recordatorios automáticos 24h antes
- [ ] Confirmación de asistencia
- [ ] Envío de ubicación con Google Maps
- [ ] Galería de trabajos (imágenes)
- [ ] Sistema de reseñas
- [ ] Programa de lealtad
- [ ] Integración con calendario de barberos

---

¿Tienes más ideas? ¡Compártelas en los issues de GitHub!

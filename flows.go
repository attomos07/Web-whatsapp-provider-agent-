package main

import (
	"fmt"
	"sync"
)

// Estado del usuario
type UserState struct {
	IsScheduling        bool
	Step                int
	Data                map[string]string
	ConversationHistory []string
	LastMessageTime     int64
	AppointmentSaved    bool
}

// Almacén de estados de usuario (en memoria)
var (
	userStates = make(map[string]*UserState)
	stateMutex sync.RWMutex
)

// Obtener o crear estado del usuario
func getUserState(userID string) *UserState {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	if state, exists := userStates[userID]; exists {
		return state
	}

	state := &UserState{
		IsScheduling:        false,
		Step:                0,
		Data:                make(map[string]string),
		ConversationHistory: []string{},
		LastMessageTime:     0,
		AppointmentSaved:    false,
	}

	userStates[userID] = state
	return state
}

// Limpiar estado del usuario
func clearUserState(userID string) {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	delete(userStates, userID)
}

// Flujo de agendamiento
type AppointmentFlow struct{}

// Iniciar flujo de agendamiento
func (f *AppointmentFlow) Start(userID, message string) string {
	state := getUserState(userID)
	state.IsScheduling = true
	state.Step = 1

	// Agregar mensaje al historial
	state.ConversationHistory = append(state.ConversationHistory, "Usuario: "+message)

	return "¡Perfecto! Vamos a agendar tu cita. 📅\n\n" +
		"Por favor, dime tu **nombre completo**:"
}

// Continuar con el flujo de agendamiento
func (f *AppointmentFlow) Continue(userID, message string) string {
	state := getUserState(userID)

	// Agregar al historial
	state.ConversationHistory = append(state.ConversationHistory, "Usuario: "+message)

	// Extraer información según el paso actual
	switch state.Step {
	case 1: // Recopilar nombre
		if !hasKey(state.Data, "nombre") {
			state.Data["nombre"] = message
			state.Step = 2
			return fmt.Sprintf("Gracias, %s. 😊\n\n¿Qué **servicio** deseas?\n\n"+
				"Algunos ejemplos:\n"+
				"• Corte Tradicional\n"+
				"• Afeitado Tradicional\n"+
				"• Arreglo de Barba\n"+
				"• Combo (especifica cuál)", state.Data["nombre"])
		}

	case 2: // Recopilar servicio
		if !hasKey(state.Data, "servicio") {
			state.Data["servicio"] = message
			state.Step = 3
			return "Perfecto. 💈\n\n¿Tienes preferencia de **barbero**?\n\n" +
				"Opciones:\n• Brandon\n• Cualquiera"
		}

	case 3: // Recopilar barbero
		if !hasKey(state.Data, "barbero") {
			if containsKeywords(normalizeText(message), []string{"cualquier", "no", "da igual", "me da igual"}) {
				state.Data["barbero"] = "Cualquiera"
			} else {
				state.Data["barbero"] = message
			}
			state.Step = 4
			return "Excelente. 📅\n\n¿Para qué **fecha** quieres tu cita?\n\n" +
				"Puedes decirme:\n• El día (Ej: Lunes, Martes)\n• Una fecha específica (Ej: 20/12/2025)\n• Mañana"
		}

	case 4: // Recopilar fecha
		if !hasKey(state.Data, "fecha") {
			state.Data["fecha"] = message
			state.Step = 5
			return "Perfecto. ⏰\n\n¿A qué **hora** prefieres?\n\n" +
				"Nuestro horario es de 9:00 AM a 7:00 PM\n" +
				"(Ejemplo: 10:00 AM, 3:00 PM, 5 de la tarde)"
		}

	case 5: // Recopilar hora y confirmar
		if !hasKey(state.Data, "hora") {
			state.Data["hora"] = message
			state.Step = 6

			// Todos los datos recopilados, confirmar
			return f.ConfirmAppointment(state)
		}
	}

	return "Por favor proporciona la información solicitada. 🙏"
}

// Confirmar la cita
func (f *AppointmentFlow) ConfirmAppointment(state *UserState) string {
	// Marcar como guardada
	state.AppointmentSaved = true
	state.IsScheduling = false

	confirmation := "¡Perfecto! 🎉 Tu cita ha sido agendada exitosamente.\n\n"
	confirmation += "📋 **Resumen de tu cita:**\n\n"
	confirmation += fmt.Sprintf("👤 Nombre: %s\n", state.Data["nombre"])
	confirmation += fmt.Sprintf("✂️ Servicio: %s\n", state.Data["servicio"])
	confirmation += fmt.Sprintf("💈 Barbero: %s\n", state.Data["barbero"])
	confirmation += fmt.Sprintf("📅 Fecha: %s\n", state.Data["fecha"])
	confirmation += fmt.Sprintf("⏰ Hora: %s\n\n", state.Data["hora"])
	confirmation += "Te esperamos en la fecha y hora acordada. ¡Gracias por confiar en nosotros! 😊"

	// Aquí podrías guardar en una base de datos o Google Sheets
	// saveAppointmentToDB(state.Data)

	return confirmation
}

// Verificar si existe una clave en el mapa
func hasKey(m map[string]string, key string) bool {
	_, exists := m[key]
	return exists
}

// Procesar mensaje mejorado con flujos
func processMessageWithFlows(message, phone, name string) string {
	state := getUserState(phone)
	message = normalizeText(message)

	// Si está en proceso de agendamiento, continuar con el flujo
	if state.IsScheduling && !state.AppointmentSaved {
		flow := &AppointmentFlow{}
		return flow.Continue(phone, message)
	}

	// Si acaba de guardar una cita, reiniciar
	if state.AppointmentSaved {
		clearUserState(phone)
		getUserState(phone)
	}

	// Detectar intención de agendamiento
	if containsKeywords(message, []string{"cita", "agendar", "turno", "reservar", "corte", "quiero"}) {
		flow := &AppointmentFlow{}
		return flow.Start(phone, message)
	}

	// Flujos normales (como antes)
	if isGreeting(message) {
		return fmt.Sprintf("¡Hola %s! ✂️ Soy el asistente virtual de la Barbería.\n\n"+
			"Puedo ayudarte a:\n"+
			"📅 Agendar tu cita\n"+
			"💰 Consultar servicios y precios\n"+
			"🎁 Ver promociones\n\n"+
			"¿En qué puedo asistirte hoy?", name)
	}

	if containsKeywords(message, []string{"servicio", "precio", "costo", "cuanto"}) {
		return getServicesAndPrices()
	}

	if containsKeywords(message, []string{"promocion", "descuento", "oferta", "2x1"}) {
		return getPromotions()
	}

	if containsKeywords(message, []string{"horario", "hora", "disponibilidad", "cuando"}) {
		return "🕐 Nuestros horarios son:\n\n" +
			"📅 Lunes a Sábado: 9:00 AM - 7:00 PM\n" +
			"📅 Domingo: 10:00 AM - 5:00 PM\n\n" +
			"¿Deseas agendar una cita?"
	}

	if containsKeywords(message, []string{"ubicacion", "direccion", "donde", "como llegar"}) {
		return "📍 Estamos ubicados en:\n\n" +
			"Calle Principal #123\n" +
			"Colonia Centro\n" +
			"Ciudad, Estado\n\n" +
			"🅿️ Contamos con estacionamiento exclusivo para clientes."
	}

	return "Lo siento, no entendí tu mensaje. 😅\n\n" +
		"Puedes preguntarme sobre:\n" +
		"• Servicios y precios 💰\n" +
		"• Promociones 🎁\n" +
		"• Agendar una cita 📅\n" +
		"• Horarios 🕐\n" +
		"• Ubicación 📍"
}

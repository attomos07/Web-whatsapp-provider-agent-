package main

// Configuración del bot
// Puedes modificar estos valores según tus necesidades

const (
	// Nombre del negocio
	BUSINESS_NAME = "Barbería Moderna"
	
	// Ubicación
	BUSINESS_ADDRESS = `Calle Principal #123
Colonia Centro
Ciudad, Estado
CP 12345`
	
	// Horarios de atención
	SCHEDULE_WEEKDAY = "Lunes a Sábado: 9:00 AM - 7:00 PM"
	SCHEDULE_SUNDAY  = "Domingo: 10:00 AM - 5:00 PM"
	
	// Teléfono de contacto (opcional)
	CONTACT_PHONE = "+52 1 234 567 8900"
	
	// Barberos disponibles
	BARBERS = "Brandon, Carlos, Miguel"
	
	// Duración promedio de servicios (minutos)
	SERVICE_DURATION_SHORT  = 30  // Arreglo de barba
	SERVICE_DURATION_MEDIUM = 45  // Corte
	SERVICE_DURATION_LONG   = 60  // Corte + Afeitado
	
	// Mensajes del bot
	WELCOME_MESSAGE = "¡Hola! ✂️ Soy el asistente virtual de la Barbería."
	
	// Base de datos
	DATABASE_FILE = "whatsapp.db"
	
	// Configuración de logs
	LOG_LEVEL = "INFO" // DEBUG, INFO, WARN, ERROR
)

// Servicios disponibles
var SERVICES = map[string]int{
	"Corte Tradicional":                           300,
	"Afeitado Tradicional":                        270,
	"Afeitado Express":                            270,
	"Arreglo de Barba":                            220,
	"Mascarillas":                                 250,
	"Combo Corte + Afeitado Express":              450,
	"Combo Corte + Afeitado Tradicional":          500,
	"Combo Corte + Arreglo":                       420,
	"Combo Corte + Afeitado Tradicional + Mascarilla": 700,
}

// Promociones activas
type Promotion struct {
	Name        string
	Description string
	Price       int
	Days        []string // Días aplicables
}

var PROMOTIONS = []Promotion{
	{
		Name:        "Martes de Estudiantes",
		Description: "Corte tradicional con credencial vigente",
		Price:       250,
		Days:        []string{"martes"},
	},
	{
		Name:        "Miércoles 2x1",
		Description: "Corte+Barba, Corte+Mascarilla, o Barba+Mascarilla",
		Price:       350,
		Days:        []string{"miércoles", "miercoles"},
	},
	{
		Name:        "Corte Mujeres",
		Description: "Todos los días",
		Price:       250,
		Days:        []string{"todos"},
	},
}

// Configuración de respuestas automáticas
var AUTO_REPLIES = map[string]string{
	"gracias":     "¡De nada! Estoy aquí para ayudarte. 😊",
	"adios":       "¡Hasta pronto! Te esperamos. 👋",
	"ok":          "Perfecto, ¿hay algo más en lo que pueda ayudarte?",
	"no":          "Entendido. Si necesitas algo más, aquí estoy. 😊",
	"si":          "¡Excelente! ¿En qué más puedo ayudarte?",
}

// Mensajes de error
const (
	ERROR_GENERIC           = "Lo siento, ocurrió un error. Por favor intenta nuevamente."
	ERROR_INVALID_DATE      = "La fecha que proporcionaste no es válida. Intenta con formato: Lunes, 20/12/2024, o 'mañana'."
	ERROR_INVALID_TIME      = "La hora no es válida. Nuestro horario es de 9:00 AM a 7:00 PM."
	ERROR_MISSING_DATA      = "Parece que falta información. ¿Podrías proporcionarla?"
	ERROR_NOT_AVAILABLE     = "Lo siento, ese horario no está disponible. ¿Tienes otra opción?"
)

// Configuración de grupos (si quieres habilitar respuestas en grupos)
const (
	RESPOND_TO_GROUPS = false // Cambiar a true para responder en grupos
	MAX_GROUP_MEMBERS = 50    // Máximo de miembros en grupo para responder
)

// Límites de rate limiting (mensajes por minuto por usuario)
const (
	RATE_LIMIT_ENABLED  = true
	RATE_LIMIT_MESSAGES = 10  // Mensajes permitidos
	RATE_LIMIT_WINDOW   = 60  // Ventana en segundos
)

// Funciones helper para configuración

// GetServicePrice retorna el precio de un servicio
func GetServicePrice(serviceName string) (int, bool) {
	price, exists := SERVICES[serviceName]
	return price, exists
}

// IsPromotionActive verifica si una promoción está activa en un día específico
func IsPromotionActive(promotionName, day string) bool {
	for _, promo := range PROMOTIONS {
		if promo.Name == promotionName {
			for _, promoDay := range promo.Days {
				if promoDay == day || promoDay == "todos" {
					return true
				}
			}
		}
	}
	return false
}

// GetAutoReply obtiene una respuesta automática si existe
func GetAutoReply(keyword string) (string, bool) {
	reply, exists := AUTO_REPLIES[normalizeText(keyword)]
	return reply, exists
}

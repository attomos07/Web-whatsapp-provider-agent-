package main

import (
	"context"
	"fmt"
	"strings"
	"unicode"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"google.golang.org/protobuf/proto"

	waProto "go.mau.fi/whatsmeow/binary/proto"
)

var client *whatsmeow.Client

// SetClient configura el cliente global
func SetClient(c *whatsmeow.Client) {
	client = c
}

// Normalizar texto (quitar acentos, minúsculas)
func normalizeText(text string) string {
	// Convertir a minúsculas
	text = strings.ToLower(text)

	// Remover acentos
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, text)

	return strings.TrimSpace(result)
}

// Verificar si el mensaje es un saludo
func isGreeting(message string) bool {
	greetings := []string{
		"hola", "buenos dias", "buenas tardes", "buenas noches",
		"hey", "hola!", "hi", "hello", "saludos", "que tal",
	}

	for _, greeting := range greetings {
		if strings.Contains(message, greeting) {
			return true
		}
	}

	return false
}

// Verificar si el mensaje contiene alguna palabra clave
func containsKeywords(message string, keywords []string) bool {
	for _, keyword := range keywords {
		if strings.Contains(message, keyword) {
			return true
		}
	}
	return false
}

// Enviar mensaje de texto
func sendMessage(jid types.JID, text string) error {
	if client == nil {
		return fmt.Errorf("cliente no configurado")
	}

	msg := &waProto.Message{
		Conversation: proto.String(text),
	}

	_, err := client.SendMessage(context.Background(), jid, msg)
	return err
}

// Obtener servicios y precios
func getServicesAndPrices() string {
	return `💈 *SERVICIOS Y PRECIOS*

*SERVICIOS INDIVIDUALES:*
✂️ Corte Tradicional - $300
  (Cualquier tipo de corte a tu gusto)

🪒 Afeitado Tradicional - $270
  (Con toallas calientes, máquina y navaja, masaje relajante)

🪒 Afeitado Express - $270
  (Rebajada con máquina y tierra, limpieza con navaja)

🧔 Arreglo de Barba - $220
  (Limpieza con navaja o tijera del contorno)

😷 Mascarillas - $250
  (Negra o de barro)

*COMBOS ESPECIALES:*
💰 Corte + Afeitado Express - $450
💰 Corte + Afeitado Tradicional - $500
💰 Corte + Arreglo - $420
💰 Corte + Afeitado Tradicional + Mascarilla - $700

*EXTRAS:*
🅿️ Estacionamiento exclusivo disponible

¿Te gustaría agendar una cita? 📅`
}

// Obtener promociones
func getPromotions() string {
	return `🎁 *PROMOCIONES ESPECIALES*

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

¿Quieres agendar tu cita? 📅`
}

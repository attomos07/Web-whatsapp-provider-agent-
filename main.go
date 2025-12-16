package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func main() {
	// Configurar logger
	dbLog := waLog.Stdout("Database", "INFO", true)

	// Crear contexto
	ctx := context.Background()

	// Crear contenedor de store SQLite
	container, err := sqlstore.New(ctx, "sqlite3", "file:whatsapp.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}

	// Si no hay dispositivos, crear uno nuevo
	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		panic(err)
	}

	clientLog := waLog.Stdout("Client", "INFO", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)

	// Configurar cliente global para envío de mensajes
	SetClient(client)

	// Registrar manejador de eventos
	client.AddEventHandler(handleEvents)

	// Si no está conectado, mostrar QR
	if client.Store.ID == nil {
		// Sin ID, muestra QR
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}

		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("🔐 Escanea este código QR con tu WhatsApp:")
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("📱 Estado de login:", evt.Event)
			}
		}
	} else {
		// Ya está conectado
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("✅ Bot conectado exitosamente!")
	fmt.Println("📱 Esperando mensajes...")

	// Mantener el programa corriendo
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	fmt.Println("\n👋 Desconectando bot...")
	client.Disconnect()
}

// Manejador de eventos
func handleEvents(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		handleMessage(v)
	case *events.Receipt:
		// Confirmaciones de lectura
		if v.Type == events.ReceiptTypeRead || v.Type == events.ReceiptTypeReadSelf {
			fmt.Printf("✓✓ Mensaje leído: %s\n", v.MessageIDs[0])
		}
	case *events.Connected:
		fmt.Println("🟢 Conectado a WhatsApp")
	case *events.Disconnected:
		fmt.Println("🔴 Desconectado de WhatsApp")
	}
}

// Manejar mensajes entrantes
func handleMessage(msg *events.Message) {
	// Ignorar mensajes propios
	if msg.Info.IsFromMe {
		return
	}

	// Ignorar mensajes de grupos (opcional)
	if msg.Info.IsGroup {
		return
	}

	sender := msg.Info.Sender.User
	senderName := msg.Info.PushName

	// Obtener el texto del mensaje
	var messageText string
	if msg.Message.GetConversation() != "" {
		messageText = msg.Message.GetConversation()
	} else if msg.Message.GetExtendedTextMessage() != nil {
		messageText = msg.Message.GetExtendedTextMessage().GetText()
	}

	fmt.Printf("📨 Mensaje de %s (%s): %s\n", senderName, sender, messageText)

	// Procesar mensaje con el flujo del bot mejorado
	response := processMessageWithFlows(messageText, sender, senderName)

	// Enviar respuesta
	if response != "" {
		err := sendMessage(msg.Info.Chat, response)
		if err != nil {
			fmt.Printf("❌ Error enviando mensaje: %v\n", err)
		} else {
			fmt.Printf("✅ Respuesta enviada a %s\n", senderName)
		}
	}
}

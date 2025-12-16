# Build stage
FROM golang:1.21-alpine AS builder

# Instalar dependencias necesarias
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar el binario
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o whatsapp-bot .

# Runtime stage
FROM alpine:latest

# Instalar dependencias de runtime
RUN apk --no-cache add ca-certificates sqlite-libs

WORKDIR /root/

# Copiar el binario desde el build stage
COPY --from=builder /app/whatsapp-bot .

# Crear volumen para persistir la sesión
VOLUME ["/root"]

# Exponer puerto (si decides agregar un servidor web)
EXPOSE 3000

# Ejecutar el bot
CMD ["./whatsapp-bot"]

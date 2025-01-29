# 📄 API en Go con Docker y Docker Compose

Una API simple en Go contenerizada con Docker y gestionada con Docker Compose. Este proyecto está orientado a demostrar cómo contenerizar una aplicación básica en Go y gestionarla a través de Docker y Docker Compose, permitiendo una implementación flexible y escalable.

## 📋 Descripción
API en Go con dos endpoints:
- /ping: Responde con "pong"
- /status: Muestra el número de visitas y configuración del puerto

Características principales:
- Empaquetado en contenedor Docker
- Configuración mediante variables de entorno
- Contador de visitas en memoria
- Gestión con Docker Compose

## 🛠 Tecnologías utilizadas
- **Go**: Lenguaje principal de desarrollo
- **Docker**: Contenerización de la aplicación
- **Docker Compose**: Orquestación de contenedores
- **Variables de entorno**: Configuración del puerto

## 📁 Estructura del proyecto

go-docker-api/
├── Dockerfile
├── docker-compose.yml
├── main.go
└── go.mod



## 📄 Detalles del código

### main.go
- Define los endpoints:
- 
  
go
  // Endpoint /ping
  http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprint(w, "pong")
  })
  
  // Endpoint /status
  http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
      visitCount++
      fmt.Fprintf(w, "Visitas: %d\nPuerto: %s", visitCount, port)
  })

### Dockerfile
  
  
FROM golang:1.21-alpine
  WORKDIR /app
  COPY go.mod ./
  RUN go mod download
  COPY *.go ./
  RUN CGO_ENABLED=0 GOOS=linux go build -o /go-docker-api
  EXPOSE 8080
  CMD ["/go-docker-api"]

### docker-compose.yml
  
  
version: '3.8'
services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - PORT=8080



## 🚀 Instrucciones de ejecución

### Construir la imagen
  
bash
  docker-compose build

### Iniciar el contenedor
  
bash
  docker-compose up

### Acceder a la API
  
bash
  # Endpoint /ping
  curl http://localhost:8080/ping
  
  # Endpoint /status
  curl http://localhost:8080/status

### Detener el contenedor
  
bash
docker-compose down

## ✅ Lo que aprendí
- **Go**: Lenguaje principal de desarrollo
- **Docker**: Contenerización de la aplicación
- **Docker Compose**: Orquestación de contenedores
- **Variables de entorno**: Configuración del puerto

## 🔜 Próximos pasos
- Agregar persistencia de datos (SQLite/MySQL)
- Implementar sistema de logs
- Añadir validación de datos
- Mejorar el manejo de errores
- Crear rutas más complejas
- Agregar autenticación
- Implementar pruebas unitarias

# Usar la imagen de Go
FROM golang:1.21

# Establecer directorio de trabajo en el contenedor
WORKDIR /app

# Copiar archivos al contenedor
COPY . .

# Descargar dependencias
RUN go mod tidy

# Compilar el binario
RUN go build -o main .

# Especificar puerto de ejecución
EXPOSE 8080

# Definir comando de arranque
CMD ["./main"]

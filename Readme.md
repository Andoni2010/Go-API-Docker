<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>API en Go con Docker y Docker Compose</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            background-color: #f4f4f4;
            color: #333;
            margin: 0;
            padding: 20px;
        }
        h1, h2, h3 {
            color: #333;
        }
        pre {
            background-color: #2e2e2e;
            color: #f5f5f5;
            padding: 15px;
            border-radius: 5px;
            overflow-x: auto;
        }
        code {
            color: #ff9e00;
        }
        ul {
            list-style-type: disc;
            margin-left: 20px;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
            background-color: #fff;
            padding: 20px;
            border-radius: 8px;
        }
        .highlight {
            color: #ff9e00;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>📄 API en Go con Docker y Docker Compose</h1>
        <p>Una API simple en Go contenerizada con Docker y gestionada con Docker Compose. Este proyecto está orientado a demostrar cómo contenerizar una aplicación básica en Go y gestionarla a través de Docker y Docker Compose, permitiendo una implementación flexible y escalable.</p>

        <h2>📋 Descripción</h2>
        <p>La API en Go cuenta con dos endpoints:</p>
        <ul>
            <li><code>/ping</code>: Responde con "pong".</li>
            <li><code>/status</code>: Muestra el número de visitas y la configuración del puerto.</li>
        </ul>
        <p><strong>Características principales:</strong></p>
        <ul>
            <li>Empaquetado en contenedor Docker.</li>
            <li>Configuración mediante variables de entorno.</li>
            <li>Contador de visitas en memoria (sin persistencia).</li>
            <li>Gestión de contenedores con Docker Compose.</li>
        </ul>

        <h2>🛠 Tecnologías utilizadas</h2>
        <ul>
            <li><strong>Go</strong>: Lenguaje principal de desarrollo de la API.</li>
            <li><strong>Docker</strong>: Contenerización de la aplicación y sus dependencias.</li>
            <li><strong>Docker Compose</strong>: Orquestación de contenedores para facilitar el despliegue y gestión de servicios.</li>
            <li><strong>Variables de entorno</strong>: Configuración del puerto de la aplicación.</li>
        </ul>

        <h2>📁 Estructura del proyecto</h2>
        <pre>
go-docker-api/
├── Dockerfile
├── docker-compose.yml
├── main.go
└── go.mod
        </pre>

        <h2>📄 Detalles del código</h2>

        <h3>main.go</h3>
        <p>Define dos endpoints simples en Go:</p>
        <pre>
        // Endpoint /ping
        http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprint(w, "pong")
        })

        // Endpoint /status
        http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
            visitCount++
            fmt.Fprintf(w, "Visitas: %d\nPuerto: %s", visitCount, port)
        })
        </pre>

        <h3>Dockerfile</h3>
        <p>Define la construcción de la imagen Docker para la API:</p>
        <pre>
        FROM golang:1.21-alpine
        WORKDIR /app
        COPY go.mod ./
        RUN go mod download
        COPY *.go ./
        RUN CGO_ENABLED=0 GOOS=linux go build -o /go-docker-api
        EXPOSE 8080
        CMD ["/go-docker-api"]
        </pre>

        <h3>docker-compose.yml</h3>
        <p>Configura el servicio de la API con Docker Compose:</p>
        <pre>
        version: '3.8'
        services:
          api:
            build: .
            ports:
              - "8080:8080"
            environment:
              - PORT=8080
        </pre>

        <h2>🚀 Instrucciones de ejecución</h2>
        <h3>Construir la imagen</h3>
        <pre><code>docker-compose build</code></pre>

        <h3>Iniciar el contenedor</h3>
        <pre><code>docker-compose up</code></pre>

        <h3>Acceder a la API</h3>
        <p>Prueba los siguientes endpoints:</p>
        <pre><code># Endpoint /ping</code>
        curl http://localhost:8080/ping
        <code># Endpoint /status</code>
        curl http://localhost:8080/status
        </pre>

        <h3>Detener el contenedor</h3>
        <pre><code>docker-compose down</code></pre>

        <h2>✅ Lo que aprendí</h2>
        <ul>
            <li>Contenerización de aplicaciones en Go con Docker.</li>
            <li>Uso de Docker Compose para la gestión de contenedores.</li>
            <li>Configuración de variables de entorno para el despliegue de la API.</li>
            <li>Optimización y gestión de la infraestructura de contenedores.</li>
        </ul>

        <h2>🔜 Próximos pasos</h2>
        <ul>
            <li>Agregar persistencia de datos (SQLite/MySQL).</li>
            <li>Implementar sistema de logs.</li>
            <li>Añadir validación de datos en los endpoints.</li>
            <li>Mejorar el manejo de errores con respuestas más detalladas.</li>
            <li>Crear rutas más complejas para la API.</li>
            <li>Agregar autenticación con tokens o sesiones.</li>
            <li>Implementar pruebas unitarias y de integración para los endpoints.</li>
        </ul>
    </div>
</body>
</html>

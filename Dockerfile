# Utiliza una imagen oficial de Golang como base
FROM golang:1.22.1-alpine as builder

ARG REPO_URL
ARG ACCESS_TOKEN
ARG BRANCH_REPO

ENV GO111MODULE=on

# Instala Git
RUN apk add --no-cache git

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Configura el token de acceso personal para clonar el repositorio
RUN git config --global url."https://${ACCESS_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
# Descarga tu repositorio desde la URL proporcionada como argumento
RUN git clone -b $BRANCH_REPO $REPO_URL /app/repo_remote

# Cambia al directorio del repositorio clonado
WORKDIR /app/repo_remote

# Descarga las dependencias
RUN go mod tidy

# Compila tu aplicación Go
RUN GOOS=linux go build -o /app/myapp

# Inicia una nueva etapa utilizando una imagen ligera de alpine como base
FROM alpine:latest

RUN apk --no-cache add tzdata
ENV TZ=America/Santiago

RUN addgroup -S app -g 1000 &&  adduser -S -g app app --uid 1000

COPY --from=builder --chown=app:app /app/myapp /app/myapp

USER app
WORKDIR /app

RUN pwd && ls -la && ldd /app/myapp || true

EXPOSE 8080

# Ejecuta tu aplicación cuando se inicie el contenedor
ENTRYPOINT ["./myapp"]

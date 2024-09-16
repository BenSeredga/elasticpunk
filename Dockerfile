# Используем официальный образ Go
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код приложения
COPY . .

# Сборка приложения
RUN GOOS=linux GOARCH=amd64 go build -o restApi .

RUN chmod +x ./restApi

# Открываем порт 8080
EXPOSE 8080



# Запускаем приложение
CMD ["./restApi"]
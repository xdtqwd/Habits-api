# базовый образ с Go
FROM golang:1.24

# рабочая папка внутри контейнера
WORKDIR /app

# копируем файлы модуля
COPY go.mod go.sum ./

# скачиваем зависимости
RUN go mod download

# копируем весь код
COPY . .

# собираем приложение
RUN go build -o main .

# запускаем
CMD ["./main"]
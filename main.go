package main

import (
	"fmt"
)

func main() {
	InitDB()
	LoadHabits()
	fmt.Println("Сервер запущен на порту 8080")
	HttpServer()
}

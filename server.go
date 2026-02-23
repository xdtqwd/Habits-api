package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HttpServer() {
	http.HandleFunc("/habits", handlerHabits)
	http.HandleFunc("/habits/", handlerHabit)
	http.ListenAndServe(":8080", nil)

}
func handlerHabits(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// текущий код который возвращает список
		data, err := json.Marshal(GetHabits())
		if err != nil {
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	case "POST":
		var newHabit Task
		err := json.NewDecoder(r.Body).Decode(&newHabit)
		if err != nil {
			http.Error(w, "Неверный формат данных", http.StatusBadRequest)
			return
		}
		newHabit.CreatedAt = time.Now()
		SaveHabitDB(newHabit)
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
func handlerHabit(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "DELETE":
		DeleteHabitDB(id)
		w.WriteHeader(http.StatusNoContent)
	case "PUT":
		CompleteHabitDB(id)
		w.WriteHeader(http.StatusOK)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int
	Done        bool
	Name        string
	CreatedAt   time.Time
	CompletedAt *time.Time
}

var habits []Task

func ListHabits() {
	for _, habit := range habits {
		fmt.Printf("ID: %d, Name: %s, Done: %v\n", habit.ID, habit.Name, habit.Done)
	}
}

func SaveHabits() {
	data, err := json.Marshal(habits)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	err = os.WriteFile("habits.json", data, 0644)
	if err != nil {
		fmt.Println("Ошибка записи файла:", err)
	}
}
func LoadHabits() {
	data, err := os.ReadFile("habits.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	err = json.Unmarshal(data, &habits)
	if err != nil {
		fmt.Println("Ошибка сервера:", err)

	}
}

func DeleteHabit(id int) {
	for i, habit := range habits {
		if habit.ID == id {
			habits = append(habits[:i], habits[i+1:]...)
			break
		}
	}
}
func CreateHabit(name string) Task {
	return Task{
		Name:      name,
		CreatedAt: time.Now(),
		ID:        1,
	}
}

func AddHabit(name string) {
	habit := CreateHabit(name)
	habit.ID = len(habits) + 1
	habits = append(habits, habit)
}

func CompleteHabit(id int) {
	for i, habit := range habits {
		if habit.ID == id {
			habits[i].Done = true
			now := time.Now()
			habits[i].CompletedAt = &now
			break
		}
	}
}

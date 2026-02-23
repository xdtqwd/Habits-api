package main

import (
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite", "/app/data/habits.db")
	if err != nil {
		panic(err)
	}

	// создаём таблицу если её нет
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS habits (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        done BOOLEAN,
        created_at DATETIME,
        completed_at DATETIME
    )`)
	if err != nil {
		panic(err)
	}
}

func GetHabits() []Task {
	rows, err := db.Query("SELECT id, name, done, created_at FROM habits")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var result []Task
	for rows.Next() {
		var t Task
		rows.Scan(&t.ID, &t.Name, &t.Done, &t.CreatedAt)
		result = append(result, t)
	}
	return result
}

func SaveHabitDB(habit Task) {
	if habit.ID == 0 {
		db.Exec("INSERT INTO habits (name, done, created_at) VALUES (?, ?, ?)", habit.Name, habit.Done, habit.CreatedAt)
	} else {
		db.Exec("UPDATE habits SET name = ?, done = ?, completed_at = ? WHERE id = ?", habit.Name, habit.Done, habit.CompletedAt, habit.ID)
	}
}

func DeleteHabitDB(id int) {
	db.Exec("DELETE FROM habits WHERE id = ?", id)
}
func CompleteHabitDB(id int) {
	db.Exec("UPDATE habits SET done = ?, completed_at = ? WHERE id = ?", true, time.Now(), id)
}

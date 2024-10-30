package main

import (
	"awesomeProject/internal/app/ds"
	"awesomeProject/internal/app/dsn"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()

	env := dsn.FromEnv()
	fmt.Println("!   !   !   DB Connection String:", env)

	db, err := gorm.Open(postgres.Open(env), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	if err := db.AutoMigrate(
		&ds.TaskItem{},
		&ds.LessonRequest{},
		&ds.User{},
		&ds.TaskLesson{},
	); err != nil {
		fmt.Println("Migration error:", err)
		panic("cant migrate db")
	}
}

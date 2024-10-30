package dsn

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// FromEnv собирает DSN строку из переменных окружения
func FromEnv() string {
	if err := godotenv.Load(); err != nil {
		return ""
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		return ""
	}
	print("FromEnv")

	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
}

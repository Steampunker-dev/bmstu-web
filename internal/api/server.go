package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Hours       int
	Image       string
}

func StartServer() {
	log.Println("Server start up")
	// Данные для карточек
	tasks := []Task{
		{ID: 1, Title: "РК 1 модуль", Description: "Рубежный контроль за первый модуль семестра 2024", Hours: 5, Image: "https://via.placeholder.com/150"},
		{ID: 2, Title: "РК 2 модуль", Description: "Рубежный контроль за второй модуль семестра 2024", Hours: 5, Image: "https://via.placeholder.com/150"},
		{ID: 3, Title: "Подготовка к экзамену", Description: "Подготовка к итоговому экзамену", Hours: 10, Image: "https://via.placeholder.com/150"},
	}
	r := gin.Default()
	// Настраиваем маршрут для статических файлов
	r.Static("/static", "./static")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"logo1": "http://127.0.0.1:9000/prog/pudge.png",
			"tasks": tasks,
		})
	})
	r.GET("/task/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || id < 1 || id > len(tasks) {
			c.String(http.StatusNotFound, "Задача не найдена")
			return
		}

		task := tasks[id-1]
		c.HTML(http.StatusOK, "task.tmpl", gin.H{
			"task": task,
		})
	})
	r.LoadHTMLGlob("templates/*")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}

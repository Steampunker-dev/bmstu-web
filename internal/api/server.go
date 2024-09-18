package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	var specialTaskIDs = []int{1, 3} // Например, задания с ID 1 и 3

	// Данные для карточек
	tasks := []Task{
		{ID: 1, Title: "РК 1 модуль", Description: "Рубежный контроль за первый модуль семестра 2024", Hours: 3, Image: "http://127.0.0.1:9000/prog/RK1.png"},
		{ID: 2, Title: "РК 2 модуль", Description: "Рубежный контроль за второй модуль семестра 2024", Hours: 3, Image: "http://127.0.0.1:9000/prog/RK1.png"},
		{ID: 3, Title: "Подготовка к экзамену", Description: "Подготовка к итоговому экзамену", Hours: 9, Image: "http://127.0.0.1:9000/prog/DZ1.png"},
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
		searchQuery := c.Query("search")
		var filteredTasks []Task

		if searchQuery != "" {
			for _, task := range tasks {
				if strings.Contains(strings.ToLower(task.Title), strings.ToLower(searchQuery)) {
					filteredTasks = append(filteredTasks, task)
				}
			}
		} else {
			filteredTasks = tasks
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"logo1":       "http://127.0.0.1:9000/prog/pudge.png",
			"tasks":       filteredTasks,
			"searchQuery": searchQuery,
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
	r.GET("/special-tasks", func(c *gin.Context) {
		var specialTasks []Task
		for _, task := range tasks {
			for _, id := range specialTaskIDs {
				if task.ID == id {
					specialTasks = append(specialTasks, task)
					break
				}
			}
		}
		c.HTML(http.StatusOK, "special_tasks.tmpl", gin.H{
			"tasks": specialTasks,
		})
	})

	r.LoadHTMLGlob("templates/*")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}

package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Minutes     int
	Image       string
	Answ        string
}

type RequestTask struct {
	TaskID     int
	RequestID  int
	IsRequired bool
}

type Request struct {
	RequestTasks []RequestTask
	DateCreate   time.Time
	RequestGroup string
}

var tasks = []Task{
	{ID: 1, Title: "Номер 3828 Демидович", Description: "Определить аналитическое выражение для данного интеграла, применяя методы теории интегралов.", Minutes: 10, Image: "http://127.0.0.1:9000/prog/num1.png", Answ: "http://127.0.0.1:9000/prog/3828answ.png"},
	{ID: 2, Title: "Номер 3805 Демидович", Description: "Найти значение этого интеграла, используя знания о нормальном распределении и преобразованиях Гаусса.", Minutes: 15, Image: "http://127.0.0.1:9000/prog/num2.png", Answ: "http://127.0.0.1:9000/prog/3805answ.png"},
	{ID: 3, Title: "Номер 3801 Демидович", Description: "Решить данный определённый интеграл, используя подходящие методики интегрирования, такие как тригонометрические подстановки или интегрирование по частям.", Minutes: 10, Image: "http://127.0.0.1:9000/prog/num3.png", Answ: "http://127.0.0.1:9000/prog/3801answ.png"},
}

var requests = []Request{
	{
		RequestTasks: []RequestTask{
			{TaskID: 1, RequestID: 1, IsRequired: true},
			{TaskID: 2, RequestID: 1, IsRequired: false},
			{TaskID: 3, RequestID: 1, IsRequired: true},
		},
		DateCreate:   time.Date(2023, time.October, 3, 0, 0, 0, 0, time.UTC),
		RequestGroup: "ИУ5-21Б",
	},
}

func getTasksForRequest(request Request) []Task {
	var result []Task
	for _, reqTask := range request.RequestTasks {
		for _, task := range tasks {
			if task.ID == reqTask.TaskID {
				result = append(result, task)
			}
		}
	}
	return result
}
func StartServer() {
	log.Println("Server start up")

	// Данные для карточек

	r := gin.Default()
	// Настраиваем маршрут для статических файлов
	r.Static("/static", "./static")

	r.GET("/tasks", func(c *gin.Context) {
		searchQuery := c.Query("find_task")
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

		// Подсчитываем количество специальных заданий
		specialTaskCount := 0
		for _, task := range tasks {
			for _, id := range requests[0].RequestTasks {
				if task.ID == id.TaskID {
					specialTaskCount++
					break
				}
			}
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"logo1":            "http://127.0.0.1:9000/prog/pudge.png",
			"tasks":            filteredTasks,
			"searchQuery":      searchQuery,
			"SpecialTaskCount": specialTaskCount,
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
	r.GET("/special-tasks/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || id < 1 || id > len(tasks) {
			c.String(http.StatusNotFound, "Задача не найдена")
			return
		}
		request := requests[id-1] // Предполагаем, что у нас одна заявка для примера
		tasksForRequest := getTasksForRequest(request)
		data := struct {
			Request Request
			Tasks   []Task
		}{
			Request: request,
			Tasks:   tasksForRequest,
		}
		log.Println(data)
		c.HTML(http.StatusOK, "special_tasks.tmpl", data)
	})

	r.LoadHTMLGlob("templates/*")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}

package main

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Represents to-do list item
type Task struct {
	Id   int
	Text string
}

// Manages tasks in the to-do list
type TaskManager struct {
	nextTaskId int
	mutex      sync.Mutex
	pool       *redis.Pool
}

var taskManager TaskManager

func main() {
	taskManager = TaskManager{
		pool: &redis.Pool{
			MaxIdle:   10,
			MaxActive: 100,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", "localhost:6379")
			},
		},
	}

	// register handlers
	router := mux.NewRouter()
	router.HandleFunc("/", GetAllTasks).Methods("GET")

	log.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])
	if err != nil || page <= 0 {
		page = 1
	}

	conn := taskManager.pool.Get()
	defer conn.Close()

	start := (page - 1) * 5
	stop := start + 5 - 1

	taskIds, err := redis.Ints(conn.Do("LRANGE", "taskIds", start, stop))
	if err != nil {
		return
	}

	var tasks []Task
	for _, id := range taskIds {
		text, err := redis.String(conn.Do("HGET", "task:"+strconv.Itoa(id), "text"))
		if err != nil {
			return
		}

		tasks = append(tasks, Task{Id: id, Text: text})
	}

}

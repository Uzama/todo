package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Uzama/todo/domain"
	"github.com/Uzama/todo/repo"
	"github.com/gorilla/mux"
)

type handler struct {
	usecase domain.TaskUsecase
}

func NewHandler() *handler {
	return &handler{
		usecase: domain.NewTaskUsecase(repo.NewTaskRepo()),
	}
}

func (handl *handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := handl.usecase.ListTasks()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	payload, _ := json.Marshal(tasks)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (handl *handler) CreateTask(w http.ResponseWriter, r *http.Request) {

	task := domain.Task{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	log.Printf("request-[title:%s, description:%s]\n", task.Title, task.Description)

	if len(task.Title) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("title is required"))
		return
	}

	id, err := handl.usecase.CreateTask(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	log.Printf("response-[id:%d]", id)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))
}

func (handl *handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.String()))
}

func (handl *handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	num, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	deleted, err := handl.usecase.DeleteTask(num)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(fmt.Sprintf("%t", deleted)))
}

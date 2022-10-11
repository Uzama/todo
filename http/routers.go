package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()

	handler := NewHandler()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Task APP"))
	})

	r.HandleFunc("/list", handler.ListTasks).Methods(http.MethodGet)
	r.HandleFunc("/create", handler.CreateTask).Methods(http.MethodPost)
	r.HandleFunc("/update", handler.UpdateTask).Methods(http.MethodPut)
	r.HandleFunc("/delete/{id}", handler.DeleteTask).Methods(http.MethodDelete)

	return r
}

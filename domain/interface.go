package domain

type TaskRepoInterface interface {
	ListTasks() ([]Task, error)
	CreateTask(task Task) (int, error)
	UpdateTask(id int, task Task) (bool, error)
	DeleteTask(id int) (bool, error)
}

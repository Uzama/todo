package domain

type TaskUsecase struct {
	repo TaskRepoInterface
}

func NewTaskUsecase(repo TaskRepoInterface) TaskUsecase {
	return TaskUsecase{
		repo: repo,
	}
}

func (usecase TaskUsecase) ListTasks() ([]Task, error) {
	return usecase.repo.ListTasks()
}

func (usecase TaskUsecase) CreateTask(task Task) (int, error) {

	return usecase.repo.CreateTask(task)
}

func (usecase TaskUsecase) UpdateTask(id int, task Task) (bool, error) {
	return false, nil
}

func (usecase TaskUsecase) DeleteTask(id int) (bool, error) {
	return usecase.repo.DeleteTask(id)
}

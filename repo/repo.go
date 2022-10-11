package repo

import (
	"database/sql"
	"fmt"

	"github.com/Uzama/todo/domain"
)

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepo() domain.TaskRepoInterface {
	return &TaskRepo{
		db: InitMysql(),
	}
}

func (repo *TaskRepo) ListTasks() ([]domain.Task, error) {

	stmt, err := repo.db.Prepare(`SELECT id,title,description from tasks`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := make([]domain.Task, 0)

	for rows.Next() {

		task := domain.Task{}

		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
		)
		if err != nil {
			fmt.Println(err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repo *TaskRepo) CreateTask(task domain.Task) (int, error) {

	stmt, err := repo.db.Prepare(`INSERT INTO tasks (title, description) VALUES(?, ?);`)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(task.Title, task.Description)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (repo *TaskRepo) UpdateTask(id int, task domain.Task) (bool, error) {
	return false, nil
}

func (usecase *TaskRepo) DeleteTask(id int) (bool, error) {
	stmt, err := usecase.db.Prepare(`DELETE FROM tasks where id=?`)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

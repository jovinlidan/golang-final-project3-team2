package task_domain

import "golang-final-project3-team2/db"

var TaskDomain taskDomainRepo = &taskDomain{}

const (
	queryCreateTask         = `INSERT INTO tasks (title, description, category_id, status) VALUES ($1, $2, $3, $4) RETURNING id, title, description, category_id, status, created_at`
	queryGetTasks           = `SELECT id, title, description, category_id, status, created_at, updated_at FROM tasks WHERE user_id = $1`
	queryUpdateTask         = `UPDATE tasks SET title = $1, description = $2, category_id = $3, status = $4, updated_at = now() WHERE id = $5 RETURNING id, title, description, category_id, status, updated_at`
	queryUpdateTaskStatus   = `UPDATE tasks SET status = $1, updated_at = now() WHERE id = $2 RETURNING id, title, description, category_id, status, updated_at`
	queryUpdateTaskCategory = `UPDATE tasks SET category_id = $1, updated_at = now() WHERE id = $2 RETURNING id, title, description, category_id, status, updated_at`
	queryDeleteTask         = `DELETE FROM tasks WHERE id = $1`
)

type taskDomainRepo interface {
	CreateTask(*Task) (*Task, error)
	GetTasks(int64) (*[]Task, error)
	UpdateTask(*Task) (*Task, error)
	UpdateTaskStatus(*Task) (*Task, error)
	UpdateTaskCategory(*Task) (*Task, error)
	DeleteTask(int64) error
}

type taskDomain struct{}

func (u *taskDomain) CreateTask(task *Task) (*Task, error) {
	row := db.GetDB().QueryRow(queryCreateTask, task.Title, task.Description, task.CategoryId, task.Status)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(&task.Id, &task.Title, &task.Description, &task.CategoryId, &task.Status, &task.CreatedAt)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (u *taskDomain) GetTasks(userId int64) (*[]Task, error) {
	rows, err := db.GetDB().Query(queryGetTasks, userId)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.CategoryId, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return &tasks, nil
}

func (u *taskDomain) UpdateTask(task *Task) (*Task, error) {
	row := db.GetDB().QueryRow(queryUpdateTask, task.Title, task.Description, task.CategoryId, task.Status, task.Id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(&task.Id, &task.Title, &task.Description, &task.CategoryId, &task.Status, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (u *taskDomain) UpdateTaskStatus(task *Task) (*Task, error) {
	row := db.GetDB().QueryRow(queryUpdateTaskStatus, task.Status, task.Id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(&task.Id, &task.Title, &task.Description, &task.CategoryId, &task.Status, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (u *taskDomain) UpdateTaskCategory(task *Task) (*Task, error) {
	row := db.GetDB().QueryRow(queryUpdateTaskCategory, task.CategoryId, task.Id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(&task.Id, &task.Title, &task.Description, &task.CategoryId, &task.Status, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (u *taskDomain) DeleteTask(id int64) error {
	_, err := db.GetDB().Exec(queryDeleteTask, id)
	if err != nil {
		return err
}


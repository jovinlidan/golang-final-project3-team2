package task_services

import (
	"golang-final-project3-team2/resources/task_resources"
	"golang-final-project3-team2/utils/error_utils"
)

var TaskService taskServiceRepo = &taskService{}

type taskServiceRepo interface {
	CreateTask(*task_resources.TaskCreateRequest, string) (*task_resources.TaskCreateResponse, error_utils.MessageErr)
	GetTasks() (*[]task_resources.TaskGetResponse, error_utils.MessageErr)
	UpdateTask(*task_resources.TaskUpdateRequest, string) (*task_resources.TaskUpdateResponse, error_utils.MessageErr)
	UpdateTaskStatus(*task_resources.TaskUpdateStatusRequest, string) (*task_resources.TaskUpdateStatusResponse, error_utils.MessageErr)
	UpdateTaskCategory(*task_resources.TaskUpdateCategoryRequest, string) (*task_resources.TaskUpdateCategoryResponse, error_utils.MessageErr)
	DeleteTask(string) error_utils.MessageErr
}

type taskService struct{}

func (u *taskService) CreateTask(request *task_resources.TaskCreateRequest, userId string) (*task_resources.TaskCreateResponse, error_utils.MessageErr) {
	task, err := TaskDomain.CreateTask(request, userId)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (u *taskService) GetTasks() (*[]task_resources.TaskGetResponse, error_utils.MessageErr) {
	tasks, err := TaskDomain.GetTasks()

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (u *taskService) UpdateTask(request *task_resources.TaskUpdateRequest, taskId string) (*task_resources.TaskUpdateResponse, error_utils.MessageErr) {
	updatedTask, err := TaskDomain.UpdateTask(request, taskId)

	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (u *taskService) UpdateTaskStatus(request *task_resources.TaskUpdateStatusRequest, taskId string) (*task_resources.TaskUpdateStatusResponse, error_utils.MessageErr) {
	updatedTask, err := TaskDomain.UpdateTaskStatus(request, taskId)

	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (u *taskService) UpdateTaskCategory(request *task_resources.TaskUpdateCategoryRequest, taskId string) (*task_resources.TaskUpdateCategoryResponse, error_utils.MessageErr) {
	updatedTask, err := TaskDomain.UpdateTaskCategory(request, taskId)

	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (u *taskService) DeleteTask(id string) error_utils.MessageErr {
	err := TaskDomain.DeleteTask(id)

	if err != nil {
		return err
	}

	return nil
}

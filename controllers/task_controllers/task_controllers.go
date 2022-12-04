package task_controllers

import (
	"github.com/gin-gonic/gin"
	"golang-final-project3-team2/resources/task_resources"
	"golang-final-project3-team2/services/task_services"
	"golang-final-project3-team2/utils/error_utils"
	"net/http"
)

func CreateTask(c *gin.Context) {
	var taskReq task_resources.TaskCreateRequest
	userIdToken := c.MustGet("user_id")
	if err := c.ShouldBindJSON(&taskReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	task, err := task_services.TaskService.CreateTask(&taskReq, userIdToken.(string))

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, task)
}

func GetTask(c *gin.Context) {
	taskId := c.Param("task_id")
	userIdToken := c.MustGet("user_id")
	task, err := task_services.TaskService.GetTask(taskId, userIdToken.(string))

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func GetTasks(c *gin.Context) {
	userIdToken := c.MustGet("user_id")
	tasks, err := task_services.TaskService.GetTasks(userIdToken.(string))

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {
	var taskReq task_resources.TaskUpdateRequest
	taskId := c.Param("task_id")
	userIdToken := c.MustGet("user_id")
	if err := c.ShouldBindJSON(&taskReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	task, err := task_services.TaskService.UpdateTask(&taskReq, taskId, userIdToken.(string))

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateStatusTask(c *gin.Context) {
	var taskReq task_resources.TaskUpdateStatusRequest
	taskId := c.Param("task_id")
	userIdToken := c.MustGet("user_id")
	if err := c.ShouldBindJSON(&taskReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	task, err := task_services.TaskService.UpdateStatusTask(&taskReq, taskId, userIdToken.(string))

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateCategoryTask(c *gin.Context) {
	var taskReq task_resources.TaskUpdateCategoryRequest
	taskId := c.Param("task_id")
	userIdToken := c.MustGet("user_id")
	if err := c.ShouldBindJSON(&taskReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	task, err := task_services.TaskService.UpdateCategoryTask(&taskReq, taskId, userIdToken.(string))

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	taskId := c.Param("task_id")
	userIdToken := c.MustGet("user_id")
	err := task_services.TaskService.DeleteTask(taskId, userIdToken.(string))

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

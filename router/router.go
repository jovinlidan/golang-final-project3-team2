package router

import (
	"github.com/gin-gonic/gin"
	"golang-final-project3-team2/controllers/category_controllers"
	"golang-final-project3-team2/controllers/task_controllers"
	"golang-final-project3-team2/controllers/user_controllers"
	"golang-final-project3-team2/db"
	"golang-final-project3-team2/middlewares"
	"log"
)

const PORT = ":8080"

func init() {
	db.InitializeDB()

}

func StartRouter() {
	router := gin.Default()
	apiRouter := router.Group("/api")
	{
		userRouter := apiRouter.Group("/users")
		{
			userRouter.POST("/register", user_controllers.CreateUser)
			userRouter.POST("/login", user_controllers.UserLogin)
			userRouter.Use(middlewares.MiddlewareAuth())
			userRouter.PUT("/update-account", user_controllers.UpdateUser)
			userRouter.DELETE("/delete-account", user_controllers.DeleteUser)
		}

		categoryRouter := apiRouter.Group("/categories")
		{
			categoryRouter.Use(middlewares.MiddlewareAuth())
			categoryRouter.GET("/", category_controllers.GetCategories)
			categoryRouter.Use(middlewares.MiddlewareOnlyAdmin())
			categoryRouter.POST("/", category_controllers.CreateCategory)
			categoryRouter.PATCH("/:categoryId", category_controllers.UpdateCategory)
			categoryRouter.DELETE("/:categoryId", category_controllers.DeleteCategory)
		}

		taskRouter := apiRouter.Group("/tasks")
		{
			taskRouter.Use(middlewares.MiddlewareAuth())
			taskRouter.GET("/", task_controllers.GetTasks)
			taskRouter.POST("/", task_controllers.CreateTask)
			taskRouter.PUT("/:taskId", task_controllers.UpdateTask)
			taskRouter.PATCH("/update-category/:taskId", task_controllers.UpdateCategoryTask)
			taskRouter.PATCH("/update-status/:taskId", task_controllers.UpdateStatusTask)
			taskRouter.DELETE("/:taskId", task_controllers.DeleteTask)
		}
	}

	err := router.Run(PORT)
	if err != nil {
		log.Fatal(err.Error())
	}
}

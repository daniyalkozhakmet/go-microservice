package application

import (
	"microservice/handler"
	"microservice/handler/template"
	"microservice/middleware"
	"microservice/repository"

	"github.com/gin-gonic/gin"
)

func (a *App) loadRoutes() {
	a.Router = a.Routes()
}

func (a *App) Routes() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	a.orderRoutes(router)
	a.userRoutes(router)
	a.basicRoutes(router)
	a.templateAuthRoutes(router)

	return router
}
func (a *App) basicRoutes(router *gin.Engine) {

	basicHandler := &template.Basic{
		BasicRepo: &repository.BasicRepo{
			DB: a.DB,
		},
	}

	router.GET("/hello", basicHandler.List)
	router.GET("/", basicHandler.Create)

}
func (a *App) templateAuthRoutes(router *gin.Engine) {

	authHandler := &template.TemplateUserHandler{
		UserRepo: &repository.UserRepo{
			DB: a.DB,
		},
	}
	router.GET("/login", authHandler.LoginView)
	router.GET("/register", authHandler.RegisterView)
}
func (a *App) orderRoutes(router *gin.Engine) {
	orderHandler := &handler.Order{
		OrderRepo: &repository.OrderRepo{
			DB: a.DB,
		},
	}
	orderGroup := router.Group("/orders")
	{
		orderGroup.GET("/", middleware.JWTMiddleware(), orderHandler.List)
		orderGroup.POST("/", orderHandler.Create)
		orderGroup.GET("/:id", orderHandler.GetByID)
		orderGroup.PUT("/:id", orderHandler.UpdateByID)
	}
}
func (a *App) userRoutes(router *gin.Engine) {
	userHandler := &handler.UserHandler{
		UserRepo: &repository.UserRepo{
			DB: a.DB,
		},
	}
	userGroup := router.Group("/auth")
	{
		userGroup.POST("/register", userHandler.SignUp)
		userGroup.POST("/login", userHandler.SignIn)
		userGroup.PUT("/:id", userHandler.UpdateByID)
	}
}

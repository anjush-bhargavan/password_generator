package server

import (
	"github.com/anjush-bhargavan/password_generator/pkg/handler"
	"github.com/anjush-bhargavan/password_generator/pkg/routes"
	"github.com/anjush-bhargavan/password_generator/pkg/service"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	passwordService := service.NewPasswordService()
	passwordHandler := handler.NewPasswordHandler(passwordService)
	routes.InitRoutes(&passwordHandler,router)


	return router
}
package main

import (
	"github.com/rohitsmart/studio/controller"
	"github.com/rohitsmart/studio/database"
	"github.com/rohitsmart/studio/route"
	"github.com/rohitsmart/studio/service"
)

func main() {
	database.InitDB()
	authService := service.NewAuthService()
	authController := controller.NewAuthController(authService)
	r := route.SetupRouter(authController)
	r.Run(":9090")
}

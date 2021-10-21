package router

import (
	"github.com/bangnh1/golang-training/10/controller"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	v1 := app.Group("/v1")
	{
		v1.Get("/users", controller.GetAllUser)
		v1.Get("/user/:id", controller.GetUserById)
		v1.Post("/user", controller.CreateUser)
		v1.Put("/user/:id", controller.UpdateUser)
		v1.Delete("/user/:id", controller.DeleteUser)
	}

}

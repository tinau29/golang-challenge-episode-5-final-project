package routes

import (
	"api/app/controller"
	"api/app/controller/employee"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

// Handle all request to route to controller
func Handle(app *fiber.App) {
	services.InitRedis()
	services.InitDatabase()
	app.Use(cors.New())

	api := app.Group(viper.GetString("ENDPOINT"))

	api.Get("/", controller.GetAPIIndex)
	api.Get("/info.json", controller.GetAPIInfo)

	api.Get("/employee", employee.GetEmployee)
	api.Get("/employee/:id", employee.GetEmployeeDetail)
	api.Post("/employee", employee.PostEmployee)
	api.Put("/employee/:id", employee.PutEmployee)
	api.Delete("/employee/:id", employee.DeleteEmployee)
	api.Patch("/employee/:id", employee.PatchEmployeeChangeActive)
}

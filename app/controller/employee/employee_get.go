package employee

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetEmployee
// @Summary Get all employee
// @Description Get all employee
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Employee{} lib.Response "success"
// @Router /employee [get]
// @Tags Employee
func GetEmployee(c *fiber.Ctx) error {
	services.InitDatabase()
	db := services.DB

	employee := []model.Employee{}
	db.Model(&model.Employee{}).Find(&employee)

	return lib.OK(c, employee)
}

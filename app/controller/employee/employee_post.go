package employee

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostEmployee Create a employee by ID
// @Summary Create a employee by ID
// @Description Create a employee by ID
// @Accept application/json
// @Produce application/json
// @Success 200 {object} lib.Response "success"
// @Failure 400 {object} lib.Response "bad request"
// @Param data body model.Employee true "Employee data"
// @Router /employee [post]
// @Tags Employee
func PostEmployee(c *fiber.Ctx) error {
	services.InitDatabase()
	db := services.DB

	employee := model.Employee{}
	if err := c.BodyParser(&employee); err != nil {
		return lib.ErrorBadRequest(c)
	}

	db.Create(&employee)
	return lib.OK(c, employee)
}

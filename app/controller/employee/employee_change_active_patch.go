package employee

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"errors"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// PatchEmployeeChangeActive Update a employee by ID
// @Summary Update a employee by ID
// @Description Update a employee by ID
// @Accept application/json
// @Produce application/json
// @Success 200 {object} lib.Response "success"
// @Failure 400 {object} lib.Response "bad request"
// @Failure 404 {object} lib.Response "not found"
// @Param id path string true "Employee ID"
// @Param data body model.Employee true "Employee data"
// @Router /employee/{id} [patch]
// @Tags Employee
func PatchEmployeeChangeActive(c *fiber.Ctx) error {
	services.InitDatabase()
	db := services.DB

	id := c.Params("id")
	regex := regexp.MustCompile(`^[0-9]+$`)
	if !regex.MatchString(id) {
		return lib.ErrorBadRequest(c)
	}

	employeeExist := model.Employee{}
	if err := db.Where(`id = ?`, id).First(&employeeExist).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return lib.ErrorNotFound(c)
	}

	employee := model.Employee{}
	if err := c.BodyParser(&employee); err != nil {
		return lib.ErrorBadRequest(c)
	}

	employeeExist.IsActive = employee.IsActive
	db.Where(`id = ?`, id).Updates(&employeeExist)
	return lib.OK(c, employeeExist)

}

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

// GetEmployeeDetail Get a employee by ID
// @Summary Get a employee by ID
// @Description Get a employee by ID
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.Employee{} lib.Response "success"
// @Failure 400 {object} lib.Response "bad request"
// @Failure 404 {object} lib.Response "not found"
// @Param id path string true "Employee ID"
// @Router /employee/{id} [get]
// @Tags Employee
func GetEmployeeDetail(c *fiber.Ctx) error {
	services.InitDatabase()
	db := services.DB

	id := c.Params("id")
	regex := regexp.MustCompile(`^[0-9]+$`)
	if !regex.MatchString(id) {
		return lib.ErrorBadRequest(c)
	}

	employee := model.Employee{}
	if err := db.Where(`id = ?`, id).First(&employee).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return lib.ErrorNotFound(c)
	}

	return lib.OK(c, employee)

}

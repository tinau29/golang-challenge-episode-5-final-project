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

// PutEmployee
// @Summary Update a employee by ID
// @Description Update a employee by ID
// @Accept application/json
// @Produce application/json
// @Success 200 {object} lib.Response "success"
// @Failure 400 {object} lib.Response "bad request"
// @Failure 404 {object} lib.Response "not found"
// @Param data body model.Employee true "Employee data"
// @Router /employee/{id} [put]
// @Tags Employee
func PutEmployee(c *fiber.Ctx) error {
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

	db.Where(`id = ?`, id).Updates(&employee)

	// name := employeeExist.Name
	// if employee.Name != "" {
	// 	name = *employee.Name
	// }

	// salary := employeeExist.Salary
	// if employee.Salary != 0.0 {
	// 	salary = employee.Salary
	// }

	// address := employeeExist.Address
	// if employee.Address != "" {
	// 	address = employee.Address
	// }

	// BirthOfDate := employeeExist.BirthOfDate
	// if employee.BirthOfDate != "" {
	// 	birthDay = employee.BirthOfDate
	// }

	// joinOfDate := employeeExist.JoinOfDate
	// if employee.JoinOfDate != "" {
	// 	joinDate = employee.JoinOfDate
	// }

	employee.ID = employeeExist.ID

	return lib.OK(c, employee)
}

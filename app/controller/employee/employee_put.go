package employee

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"errors"
	"log"
	"reflect"
	"regexp"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// PutEmployee Update a employee by ID
// @Summary Update a employee by ID
// @Description Update a employee by ID
// @Accept application/json
// @Produce application/json
// @Success 200 {object} lib.Response "success"
// @Failure 400 {object} lib.Response "bad request"
// @Failure 404 {object} lib.Response "not found"
// @Param id path string true "Employee ID"
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

	log.Println("cek aja ", reflect.TypeOf(employeeExist.AddressDetail))
	log.Println("cek kedua", employeeExist.AddressDetail)

	db.Where(`id = ?`, id).Updates(&employee)
	employee.ID = employeeExist.ID

	if employee.Name == nil {
		employee.Name = lib.Strptr(*employeeExist.Name)
	}

	if employee.Address == nil {
		employee.Address = lib.Strptr(*employeeExist.Address)
	}

	if employee.Position == nil {
		employee.Position = lib.Strptr(*employeeExist.Position)
	}

	if employee.Salary == nil {
		employee.Salary = lib.Float64ptr(*employeeExist.Salary)
	}

	if employee.IsActive == nil {
		employee.IsActive = lib.Boolptr(*employeeExist.IsActive)
	}

	if employee.IsContract == nil {
		employee.IsContract = lib.Boolptr(*employeeExist.IsContract)
	}

	if employee.BirthOfDate == nil {
		employee.BirthOfDate = (*strfmt.Date)(lib.Timeptr(time.Time(*employeeExist.BirthOfDate)))
	}

	if employee.JoinOfDate == nil {
		employee.JoinOfDate = (*strfmt.Date)(lib.Timeptr(time.Time(*employeeExist.JoinOfDate)))
	}

	log.Println(len(employeeExist.AddressDetail))

	if len(employee.AddressDetail) == 0 {
		employee.AddressDetail = employeeExist.AddressDetail
	}

	return lib.OK(c, employee)
}

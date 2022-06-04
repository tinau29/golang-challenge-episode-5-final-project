package employee

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"net/http"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestDeleteEmployee(t *testing.T) {
	services.InitDatabaseForTest()
	db := services.DB

	app := fiber.New()
	app.Delete("/api/v1/employee/:id", DeleteEmployee)

	employee := model.Employee{}
	employee.Name = lib.Strptr("Joni")
	employee.Address = lib.Strptr("Jl. jalan no 1")
	employee.Position = lib.Strptr("Programmer")
	employee.Salary = lib.Float64ptr(2000)
	employee.IsActive = lib.Boolptr(true)
	employee.IsContract = lib.Boolptr(false)

	db.Create(&employee)
	id := strconv.Itoa(*employee.ID)

	// positive case
	url := "/api/v1/employee/" + id
	req, _ := http.NewRequest("DELETE", url, nil)
	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "response code")

	// regex error case
	url = "/api/v1/employee/non-existing-id"
	req, _ = http.NewRequest("DELETE", url, nil)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "response bad body parser")

	// non existing id case
	url = "/api/v1/employee/0"
	req, _ = http.NewRequest("DELETE", url, nil)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 404, res.StatusCode, "response not existing id")

}

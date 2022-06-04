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

func TestEmployeeDetail(t *testing.T) {
	services.InitDatabaseForTest()
	db := services.DB

	app := fiber.New()
	app.Get("/api/v1/employee/:id", GetEmployeeDetail)

	employee := model.Employee{}
	employee.Name = lib.Strptr("Joni")
	employee.Address = lib.Strptr("Jl. jalan no 1")
	employee.Position = lib.Strptr("Programmer")
	employee.Salary = lib.Float64ptr(2000)
	employee.IsActive = lib.Boolptr(true)
	employee.IsContract = lib.Boolptr(false)

	db.Create(&employee)
	id := strconv.Itoa(*employee.ID)

	url := "/api/v1/employee/" + id
	req, _ := http.NewRequest("GET", url, nil)
	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "response 200")

	url = "/api/v1/employee/test"
	req, _ = http.NewRequest("GET", url, nil)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "response 400")

	url = "/api/v1/employee/0"
	req, _ = http.NewRequest("GET", url, nil)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 404, res.StatusCode, "response 404")

}

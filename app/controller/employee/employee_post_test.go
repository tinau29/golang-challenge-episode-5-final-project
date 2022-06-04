package employee

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"bytes"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPostEmployee(t *testing.T) {
	services.InitDatabaseForTest()
	db := services.DB

	app := fiber.New()
	app.Post("/api/v1/employee", PostEmployee)

	payload := bytes.NewReader([]byte(`
	{ 
		"Name": "Tino",
		"Address": "Jl, jalan no 1",
		"position": "Programmer",
		"salary": 2000,
		"is_active": true,
		"is_contract": false
	}
	`))

	employee := model.Employee{}
	employee.Name = lib.Strptr("Joni")
	employee.Address = lib.Strptr("Jl. jalan no 1")
	employee.Position = lib.Strptr("Programmer")
	employee.Salary = lib.Float64ptr(2000)
	employee.IsActive = lib.Boolptr(true)
	employee.IsContract = lib.Boolptr(false)
	db.Create(&employee)

	req, _ := http.NewRequest("POST", "/api/v1/employee", nil)
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid request")

	req, _ = http.NewRequest("POST", "/api/v1/employee", payload)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "response code")

}

package employee

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"bytes"
	"net/http"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPatchEmployeeChangeActive(t *testing.T) {
	services.InitDatabaseForTest()
	db := services.DB

	app := fiber.New()
	app.Patch("/api/v1/employee/:id", PatchEmployeeChangeActive)

	employee := model.Employee{}
	employee.Name = lib.Strptr("Joni")
	employee.Address = lib.Strptr("Jl. jalan no 1")
	employee.Position = lib.Strptr("Programmer")
	employee.Salary = lib.Float64ptr(2000)
	employee.IsActive = lib.Boolptr(true)
	employee.IsContract = lib.Boolptr(false)

	employeeDeleted := model.Employee{}
	employeeDeleted.Name = lib.Strptr("Joni")
	employeeDeleted.Address = lib.Strptr("Jl. jalan no 1")
	employeeDeleted.Position = lib.Strptr("Programmer")
	employeeDeleted.Salary = lib.Float64ptr(2000)
	employeeDeleted.IsActive = lib.Boolptr(true)
	employeeDeleted.IsContract = lib.Boolptr(false)

	db.Create(&employee)
	id := strconv.Itoa(*employee.ID)

	db.Create(&employeeDeleted)
	idDeleted := strconv.Itoa(*employeeDeleted.ID)
	db.Where(`id = ?`, idDeleted).Delete(&employeeDeleted)

	payload := bytes.NewReader([]byte(`
	{ 
		"is_active": true
	}
	`))

	req, _ := http.NewRequest("PATCH", "/api/v1/employee/"+id, payload)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "response code")

	req, _ = http.NewRequest("PATCH", "/api/v1/employee/"+id, nil)
	req.Header.Set("accept", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "response bad body parser")

	req, _ = http.NewRequest("PATCH", "/api/v1/employee/test", nil)
	req.Header.Set("accept", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid id format")

	req, _ = http.NewRequest("PATCH", "/api/v1/employee/"+idDeleted, nil)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 404, res.StatusCode, "not found")
}

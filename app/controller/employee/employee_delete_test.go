package employee

import (
	"testing"
)

func TestDeleteEmployee(t *testing.T) {
	// type TestDeleteEmployeeStruct struct {
	// 	description  string
	// 	route        string
	// 	expectedCode int
	// }

	// app := fiber.New()
	// app.Delete("/api/v1/employee/:id", DeleteEmployee)

	// employee := model.Employee{}
	// employee.Name = "Joni"
	// employee.Salary = 2000
	// employee.IsActive = true
	// employee.IsContract = false
	// employee.Address = "jl jalan no 1"
	// employee.Position = "sales marketing"

	// db := services.InitDatabaseForTest()
	// db.Create(&employee)
	// id := strconv.Itoa(*&employee.ID)

	// caseTest := []TestDeleteEmployeeStruct{
	// 	{
	// 		description:  "Get response 200",
	// 		route:        "/ap1/v1/employee/" + id,
	// 		expectedCode: 200,
	// 	},
	// 	{
	// 		description:  "Get response 400",
	// 		route:        "/api/v1/employee/test",
	// 		expectedCode: 400,
	// 	},
	// 	{
	// 		description:  "Get response 404",
	// 		route:        "/api/v1/employee/0",
	// 		expectedCode: 404,
	// 	},
	// }

	// for _, test := range caseTest {
	// 	req, _ := http.NewRequest("DELETE", test.route, nil)
	// 	req.Header.Set("accept", "application/json")
	// 	res, err := app.Test(req)
	// 	utils.AssertEqual(t, nil, err, "send request")
	// 	utils.AssertEqual(t, test.expectedCode, res.StatusCode, test.description)
	// }

}

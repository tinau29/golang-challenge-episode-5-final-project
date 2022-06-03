package employee

import (
	"testing"
)

func TestEmployeeDetail(t *testing.T) {
	// db := services.InitDatabaseForTest()
	// // db := services.DB

	// type TestEmployeeDetailStruct struct {
	// 	descriptionRequest  string
	// 	descriptionResponse string
	// 	route               string
	// 	expectedCode        int
	// }

	// app := fiber.New()
	// app.Get("/api/v1/employee/:id", GetDetailEmployee)

	// employee := model.Employee{}
	// *employee.Name = "Joni"
	// *employee.Address = "Jl. jalan no 1"
	// *employee.Position = "Programmer"
	// *employee.Salary = 2000
	// *employee.IsActive = true
	// *employee.IsContract = false

	// db.Create(&employee)
	// id := strconv.Itoa(*employee.ID)

	// caseTest := []TestEmployeeDetailStruct{
	// 	{
	// 		descriptionRequest:  "GET /api/v1/movies/" + id,
	// 		descriptionResponse: "get response success",
	// 		route:               "/api/v1/movies/" + id,
	// 		expectedCode:        200,
	// 	},
	// 	{
	// 		descriptionRequest:  "GET /api/v1/movies/test",
	// 		descriptionResponse: "get response bad request",
	// 		route:               "/api/v1/movies/test",
	// 		expectedCode:        400,
	// 	},
	// 	{
	// 		descriptionRequest:  "GET /api/v1/movies/0",
	// 		descriptionResponse: "get response not found",
	// 		route:               "/api/v1/movies/0",
	// 		expectedCode:        404,
	// 	},
	// }

	// for _, test := range caseTest {
	// 	res, _, err := lib.GetTest(app, test.route, nil)
	// 	utils.AssertEqual(t, nil, err, test.descriptionRequest)
	// 	utils.AssertEqual(t, 200, res.StatusCode, test.descriptionResponse)
	// }

}

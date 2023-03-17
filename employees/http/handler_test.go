package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	empHTTP "algogrit.com/emp_server/employees/http"
	"algogrit.com/emp_server/employees/service"
	"algogrit.com/emp_server/entities"
)

func TestIndexV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.New(mockSvc)

	req := httptest.NewRequest("GET", "/v1/employees", nil)
	respRec := httptest.NewRecorder()

	expectedEmps := []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
	}

	// mockSvc.EXPECT().Index(req.Context()).Return(expectedEmps, nil)
	mockSvc.EXPECT().Index(gomock.Any()).Return(expectedEmps, nil)

	// sut.IndexV1(respRec, req)
	sut.ServeHTTP(respRec, req)

	resp := respRec.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var actualEmps []entities.Employee

	err := json.NewDecoder(resp.Body).Decode(&actualEmps)

	assert.Nil(t, err)

	assert.NotNil(t, actualEmps)
	assert.NotEqual(t, 0, len(actualEmps))

	assert.Equal(t, expectedEmps[0].ID, actualEmps[0].ID)
	assert.Equal(t, expectedEmps[0].Name, actualEmps[0].Name)
	assert.Equal(t, expectedEmps[0].Department, actualEmps[0].Department)
}

func TestCreateV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.New(mockSvc)

	jsonBody := `{"name": "Gaurav", "speciality": "LnD"}`
	reqBody := strings.NewReader(jsonBody)

	req := httptest.NewRequest("POST", "/v1/employees", reqBody)
	respRec := httptest.NewRecorder()

	expectedEmp := entities.Employee{Name: "Gaurav", Department: "LnD"}
	createdEmp := entities.Employee{ID: 1, Name: "Gaurav", Department: "LnD"}

	mockSvc.EXPECT().Create(gomock.Any(), expectedEmp).Return(&createdEmp, nil)

	sut.ServeHTTP(respRec, req)

	resp := respRec.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var actualEmp entities.Employee

	json.NewDecoder(resp.Body).Decode(&actualEmp)

	assert.Equal(t, createdEmp, actualEmp)
}

// func FuzzCreateV1(f *testing.F) {
// 	f.Add(`{"name": "Gaurav", "speciality": "LnD"`)
// 	f.Add(`2312tsgfas`)

// 	f.Fuzz(func(t *testing.T, jsonBody string) {
// 		ctrl := gomock.NewController(t)
// 		defer ctrl.Finish()

// 		mockSvc := service.NewMockEmployeeService(ctrl)

// 		sut := empHTTP.New(mockSvc)

// 		reqBody := strings.NewReader(jsonBody)

// 		req := httptest.NewRequest("POST", "/v1/employees", reqBody)
// 		respRec := httptest.NewRecorder()

// 		sut.ServeHTTP(respRec, req)

// 		resp := respRec.Result()

// 		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
// 	})
// }

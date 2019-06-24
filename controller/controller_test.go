package controller

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tty2/company/model"
)

func TestCreateEmployee(t *testing.T) {
	ctx := context.Background()
	es := new(employeeStorageMock)
	firstEmployee := model.Employee{
		ID:        1,
		FirstName: "Name1",
		LastName:  "LastName1",
		Age:       33,
		Salary:    100,
	}
	secondEmployee := model.Employee{
		ID:        2,
		FirstName: "Name2",
		LastName:  "LastName2",
		Age:       33,
		Salary:    700,
	}

	es.On("Create", ctx, firstEmployee).Return(firstEmployee, nil)

	es.On("Create", ctx, secondEmployee).Return(
		emptyEmployee, fmt.Errorf("salary can't be higher then %d, set %d", 500, 700))

	s := New(es)

	tt := []struct {
		name             string
		employee         model.Employee
		expectedResponse model.Employee
		expectedErr      error
	}{
		{
			name:             "pass",
			employee:         firstEmployee,
			expectedResponse: firstEmployee,
			expectedErr:      nil,
		},
		{
			name:             "validation error",
			employee:         secondEmployee,
			expectedResponse: emptyEmployee,
			expectedErr: fmt.Errorf("validation failure on create employee with id %d: %v", 2,
				fmt.Errorf("salary can't be higher then %d, set %d", 500, 700)),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e, err := s.CreateEmployee(ctx, tc.employee)
			assert.Equal(t, tc.expectedResponse, e)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestGetEmployee(t *testing.T) {
	ctx := context.Background()
	es := new(employeeStorageMock)
	firstEmployee := model.Employee{
		ID:        1,
		FirstName: "Name1",
		LastName:  "LastName1",
		Age:       33,
		Salary:    100,
	}

	es.On("ByID", ctx, 1).Return(firstEmployee, nil)

	s := New(es)

	tt := []struct {
		name             string
		id               int
		expectedResponse model.Employee
		expectedErr      error
	}{
		{
			name:             "pass",
			id:               1,
			expectedResponse: firstEmployee,
			expectedErr:      nil,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e, err := s.GetEmployee(ctx, tc.id)
			assert.Equal(t, tc.expectedResponse, e)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestDeleteEmployee(t *testing.T) {
	ctx := context.Background()
	es := new(employeeStorageMock)

	es.On("Delete", ctx, 1).Return(nil)

	s := New(es)

	tt := []struct {
		name        string
		id          int
		expectedErr error
	}{
		{
			name:        "pass",
			id:          1,
			expectedErr: nil,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := s.DeleteEmployee(ctx, tc.id)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestRaiseSalary(t *testing.T) {
	ctx := context.Background()
	es := new(employeeStorageMock)

	firstEmployee := model.Employee{
		ID:        1,
		FirstName: "Name1",
		LastName:  "LastName1",
		Age:       33,
		Salary:    100,
	}
	secondEmployee := model.Employee{
		ID:        2,
		FirstName: "Name2",
		LastName:  "LastName2",
		Age:       33,
		Salary:    200,
	}
	thirdEmployee := model.Employee{
		ID:        3,
		FirstName: "Name3",
		LastName:  "LastName3",
		Age:       33,
		Salary:    100,
	}

	es.On("ByID", ctx, 1).Return(firstEmployee, nil)
	es.On("ByID", ctx, 2).Return(emptyEmployee,
		fmt.Errorf("employee with id %d doesn't exist", secondEmployee.ID))
	es.On("ByID", ctx, 3).Return(thirdEmployee, nil)

	es.On("Update", ctx, model.Employee{
		ID:        1,
		FirstName: "Name1",
		LastName:  "LastName1",
		Age:       33,
		Salary:    200,
	}).Return(nil)
	es.On("Update", ctx, thirdEmployee).Return(
		fmt.Errorf("employee with id %d doesn't exist", thirdEmployee.ID))

	s := New(es)

	tt := []struct {
		name        string
		id          int
		amount      int
		expectedErr error
	}{
		{
			name:        "pass",
			id:          firstEmployee.ID,
			amount:      100,
			expectedErr: nil,
		},
		{
			name:   "error: no employee",
			id:     secondEmployee.ID,
			amount: 100,
			expectedErr: fmt.Errorf("couldn't get employee with id %d: %v", 2,
				fmt.Errorf("employee with id %d doesn't exist", 2)),
		},
		{
			name:   "validation error",
			id:     firstEmployee.ID,
			amount: 500,
			expectedErr: fmt.Errorf("validation failure on update employee with id %d: %v", firstEmployee.ID,
				fmt.Errorf("salary can't be higher then %d, set %d", 500, 600)),
		},
		{
			name:   "update error",
			id:     thirdEmployee.ID,
			amount: 0,
			expectedErr: fmt.Errorf("couldn't update employee with id %d: %v", thirdEmployee.ID,
				fmt.Errorf("employee with id %d doesn't exist", thirdEmployee.ID)),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := s.RaiseSalary(ctx, tc.id, tc.amount)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

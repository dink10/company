// Package controller is a layer for business logic.
// All computation MUST be in this layer.
// It's a `brain` of the application.
//
// This layer has interface to the datastore and depends only from it.
// controller prepares request for datastore, gets data from it,
// does all computations and send data to handler layer.
//
// Services in controller MUST implement interfaces that was declared on
// handler layer.
package controller

import (
	"context"
	"fmt"

	"github.com/tty2/company/model"
)

type (
	employeeStorage interface {
		Create(ctx context.Context, employee model.Employee) (model.Employee, error)
		ByID(ctx context.Context, id int) (model.Employee, error)
		Update(ctx context.Context, employee model.Employee) error
		Delete(ctx context.Context, id int) error
	}

	// Service allows to communicate with datastore and is responsible for business logic.
	Service struct {
		EmployeeStore employeeStorage
	}
)

var emptyEmployee = model.Employee{}

// New creates employee service with dadtastore access.
func New(ds employeeStorage) Service {
	return Service{
		EmployeeStore: ds,
	}
}

// CreateEmployee method.
func (s Service) CreateEmployee(ctx context.Context, employee model.Employee) (model.Employee, error) {
	err := employee.Validate()
	if err != nil {
		return emptyEmployee, fmt.Errorf("validation failure on create employee with id %d: %v", employee.ID, err)
	}
	return s.EmployeeStore.Create(ctx, employee)
}

// GetEmployee gets employee from datastore by id.
func (s Service) GetEmployee(ctx context.Context, id int) (model.Employee, error) {
	return s.EmployeeStore.ByID(ctx, id)
}

// RaiseSalary method.
func (s Service) RaiseSalary(ctx context.Context, id int, amount int) error {
	em, err := s.EmployeeStore.ByID(ctx, id)
	if err != nil {
		return fmt.Errorf("couldn't get employee with id %d: %v", id, err)
	}
	em.Salary += amount
	err = em.Validate()
	if err != nil {
		return fmt.Errorf("validation failure on update employee with id %d: %v", id, err)
	}
	err = s.EmployeeStore.Update(ctx, em)
	if err != nil {
		return fmt.Errorf("couldn't update employee with id %d: %v", id, err)
	}
	return nil
}

// DeleteEmployee method.
func (s Service) DeleteEmployee(ctx context.Context, id int) error {
	return s.EmployeeStore.Delete(ctx, id)
}

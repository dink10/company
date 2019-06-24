// Package datastore is responsible for data storage.
// Doesn't matter which datastore has been chosen, there is only on requirement:
// it MUST implement the interface that declared on controller layer.
package datastore

import (
	"context"
	"fmt"

	"github.com/tty2/company/model"
)

var emptyEmployee = model.Employee{}

type employeeCollection map[int]model.Employee

func initEmployeeCollection() employeeCollection {
	return make(employeeCollection)
}

// Create method.
func (ec employeeCollection) Create(ctx context.Context, employee model.Employee) (model.Employee, error) {
	_, ok := ec[employee.ID]
	if ok {
		return emptyEmployee, fmt.Errorf("employee with id %d already exists", employee.ID)
	}
	ec[employee.ID] = employee
	return employee, nil
}

// ByID gets employee from datastore by id.
func (ec employeeCollection) ByID(ctx context.Context, id int) (model.Employee, error) {
	emp, ok := ec[id]
	if !ok {
		return emptyEmployee, fmt.Errorf("employee with id %d doesn't exist", id)
	}
	return emp, nil
}

// Update method.
func (ec employeeCollection) Update(ctx context.Context, employee model.Employee) error {
	_, ok := ec[employee.ID]
	if !ok {
		return fmt.Errorf("employee with id %d doesn't exist", employee.ID)
	}
	ec[employee.ID] = employee
	return nil
}

// Delete method.
func (ec employeeCollection) Delete(ctx context.Context, id int) error {
	_, ok := ec[id]
	if !ok {
		return fmt.Errorf("employee with id %d doesn't exist", id)
	}
	delete(ec, id)
	return nil
}

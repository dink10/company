// Package model keeps data models.
// It doesn't have any dependencies of other layers.
// It can be imported on any other layer.
package model

import (
	"errors"
	"fmt"
)

const (
	maxUint       = ^uint(0)
	maxInt        = int(maxUint >> 1)
	highestSalary = 500
	lowestSalary  = 100
)

type (
	// Employee represents type of user.
	Employee struct {
		ID        int    `json:"id,omitempty"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Age       int    `json:"age,omitempty"`
		Salary    int    `json:"salary,omitempty"`
	}
)

// Validate method.
func (e Employee) Validate() error {
	if e.ID <= 0 || e.ID >= maxInt {
		return errors.New("unavailable value of ID")
	}
	if e.Salary > highestSalary {
		return fmt.Errorf("salary can't be higher then %d, set %d", highestSalary, e.Salary)
	}
	if e.Salary < lowestSalary {
		return fmt.Errorf("salary can't be lower then %d, set %d", lowestSalary, e.Salary)
	}
	return nil
}

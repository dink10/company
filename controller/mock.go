package controller

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/tty2/company/model"
)

type employeeStorageMock struct {
	mock.Mock
}

// Create method.
func (es *employeeStorageMock) Create(ctx context.Context, employee model.Employee) (model.Employee, error) {
	args := es.Called(ctx, employee)
	return args.Get(0).(model.Employee), args.Error(1)
}

// ByID method.
func (es *employeeStorageMock) ByID(ctx context.Context, id int) (model.Employee, error) {
	args := es.Called(ctx, id)
	return args.Get(0).(model.Employee), args.Error(1)
}

// Update method.
func (es *employeeStorageMock) Update(ctx context.Context, employee model.Employee) error {
	args := es.Called(ctx, employee)
	return args.Error(0)
}

// Delete method.
func (es *employeeStorageMock) Delete(ctx context.Context, id int) error {
	args := es.Called(ctx, id)
	return args.Error(0)
}

package services

import (
	"github.com/imrinry/assessment-tax/models"
	"github.com/stretchr/testify/mock"
)

type taxServiceMock struct {
	mock.Mock
}

func NewTaxServiceMock() *taxServiceMock {
	return &taxServiceMock{}
}

func (m *taxServiceMock) TaxCalculations() (models.TaxResponse, error) {
	args := m.Called()
	return args.Get(0).(models.TaxResponse), args.Error(1)
}

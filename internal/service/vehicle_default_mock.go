package service

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

// NewServiceVehicleDefaultMock is a function that returns a new instance of ServiceVehicleDefaultMock
func NewServiceVehicleDefaultMock() *ServiceVehicleDefaultMock {
	return &ServiceVehicleDefaultMock{}
}

// ServiceVehicleDefaultMock is a struct that represents the mock service for vehicles
type ServiceVehicleDefaultMock struct {
	mock.Mock
}

// FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) is a method that returns a map of vehicles that match the color and fabrication year
func (s *ServiceVehicleDefaultMock) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	args := s.Called(color, fabricationYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) is a method that returns a map of vehicles that match the brand and a range of fabrication years
func (s *ServiceVehicleDefaultMock) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	args := s.Called(brand, startYear, endYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// AverageMaxSpeedByBrand(brand string) (a float64, err error) is a method that returns the average speed of the vehicles by brand
func (s *ServiceVehicleDefaultMock) AverageMaxSpeedByBrand(brand string) (a float64, err error) {
	args := s.Called(brand)
	return args.Get(0).(float64), args.Error(1)
}

// AverageCapacityByBrand(brand string) (a int, err error) is a method that returns the average capacity of the vehicles by brand
func (s *ServiceVehicleDefaultMock) AverageCapacityByBrand(brand string) (a int, err error) {
	args := s.Called(brand)
	return args.Get(0).(int), args.Error(1)
}

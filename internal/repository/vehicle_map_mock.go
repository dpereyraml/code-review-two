package repository

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

// NewRepositoryWriteVehicleMapMock is a function that returns a new instance of RepositoryWriteVehicleMapMock
func NewRepositoryWriteVehicleMapMock() *RepositoryWriteVehicleMapMock {
	return &RepositoryWriteVehicleMapMock{}
}

// RepositoryWriteVehicleMapMock is a struct that represents a vehicle repository mock
type RepositoryWriteVehicleMapMock struct {
	mock.Mock
}

// FindAll() (v map[int]Vehicle, err error) is a method that returns a map of all vehicles
func (r *RepositoryWriteVehicleMapMock) FindAll() (v map[int]internal.Vehicle, err error) {
	args := r.Called()
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// FindByColorAndYear(color string, fabricationYear int) (v map[int]Vehicle, err error) is a method that returns a map of vehicles that match the color and fabrication year
func (r *RepositoryWriteVehicleMapMock) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	args := r.Called(color, fabricationYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]Vehicle, err error) is a method that returns a map of vehicles that match the brand and a range of fabrication years
func (r *RepositoryWriteVehicleMapMock) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	args := r.Called(brand, startYear, endYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// FindByBrand(brand string) (v map[int]Vehicle, err error) is a method that returns a map of vehicles that match the brand
func (r *RepositoryWriteVehicleMapMock) FindByBrand(brand string) (v map[int]internal.Vehicle, err error) {
	args := r.Called(brand)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// FindByWeightRange(fromWeight float64, toWeight float64) (v map[int]Vehicle, err error) is a method that returns a map of vehicles that match the weight range
func (r *RepositoryWriteVehicleMapMock) FindByWeightRange(fromWeight float64, toWeight float64) (v map[int]internal.Vehicle, err error) {
	args := r.Called(fromWeight, toWeight)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

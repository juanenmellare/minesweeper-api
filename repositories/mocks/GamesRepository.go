// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	errors "minesweeper-api/errors"

	mock "github.com/stretchr/testify/mock"

	models "minesweeper-api/models"
)

// GamesRepository is an autogenerated mock type for the GamesRepository type
type GamesRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: game
func (_m *GamesRepository) Create(game *models.Game) *errors.ApiError {
	ret := _m.Called(game)

	var r0 *errors.ApiError
	if rf, ok := ret.Get(0).(func(*models.Game) *errors.ApiError); ok {
		r0 = rf(game)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.ApiError)
		}
	}

	return r0
}

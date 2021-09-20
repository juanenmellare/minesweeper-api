// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	errors "minesweeper-api/errors"

	mock "github.com/stretchr/testify/mock"

	models "minesweeper-api/models"

	uuid "github.com/google/uuid"
)

// FieldsRepository is an autogenerated mock type for the FieldsRepository type
type FieldsRepository struct {
	mock.Mock
}

// FindByIdAndGameId provides a mock function with given fields: _a0, gameUuid
func (_m *FieldsRepository) FindByIdAndGameId(_a0 *uuid.UUID, gameUuid *uuid.UUID) (*models.Field, *errors.ApiError) {
	ret := _m.Called(_a0, gameUuid)

	var r0 *models.Field
	if rf, ok := ret.Get(0).(func(*uuid.UUID, *uuid.UUID) *models.Field); ok {
		r0 = rf(_a0, gameUuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Field)
		}
	}

	var r1 *errors.ApiError
	if rf, ok := ret.Get(1).(func(*uuid.UUID, *uuid.UUID) *errors.ApiError); ok {
		r1 = rf(_a0, gameUuid)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.ApiError)
		}
	}

	return r0, r1
}

// FindMineFieldsFlaggedByGame provides a mock function with given fields: gameUuid
func (_m *FieldsRepository) FindMineFieldsFlaggedByGame(gameUuid *uuid.UUID) (*[]models.Field, *errors.ApiError) {
	ret := _m.Called(gameUuid)

	var r0 *[]models.Field
	if rf, ok := ret.Get(0).(func(*uuid.UUID) *[]models.Field); ok {
		r0 = rf(gameUuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Field)
		}
	}

	var r1 *errors.ApiError
	if rf, ok := ret.Get(1).(func(*uuid.UUID) *errors.ApiError); ok {
		r1 = rf(gameUuid)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.ApiError)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: field
func (_m *FieldsRepository) Update(field *models.Field) *errors.ApiError {
	ret := _m.Called(field)

	var r0 *errors.ApiError
	if rf, ok := ret.Get(0).(func(*models.Field) *errors.ApiError); ok {
		r0 = rf(field)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.ApiError)
		}
	}

	return r0
}

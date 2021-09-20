package helpers

import (
	"minesweeper-api/errors"
)

func ValidateDatabaseTxError(err error, baseMessage string) *errors.ApiError {
	if err != nil {
		switch err.Error() {
		case "record not found":
			return errors.NewNotFoundError(errors.NewError(baseMessage + " not found"))
		}
		return errors.NewInternalServerApiError(err)
	}

	return nil
}

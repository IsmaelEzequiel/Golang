package internalerrors

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
)

var ErrInternal error = errors.New(http.StatusText(http.StatusInternalServerError))

func ProcessErrorNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return ErrInternal
}

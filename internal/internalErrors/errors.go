package internalerrors

import (
	"errors"
	"net/http"
)

var ErrInternal error = errors.New(http.StatusText(http.StatusInternalServerError))

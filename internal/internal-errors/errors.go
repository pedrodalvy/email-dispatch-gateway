package internal_errors

import "errors"

var ErrInternalServerError = errors.New("internal server error")
var ErrResourceNotFound = errors.New("resource not found")

package backend

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var (
	ErrNotFound = ErrorWithCode{
		Err:  errors.New(http.StatusText(http.StatusNotFound)),
		Code: NotFound,
	}
)

type ErrorCode uint

const (
	NotFound        ErrorCode = http.StatusNotFound
	Unprocessable   ErrorCode = http.StatusUnprocessableEntity
	Internal        ErrorCode = http.StatusInternalServerError
	Unauthorized    ErrorCode = http.StatusUnauthorized
	Forbidden       ErrorCode = http.StatusForbidden
	InvalidArgument ErrorCode = http.StatusBadRequest
)

func (e ErrorCode) String() string {
	return strconv.FormatUint(uint64(e), 10)
}

type ErrorWithCode struct {
	Err  error
	Code ErrorCode
}

func (e ErrorWithCode) Error() string {
	return fmt.Sprintf("%v:%v", e.Code, e.Err.Error())
}

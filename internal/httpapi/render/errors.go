package render

import (
	"errors"
	"net/http"
)

// AppError — ошибка уровня API, которую можно безопасно показать клиенту.
type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
	Err        error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Code + ": " + e.Err.Error()
	}
	return e.Code + ": " + e.Message
}

func (e *AppError) Unwrap() error { return e.Err }

// E — конструктор AppError
func E(httpStatus int, code, message string, err error) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
		Err:        err,
	}
}

// Часто используемые ошибки
func BadRequest(code, message string, err error) *AppError {
	return E(http.StatusBadRequest, code, message, err)
}

func Unauthorized(code, message string, err error) *AppError {
	return E(http.StatusUnauthorized, code, message, err)
}

func Forbidden(code, message string, err error) *AppError {
	return E(http.StatusForbidden, code, message, err)
}

func NotFound(code, message string, err error) *AppError {
	return E(http.StatusNotFound, code, message, err)
}

func Conflict(code, message string, err error) *AppError {
	return E(http.StatusConflict, code, message, err)
}

func Internal(code, message string, err error) *AppError {
	return E(http.StatusInternalServerError, code, message, err)
}

// AsAppError — привести любую ошибку к AppError.
// Если это не AppError — оборачиваем в Internal.
func AsAppError(err error) *AppError {
	if err == nil {
		return nil
	}
	var appErr *AppError
	if errors.As(err, &appErr) {
		if appErr.HTTPStatus == 0 {
			appErr.HTTPStatus = http.StatusInternalServerError
		}
		if appErr.Code == "" {
			appErr.Code = "internal_error"
		}
		if appErr.Message == "" {
			appErr.Message = "Internal error"
		}
		return appErr
	}
	return Internal("internal_error", "Internal error", err)
}

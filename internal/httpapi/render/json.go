package render

import (
	"encoding/json"
	"net/http"
)

type envelope map[string]any

func JSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	// encode errors intentionally ignored to avoid double-write; logging will be in middleware later
	_ = json.NewEncoder(w).Encode(v)
}

func OK(w http.ResponseWriter, v any) {
	JSON(w, http.StatusOK, v)
}

func Accepted(w http.ResponseWriter, v any) {
	JSON(w, http.StatusAccepted, v)
}

// Error — единый формат ошибок.
func Error(w http.ResponseWriter, err error) {
	appErr := AsAppError(err)

	body := envelope{
		"error": envelope{
			"code":    appErr.Code,
			"message": appErr.Message,
		},
	}

	JSON(w, appErr.HTTPStatus, body)
}

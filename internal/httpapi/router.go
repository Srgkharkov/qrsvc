package httpapi

import (
	"net/http"

	"github.com/Srgkharkov/qrsvc/internal/httpapi/handlers"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.Health)

	return mux
}

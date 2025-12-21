package handlers

import (
	"net/http"
	"time"

	"github.com/Srgkharkov/qrsvc/internal/httpapi/render"
)

var startedAt = time.Now()

type healthResponse struct {
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	render.OK(w, healthResponse{
		Status: "ok",
		Uptime: time.Since(startedAt).String(),
	})
}

package controllers

import (
	"github.com/unrolled/render"
	"net/http"
)

func HealthCheckController(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Status string }{"OK"})
	}
}

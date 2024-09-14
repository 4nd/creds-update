package handlers

import (
	"net/http"
)

func (h *RouteHandler) Welcome(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "welcome", map[string]interface{}{
		"Name": "Andy",
	})
}

package handlers

import (
	"net/http"
)

func (h *RouteHandler) ListCreds(w http.ResponseWriter, r *http.Request) {

	h.render(w, "creds.html", map[string]interface{}{
		"Name": "Andy",
	})
}

package handlers

import (
	"creds-update/models"
	"fmt"
	"net/http"
)

func (h *RouteHandler) ListCreds(w http.ResponseWriter, r *http.Request) {
	credentials, err := models.GetAllCredentials(h.db)
	if err != nil {
		fmt.Println(err)
	}

	h.render(w, "creds.html", map[string]interface{}{
		"credentials": credentials,
	})
}

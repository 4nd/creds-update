package handlers

import (
	"creds-update/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/render"
	"net/http"
)

func (h *RouteHandler) ListCreds(w http.ResponseWriter, r *http.Request) {
	credentials, err := models.GetAllCredentials(h.db)
	if err != nil {
		fmt.Println(err)
	}

	h.renderTemplate(w, "creds.html", map[string]interface{}{
		"credentials": credentials,
	})
}

func (h *RouteHandler) CompleteCred(w http.ResponseWriter, r *http.Request) {
	var credentialId models.CredentialId

	err := json.NewDecoder(r.Body).Decode(&credentialId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	credential, err := models.CompleteCredential(h.db, credentialId.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, credential)
}

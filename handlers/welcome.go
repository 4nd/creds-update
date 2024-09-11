package handlers

import (
	"creds-update/resources"
	"fmt"
	"log"
	"net/http"
)

func (h *RouteHandler) Welcome(w http.ResponseWriter, r *http.Request) {
	var viewFile = "welcome.html"

	view, ok := resources.Views[viewFile]
	if !ok {
		log.Printf("view %s not found", viewFile)
		return
	}

	fmt.Printf("%v", resources.Views)

	h.render(w, view, map[string]interface{}{
		"Name": "Andy",
	})
}

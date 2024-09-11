package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
)

type RouteHandler struct {
	db *sql.DB
}

func NewRouteHandler(db *sql.DB) *RouteHandler {
	return &RouteHandler{
		db: db,
	}
}

func (h *RouteHandler) render(
	resp http.ResponseWriter,
	view *template.Template,
	data map[string]interface{},
) {
	if viteHot() == true {
		data["viteHot"] = true
	} else {
		data["viteHot"] = false
		data["styles"], data["scripts"] = ParseManifest()
	}
	if err := view.Execute(resp, data); err != nil {
		log.Println(err)
	}
}

func viteHot() bool {
	_, err := os.Stat("./vite-hot")

	return err == nil
}

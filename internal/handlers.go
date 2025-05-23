package handlers

import (
	"context"
	"digital-garden/internal/templates/pages"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	component := pages.Home()
	component.Render(context.Background(), w)
}

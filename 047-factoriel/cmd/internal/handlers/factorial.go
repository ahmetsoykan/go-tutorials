package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ahmetsoykan/go-tutorials/047-factoriel/internal/platform/web"
	"github.com/go-chi/chi"
)

func API() http.Handler {

	app := web.NewApp()

	app.Handle(http.MethodGet, "/v1/{id}", Factorial)

	return app
}

func Factorial(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println(id)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(id)); err != nil {
		log.Println("error writing result", err)
	}
}

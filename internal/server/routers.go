package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm_test_proj/internal/userApi"
	"log/slog"
	"net/http"
)

func routersInit() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10))

	r.Get("/register", userApi.UserRegister)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		slog.Error(err.Error())
	}
}

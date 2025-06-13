package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Mhirii/TaskHub/backend/handlers"
	"github.com/Mhirii/TaskHub/backend/pkg/config"
	"github.com/Mhirii/TaskHub/backend/pkg/infra"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Load Config
	configPath := flag.String("config", "etc/config.yaml", "path to config file")
	flag.Parse()
	cfg, err := config.ParseConfig(*configPath)
	if err != nil {
		panic(err)
	}

	// DB
	_, err = infra.InitDB(*cfg)
	if err != nil {
		panic(err)
	}

	// Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	auth := handlers.NewAuthHandlers()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", auth.Login)
		r.Post("/register", auth.Register)
		r.Post("/refresh", auth.Refresh)
	})
	fmt.Println("server listening on port ", cfg.Server.Port)
	err = http.ListenAndServe(":"+cfg.Server.Port, r)
	if err != nil {
		panic(err)
	}
}

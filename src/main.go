package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/5aradise/media-content-api/src/config"
	"github.com/5aradise/media-content-api/src/internal/database"
	"github.com/5aradise/media-content-api/src/internal/handlers"
	"github.com/5aradise/media-content-api/src/pkg/httpserver"
	swagger "github.com/swaggo/http-swagger/v2"

	_ "github.com/5aradise/media-content-api/src/docs"
	_ "github.com/lib/pq"
)

// @title Media content analysis system
// @version 1.0
// @description RESTful service for analyzing media content

// @contact.name Danyil Rozumovskyi
// @contact.url https://t.me/Danya_Rozum
// @contact.email rozumovskyi.daniil@lll.kpi.ua

// @host localhost:8080
func main() {
	// Load config
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to storage
	conn, err := sql.Open("postgres", config.Cfg.DB.URL)
	if err != nil {
		log.Fatal("can't open sql", err)
	}
	defer conn.Close()

	db := database.Create(conn)

	// Set handlers
	r := http.NewServeMux()

	us := handlers.NewUserService(db)
	r.HandleFunc("POST /users", us.CreateUser)
	r.HandleFunc("GET /users", us.ListUsers)
	r.HandleFunc("GET /users/{id}", us.GetUser)
	r.HandleFunc("PUT /users/{id}", us.UpdateUser)
	r.HandleFunc("DELETE /users/{id}", us.DeleteUser)

	r.HandleFunc("GET /api/", swagger.Handler(
		swagger.URL(fmt.Sprintf("http://localhost:%s/api/doc.json", config.Cfg.Server.Port)),
	))

	// Run server
	server := httpserver.New(
		r,
		httpserver.Port(config.Cfg.Server.Port),
		httpserver.ReadTimeout(4),
		httpserver.IdleTimeout(60),
		httpserver.ErrorLog(log.Default()),
	)

	log.Println("starting server at address:", server.Addr())
	go server.Run()

	// Waiting signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("signal interrupt:", s.String())
	case err := <-server.Notify():
		log.Println("server notify:", err)
	}

	// Shutdown server
	err = server.Shutdown()
	if err != nil {
		log.Println("can't shutdown server:", err)
	}
}

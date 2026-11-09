package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"tempo-server/internal/config"
	"tempo-server/internal/store"
	"tempo-server/internal/subsonic"
)

type HealthResponse struct {
    Message string `json:"message"`
    Status  string `json:"status"`
}

func main() {
	cfg, err:=config.Load()
	if err!=nil {
		log.Fatal(err)
	}
	ctx:=context.Background()
	db,err:=store.NewPostgres(ctx, cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	mux := http.NewServeMux()
	log.Println("Server running on :" + cfg.Port)

	mux.HandleFunc("GET /rest/ping", subsonic.Ping)
	mux.HandleFunc("GET /rest/ping.view", subsonic.Ping)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HealthResponse{
        Message: "API is running",
        Status:  "healthy",
    })
	})

	mux.HandleFunc("GET /health/db", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(r.Context()); err != nil {
			http.Error(w, "Database unavailable", http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Database connected"))
	})


	handler:=subsonic.AuthMiddleWare(cfg.Username, cfg.Password)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler(mux)))

}
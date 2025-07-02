package main

import (
	"context"
	"embed"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/byterotom/infinity-play/config"
	"github.com/byterotom/infinity-play/internal/auth"
	"github.com/byterotom/infinity-play/internal/db"
	"github.com/byterotom/infinity-play/internal/web"
	"github.com/byterotom/infinity-play/pkg"
)

//go:embed static/**
var staticFiles embed.FS

func main() {
	ctx := context.Background()
	env := config.LoadConfig()
	r2 := pkg.NewR2(env)

	conn := db.ConnectDB(ctx, env)
	defer conn.Close(ctx)

	auth.JwtSecret = []byte(env.JwtSecret)

	mux := web.NewInfinityMux(r2, conn, &staticFiles)

	srv := &http.Server{
		Addr:    ":6969",
		Handler: mux,
	}
	log.Println("infinity server running on 6969")
	go srv.ListenAndServe()

	sigChan := make(chan os.Signal, 1)
	// notify the channel in specified signals
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	srv.Shutdown(ctx)

}

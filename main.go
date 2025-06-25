package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/byterotom/infinity-play/config"
	"github.com/byterotom/infinity-play/internal/db"
	"github.com/byterotom/infinity-play/internal/web"
	"github.com/byterotom/infinity-play/pkg"
)

//go:embed templates/*
var content embed.FS

func main() {

	env := config.LoadConfig()
	r2 := pkg.NewR2(env)

	conn := db.ConnectDB()

	defer conn.Close()

	mux := web.NewInfinityMux(&content, r2, conn)

	log.Println("infinity server running on 6969")
	http.ListenAndServe(":6969", mux)

}

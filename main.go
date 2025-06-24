package main

import (
	"embed"
	"net/http"

	"github.com/byterotom/infinity-play/internal/web"
)

//go:embed templates/*
var content embed.FS

func main() {

	mux := web.NewInfinityMux(&content)

	http.ListenAndServe(":8080", mux)
}

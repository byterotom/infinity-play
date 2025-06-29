package main

import (
	"context"
	"log"
	"net/http"

	"github.com/byterotom/infinity-play/config"
	"github.com/byterotom/infinity-play/internal/auth"
	"github.com/byterotom/infinity-play/internal/db"
	"github.com/byterotom/infinity-play/internal/web"
	"github.com/byterotom/infinity-play/pkg"
)

func main() {
	ctx := context.Background()
	env := config.LoadConfig()
	r2 := pkg.NewR2(env)

	conn := db.ConnectDB(ctx, env)
	defer conn.Close(ctx)

	auth.JwtSecret = []byte(env.JwtSecret)

	mux := web.NewInfinityMux(r2, conn)

	log.Println("infinity server running on 6969")
	http.ListenAndServe(":6969", mux)

}

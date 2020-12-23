package main

import (
	"fmt"
	"net/http"

	"github.com/ryuzaki01/go-ms/encrypt/app/config"
	_ "github.com/ryuzaki01/go-ms/encrypt/app/controllers"
	"github.com/ryuzaki01/go-ms/encrypt/app/logs"
)

func main() {
	cfg := config.NewConfig()
	logs.Debug.Print("[config] " + cfg.String())
	logs.Info.Printf("[service] listening on port %v", cfg.Port)
	logs.Fatal.Print(http.ListenAndServe(":"+fmt.Sprint(cfg.Port), nil))
}
package main

import (
	"fmt"
	"lms/server"
)

func main() {
	cfg := server.GetConfig()
	r := server.NewRouter(cfg)
	r.Run(fmt.Sprintf(":%d", cfg.Port))
}

package main

import (
	"fmt"
	"log"

	"github.com/Levisogembo/flowforge/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Environment:", cfg.Env)
	fmt.Println("Address:", cfg.Addr)
	fmt.Println("Database host:", cfg.DB.Host)
	fmt.Println("Database name:", cfg.DB.Name)
	fmt.Println("Maximum connections:", cfg.DB.MaxConns)
}

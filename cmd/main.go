package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jim124/comfy-backend/database"
	"github.com/sethvargo/go-envconfig"
)

type MyConfig struct {
	Port  string `env:"PORT"`
	DbUrl string `env:"DB_URL"`
}

func main() {
	ctx := context.Background()
	var c MyConfig
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}
	database.Init(c.DbUrl)
	server := gin.Default()
	// server.Use()
	server.Run(fmt.Sprintf(":%v", c.Port))
	fmt.Printf("server is running on port: %v", c.Port)
}

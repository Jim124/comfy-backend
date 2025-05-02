package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jim124/comfy-backend/database"
	"github.com/sethvargo/go-envconfig"
)

type MyConfig struct {
	Port        string `env:"PORT"`
	DbUrl       string `env:"DB_URL"`
	TokenSecret string `env:"TOKEN_SECRET"`
}

func main() {
	ctx := context.Background()
	var c MyConfig
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)

	}
	database.Init(c.DbUrl)
	server := gin.Default()
	// cors config
	// server.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
	// 	AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
	// 	AllowCredentials: true, // Enable cookies/auth
	// }))
	server.Use(cors.Default())
	server.Run(fmt.Sprintf(":%v", c.Port))
	log.Printf("server is running on port: %v", c.Port)
}

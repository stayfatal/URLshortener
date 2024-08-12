package main

import (
	"log"
	"url/internal/middleware"
	"url/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	srv := server.NewServer()

	router.POST("/register", srv.CreateUserHandler)
	router.POST("/login", srv.LoginHandler)
	router.GET("/:link", srv.RedirectHandler)

	shortener := router.Group("/")
	shortener.Use(middleware.Authentication())

	shortener.POST("/shorten", srv.ShortenHandler)

	log.Fatal(router.Run(":8080"))
}

package main

import "github.com/anjush-bhargavan/password_generator/pkg/server"

func main() {
	router := server.Init()

	router.LoadHTMLGlob("templates/*.html")

	router.Run(":8080")
}
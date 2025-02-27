package main

import (
	"admin-v1/app"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)


// @title CRM
// @version 1.0
// @host 192.168.0.121:8000
// @BasePath /api/v1
func main() {
	err := godotenv.Load()
	
	if err != nil {
		panic(err)
	}

	route := app.Init()

	port := os.Getenv("PORT")

	ln, err := net.Listen("tcp", "0.0.0.0:" + port)

	if err != nil {
		panic(err)
	}

	_ = http.Serve(ln, route)
}
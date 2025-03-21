package main

import (
	"admin-v1/app"
	"admin-v1/app/helpers"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var redisContext = context.Background()

// @title CRM
// @version 1.0
// @host localhost:1006
// @BasePath /api/v1
func main() {
	helpers.Ctx = redisContext

	err := godotenv.Load()

	helpers.Redis = redis.NewClient(&redis.Options{
		Addr:         "172.26.168.7:6379",
		PoolSize:     20,
		MinIdleConns: 5,  
	})

	pong, err := helpers.Redis.Ping(helpers.Ctx).Result()

	if err != nil {
		fmt.Println("Không thể kết nối Redis:", err)
		return
	}

	fmt.Println("Kết nối Redis thành công:", pong)
	
	if err != nil {
		panic(err)
	}

	route := app.Init()

	port := os.Getenv("PORT")

	ln, err := net.Listen("tcp", "localhost:" + port)

	if err != nil {
		panic(err)
	}

	_ = http.Serve(ln, route)
}
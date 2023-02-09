package main

import (
	"time"

	"github.com/labstack/echo"
	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})

	if err := Open(); err != nil {
		panic("db failed")
	}

	ttl := time.Second * 100
	s := NewStore(NewRedisCache(client, ttl))

	e := echo.New()
	e.GET("/user/:id", s.GetUser)
	e.Logger.Fatal(e.Start(":3003"))

}

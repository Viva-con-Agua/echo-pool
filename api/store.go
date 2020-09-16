package api

import (
	"encoding/gob"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func RedisSession() echo.MiddlewareFunc {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})

	redis, err := NewRedisStore(client)

	if err != nil {
		log.Fatal("failed to create redis store: ", err)
	}
	gob.Register(&UserSession{})
	gob.Register(&Additional{})
	log.Println("Redis successfully connected!")
	return session.Middleware(redis)
}

package auth

import (
	"encoding/gob"
	"log"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func RedisSession(host string, port string) echo.MiddlewareFunc {
	client := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})

	redis, err := NewRedisStore(client)

	if err != nil {
		log.Fatal("failed to create redis store: ", err)
	}
	gob.Register(&User{})
	gob.Register(&M{})
	log.Println("Redis successfully connected!")
	return session.Middleware(redis)
}

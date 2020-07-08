package auth

import (
	"encoding/gob"
	"log"

	"github.com/Viva-con-Agua/echo-pool/config"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func RedisSession() echo.MiddlewareFunc {
	client := redis.NewClient(&redis.Options{
		Addr: config.Config.Redis.Url,
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

package redis

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"web_start/app/config"
)

var Redis *redis.Client

func Init() {
	conf := config.Conf.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.Db, // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		logrus.Error("redis connection error: ", err)
		panic(err)
	}
	logrus.Info("redis connection success")
	Redis = client
}

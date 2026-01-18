package redis

import (
	"github.com/webcore-go/webcore/infra/config"
	"github.com/webcore-go/webcore/port"
)

type RedisLoader struct {
	Redis *Redis
	name  string
}

func (a *RedisLoader) SetName(name string) {
	a.name = name
}

func (a *RedisLoader) Name() string {
	return a.name
}

func (l *RedisLoader) Init(args ...any) (port.Library, error) {
	config := args[0].(config.RedisConfig)
	redis := NewRedis(config)
	err := redis.Install(args...)
	if err != nil {
		return nil, err
	}

	redis.Connect()

	l.Redis = redis
	return redis, nil
}

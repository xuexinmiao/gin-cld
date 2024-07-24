package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client
var err error

func Init() (err error) {
	addr := fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetInt("redis.host"))
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})
	return nil
}

func Close() {
	_ = rdb.Close()
}

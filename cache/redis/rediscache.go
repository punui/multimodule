package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type redisCacheConfig struct {
	ADDRESS   string `mapstructure:"CLUSTER_ADDRESS"`
	PASSWORD  string `mapstructure:"CLUSTER_PASSWORD"`
	UseSSL    bool   `mapstructure:"CLUSTER_USE_SSL"`
	RemoveTtl int
}

type Cache struct {
	redis *redis.Client
	ttl   int64
}

func (c *Cache) Init(conf *redisCacheConfig) {

	c.redis = redis.NewClient(&redis.Options{
		Addr:     conf.ADDRESS,
		Password: conf.PASSWORD,
		DB:       0,
	})
	c.ttl = int64(conf.RemoveTtl)
}
func (c *Cache) GetKey(key string) int {
	val, err := c.redis.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	if val != "" {
		n, _ := strconv.Atoi(val)
		return n
	} else {
		return 1
	}
}

func (c *Cache) SetKey(key string, value int) {

	err := c.redis.Set(key, value, time.Duration(time.Duration.Microseconds()*c.ttl))

	if err != nil {
		println(err)
	}
}

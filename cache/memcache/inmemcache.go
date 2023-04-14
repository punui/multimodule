package memcache

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"strconv"
	"time"
)

type Cache struct {
	inMemCache *bigcache.BigCache
	ttl        int
}

type InMemCacheConfig struct {
	Shards             int
	LifeWindow         time.Duration
	CleanWindow        time.Duration
	MaxEntriesInWindow int
	MaxEntrySize       int
	Verbose            bool
	HardMaxCacheSize   int
	OnRemove           func(string2 string, data []byte)
	OnRemoveWithReason func(string2 string, data []byte, reason bigcache.RemoveReason)
	RemoveTtl          int
}

func (c *Cache) Init(conf *InMemCacheConfig) {

	config := bigcache.Config{
		Shards:             conf.Shards,
		LifeWindow:         conf.LifeWindow,
		CleanWindow:        conf.CleanWindow,
		MaxEntriesInWindow: conf.MaxEntriesInWindow,
		MaxEntrySize:       conf.MaxEntrySize,
		Verbose:            conf.Verbose,
		HardMaxCacheSize:   conf.HardMaxCacheSize,
		OnRemove:           conf.OnRemove,
		OnRemoveWithReason: conf.OnRemoveWithReason,
	}

	c.inMemCache, err = bigcache.NewBigCache(config)
	c.ttl = conf.RemoveTtl
}

func (c *Cache) GetKey(key string) []byte {
	val, err := c.inMemCache.Get(key)
	if err != nil {
		fmt.Println(err)
	}
	return val
}

func (c *Cache) SetKey(key string, value []byte) {

	err := c.inMemCache.Set(key, value)

	if err != nil {
		println(err)
	}
}

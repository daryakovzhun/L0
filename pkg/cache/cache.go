package cache

import (
	"L0WB/internal/domain"
	"sync"
	"time"
)

type CacheItem struct {
	value        domain.Order
	timeCreation time.Time
	timeDuration time.Duration
}

type Cache struct {
	cache map[string]CacheItem
	mu    *sync.Mutex
}

func New() *Cache {
	cacheNew := Cache{
		cache: make(map[string]CacheItem),
		mu:    new(sync.Mutex),
	}
	go cacheNew.TimeExpireTask()
	return &cacheNew
}

func (c *Cache) Set(key string, value domain.Order, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = CacheItem{
		value:        value,
		timeCreation: time.Now(),
		timeDuration: ttl,
	}
}

func (c *Cache) SetFew(ords map[string]domain.Order, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range ords {
		c.cache[k] = CacheItem{
			value:        v,
			timeCreation: time.Now(),
			timeDuration: ttl,
		}
	}
}

func (c *Cache) Get(key string) (domain.Order, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	k, ok := c.cache[key]

	if !ok {
		//return domain.Order{Order_uid: ""}, errors.New("invalid key")
	}
	return k.value, nil
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.cache, key)
}

func (c *Cache) TimeExpireTask() {
	for {
		c.DeleteExpired()
		time.Sleep(1 * time.Second)
	}
}

func (c *Cache) DeleteExpired() {
	for key, value := range c.cache {
		if time.Since(value.timeCreation) > value.timeDuration {
			c.Delete(key)
		}
	}
}

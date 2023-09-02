package cache

import "errors"

var ErrKeyMissing = errors.New("cache key missing")

type CachePolicy interface {
	Add(CacheKey)
	Remove(CacheKey)
	Access(CacheKey)
	Victim() CacheKey
}

func (c *Cache) Add(key CacheKey, value string) {
	if c.currSize == c.maxSize {
		victimKey := c.policy.Victim()
		delete(c.data, victimKey)
		c.currSize -= 1
	}
	c.policy.Add(key)
	c.data[key] = value
	c.currSize += 1
}

func (c *Cache) Remove(key CacheKey) (*string, error) {
	if value, ok := c.data[key]; ok {
		c.policy.Access(key)
		return &value, nil
	}
	return nil, ErrKeyMissing
}

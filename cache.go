package cache

type CacheKey string

type CacheData map[CacheKey]string

type Cache struct {
	maxSize  int
	currSize int
	policy   CachePolicy
	data     CacheData
}

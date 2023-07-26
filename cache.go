package in_memory_cache

type Cache map[string]interface{}

func New() Cache {
	newCache := make(Cache)
	return newCache
}

func (c *Cache) Set(key string, value interface{}) {
	cache := *c
	cache[key] = value
}

func (c *Cache) Delete(key string) {
	cache := *c
	delete(cache, key)
}

func (c *Cache) Get(key string) interface{} {
	cache := *c
	return cache[key]
}

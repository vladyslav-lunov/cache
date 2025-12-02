package github.com/vladyslav-lunov/cache

type Cache struct {
	m map[string]interface{}
}

func New() *Cache {
	return &Cache{
		m: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.m[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	value, ok := c.m[key]
	return value, ok
}

func (c *Cache) Delete(key string) {
	delete(c.m, key)
}

func (c *Cache) Clear() {
	c.m = make(map[string]interface{})
}

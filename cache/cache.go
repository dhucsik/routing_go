package cache

type Cache interface {
	Add(key, value string)
	Get(key string) (value string, ok bool)
	Len() int
}

func NewMyCache() MyCache {
	m := MyCache{
		make(map[string]string),
	}
	return m
}

type MyCache struct {
	cache map[string]string
}

func (m MyCache) Add(key, value string) {
	m.cache[key] = value
}

func (m MyCache) Get(key string) (string, bool) {
	value, ok := m.cache[key]
	return value, ok
}

func (m MyCache) Len() int {
	return len(m.cache)
}

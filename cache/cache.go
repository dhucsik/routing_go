package cache

type Cache interface {
	Add(key, value string)
	Get(key string) (value string, ok bool)
	Len() int
}

type MyCache map[string]string

func (m MyCache) Add(key, value string) {
	m[key] = value
}

func (m MyCache) Get(key string) (string, bool) {
	value, ok := m[key]
	return value, ok
}

func (m MyCache) Len() int {
	return len(m)
}

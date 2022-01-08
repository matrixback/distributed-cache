package cache

type Cache interface {
	Get(string) ([]byte, error)
	Set(string, []byte) error
	Del(string) error
}

type MemoryCache struct {
	hash map[string][]byte
}

func NewMemoryCache() *MemoryCache {
	h := make(map[string][]byte)
	return &MemoryCache{hash: h}
}

func(mc *MemoryCache) Get(key string) ([]byte, error) {
	val, _ := mc.hash[key]
	return val, nil
}

func(mc *MemoryCache) Set(key string, val []byte) error {
	mc.hash[key] = val
	return nil
}

func(mc *MemoryCache) Del(key string) error {
	return nil
}

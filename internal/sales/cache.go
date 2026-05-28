package sales

import "sync"

const (
	cacheKeyTotalVolume              = "total_volume"
	cacheKeyVolumeByCenter           = "volume_by_center"
	cacheKeyModelPercentagesByCenter = "model_percentages_by_center"
)

type cache struct {
	mu     sync.RWMutex
	values map[string]any
}

func newCache() *cache {

	return &cache{
		values: make(map[string]any),
	}
}

func (c *cache) Get(key string) (any, bool) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.values[key]
	return value, ok
}

func (c *cache) Set(key string, value any) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.values[key] = value
}

func (c *cache) Clear() {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.values = make(map[string]any)
}

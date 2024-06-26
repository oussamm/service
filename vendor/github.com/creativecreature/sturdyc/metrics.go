package sturdyc

type MetricsRecorder interface {
	// CacheHit is called for every key that results in a cache hit.
	CacheHit()
	// CacheMiss is called for every key that results in a cache miss.
	CacheMiss()
	// ForcedEviction is called when the cache reaches its capacity, and has to
	// evict keys in order to write a new one.
	ForcedEviction()
	// EntiresEvicted is called when the cache evicts keys from a shard.
	EntriesEvicted(int)
	// ShardIndex is called to report which shard it was that performed an operation.
	ShardIndex(int)
	// CacheBatchRefreshSize is called to report the size of the batch refresh.
	CacheBatchRefreshSize(size int)
	// ObserveCacheSize is called to report the size of the cache.metri
	ObserveCacheSize(callback func() int)
}

type DistributedMetricsRecorder interface {
	MetricsRecorder
	// DistributedCacheHit is called for every key that results in a cache hit.
	DistributedCacheHit()
	// DistributedCacheHit is called for every key that results in a cache miss.
	DistributedCacheMiss()
	// DistributedFallback is called when you are using a distributed storage
	// with early refreshes, and the call for a value was supposed to refresh it,
	// but the call failed. When that happens, the cache fallbacks to the latest
	// value from the distributed storage.
	DistributedFallback()
}

type distributedMetricsRecorder struct {
	MetricsRecorder
}

func (d *distributedMetricsRecorder) DistributedCacheHit() {}

func (d *distributedMetricsRecorder) DistributedCacheMiss() {}

func (d *distributedMetricsRecorder) DistributedFallback() {}

func (s *shard[T]) reportForcedEviction() {
	if s.metricsRecorder == nil {
		return
	}
	s.metricsRecorder.ForcedEviction()
}

func (s *shard[T]) reportEntriesEvicted(n int) {
	if s.metricsRecorder == nil {
		return
	}
	s.metricsRecorder.EntriesEvicted(n)
}

// reportCacheHits is used to report cache hits and misses to the metrics recorder.
func (c *Client[T]) reportCacheHits(cacheHit bool) {
	if c.metricsRecorder == nil {
		return
	}
	if !cacheHit {
		c.metricsRecorder.CacheMiss()
		return
	}
	c.metricsRecorder.CacheHit()
}

func (c *Client[T]) reportShardIndex(index int) {
	if c.metricsRecorder == nil {
		return
	}
	c.metricsRecorder.ShardIndex(index)
}

func (c *Client[T]) reportBatchRefreshSize(n int) {
	if c.metricsRecorder == nil {
		return
	}
	c.metricsRecorder.CacheBatchRefreshSize(n)
}

func (c *Client[T]) reportDistributedCacheHit(cacheHit bool) {
	if c.metricsRecorder == nil {
		return
	}
	if !cacheHit {
		c.metricsRecorder.DistributedCacheMiss()
		return
	}
	c.metricsRecorder.DistributedCacheHit()
}

func (c *Client[T]) reportDistributedStaleFallback() {
	if c.metricsRecorder == nil {
		return
	}
	c.metricsRecorder.DistributedFallback()
}

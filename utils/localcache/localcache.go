package localcache

import (
	"errors"

	"github.com/coocood/freecache"
)

type LocalCache interface {
	Get(key string) ([]byte, error)
	SetWithExpire(key string, value []byte, expireSecond int) error
	GetOrSetWithExpire(key string, value []byte, expireSecond int) ([]byte, error)
	Delete(key string) error
	GetStatistics() map[string]interface{}
	EntryCount() int64
	EvacuateCount() int64
	ExpiredCount() int64
	HitCount() int64
	HitRate() float64
	LookupCount() int64
	MissCount() int64
}

type cacheRepo struct {
	cacheClient *freecache.Cache
}

func NewLocalCacheRepository(cache *freecache.Cache) LocalCache {
	return &cacheRepo{cacheClient: cache}
}

func (s *cacheRepo) Get(key string) ([]byte, error) {
	return s.cacheClient.Get([]byte(key))
}

func (s *cacheRepo) SetWithExpire(key string, value []byte, expireSecond int) error {
	return s.cacheClient.Set([]byte(key), value, expireSecond)
}

func (s *cacheRepo) GetOrSetWithExpire(key string, value []byte, expireSecond int) ([]byte, error) {
	return s.cacheClient.GetOrSet([]byte(key), value, expireSecond)
}

func (s *cacheRepo) Delete(key string) error {
	if s.cacheClient.Del([]byte(key)) {
		return nil
	}
	return errors.New("delete failed")
}

func (s *cacheRepo) EntryCount() int64 {
	return s.cacheClient.EntryCount()
}

func (s *cacheRepo) EvacuateCount() int64 {
	return s.cacheClient.EntryCount()
}

func (s *cacheRepo) ExpiredCount() int64 {
	return s.cacheClient.ExpiredCount()
}

func (s *cacheRepo) HitCount() int64 {
	return s.cacheClient.HitCount()
}

func (s *cacheRepo) HitRate() float64 {
	return s.cacheClient.HitRate()
}

func (s *cacheRepo) LookupCount() int64 {
	return s.cacheClient.LookupCount()
}

func (s *cacheRepo) MissCount() int64 {
	return s.cacheClient.MissCount()
}

func (s *cacheRepo) GetStatistics() map[string]interface{} {
	return map[string]interface{}{
		"average_access_time": s.cacheClient.AverageAccessTime(),
		"entry_count":         s.cacheClient.EntryCount(),
		"evacuate_count":      s.cacheClient.EvacuateCount(),
		"expired_count":       s.cacheClient.ExpiredCount(),
		"hit_count":           s.cacheClient.HitCount(),
		"hit_rate":            s.cacheClient.HitRate(),
		"lookup_count":        s.cacheClient.LookupCount(),
		"miss_count":          s.cacheClient.MissCount(),
		"overwrite_count":     s.cacheClient.OverwriteCount(),
		"touched_count":       s.cacheClient.TouchedCount(),
	}
}

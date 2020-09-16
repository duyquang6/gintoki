package middleware

import (
	"context"
	"gintoki/utils/localcache"

	custom_metric "gintoki/utils/prometheus"

	"google.golang.org/grpc"
)

func MetricInterceptor(cacheRepo localcache.LocalCache) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		res, err := handler(ctx, req)
		custom_metric.CacheHit.Set(float64(cacheRepo.HitCount()))
		custom_metric.CacheMiss.Set(float64(cacheRepo.MissCount()))
		custom_metric.CacheEntry.Set(float64(cacheRepo.EntryCount()))
		custom_metric.CacheEvacuate.Set(float64(cacheRepo.EvacuateCount()))
		custom_metric.CacheHitRate.Set(cacheRepo.HitRate())
		custom_metric.LookupCache.Set(float64(cacheRepo.LookupCount()))
		custom_metric.LookupCache.Set(float64(cacheRepo.ExpiredCount()))
		return res, err
	}
}

package cache

import (
	"context"
	"strings"
	"time"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"

	"weaving_net/internal/model"
)

const (
	// cache prefix key, must end with a colon
	educationsCachePrefixKey = "educations:"
	// EducationsExpireTime expire time
	EducationsExpireTime = 5 * time.Minute
)

var _ EducationsCache = (*educationsCache)(nil)

// EducationsCache cache interface
type EducationsCache interface {
	Set(ctx context.Context, id uint64, data *model.Educations, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Educations, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Educations, error)
	MultiSet(ctx context.Context, data []*model.Educations, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// educationsCache define a cache struct
type educationsCache struct {
	cache cache.Cache
}

// NewEducationsCache new a cache
func NewEducationsCache(cacheType *model.CacheType) EducationsCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Educations{}
		})
		return &educationsCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Educations{}
		})
		return &educationsCache{cache: c}
	}

	return nil // no cache
}

// GetEducationsCacheKey cache key
func (c *educationsCache) GetEducationsCacheKey(id uint64) string {
	return educationsCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *educationsCache) Set(ctx context.Context, id uint64, data *model.Educations, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetEducationsCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *educationsCache) Get(ctx context.Context, id uint64) (*model.Educations, error) {
	var data *model.Educations
	cacheKey := c.GetEducationsCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *educationsCache) MultiSet(ctx context.Context, data []*model.Educations, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetEducationsCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *educationsCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Educations, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetEducationsCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Educations)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Educations)
	for _, id := range ids {
		val, ok := itemMap[c.GetEducationsCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *educationsCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetEducationsCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *educationsCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetEducationsCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

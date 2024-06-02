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
	workexperiencesCachePrefixKey = "workexperiences:"
	// WorkexperiencesExpireTime expire time
	WorkexperiencesExpireTime = 5 * time.Minute
)

var _ WorkexperiencesCache = (*workexperiencesCache)(nil)

// WorkexperiencesCache cache interface
type WorkexperiencesCache interface {
	Set(ctx context.Context, id uint64, data *model.Workexperiences, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Workexperiences, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Workexperiences, error)
	MultiSet(ctx context.Context, data []*model.Workexperiences, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// workexperiencesCache define a cache struct
type workexperiencesCache struct {
	cache cache.Cache
}

// NewWorkexperiencesCache new a cache
func NewWorkexperiencesCache(cacheType *model.CacheType) WorkexperiencesCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Workexperiences{}
		})
		return &workexperiencesCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Workexperiences{}
		})
		return &workexperiencesCache{cache: c}
	}

	return nil // no cache
}

// GetWorkexperiencesCacheKey cache key
func (c *workexperiencesCache) GetWorkexperiencesCacheKey(id uint64) string {
	return workexperiencesCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *workexperiencesCache) Set(ctx context.Context, id uint64, data *model.Workexperiences, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetWorkexperiencesCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *workexperiencesCache) Get(ctx context.Context, id uint64) (*model.Workexperiences, error) {
	var data *model.Workexperiences
	cacheKey := c.GetWorkexperiencesCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *workexperiencesCache) MultiSet(ctx context.Context, data []*model.Workexperiences, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetWorkexperiencesCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *workexperiencesCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Workexperiences, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetWorkexperiencesCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Workexperiences)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Workexperiences)
	for _, id := range ids {
		val, ok := itemMap[c.GetWorkexperiencesCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *workexperiencesCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetWorkexperiencesCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *workexperiencesCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetWorkexperiencesCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

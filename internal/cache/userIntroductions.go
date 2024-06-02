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
	userIntroductionsCachePrefixKey = "userIntroductions:"
	// UserIntroductionsExpireTime expire time
	UserIntroductionsExpireTime = 5 * time.Minute
)

var _ UserIntroductionsCache = (*userIntroductionsCache)(nil)

// UserIntroductionsCache cache interface
type UserIntroductionsCache interface {
	Set(ctx context.Context, id uint64, data *model.UserIntroductions, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.UserIntroductions, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserIntroductions, error)
	MultiSet(ctx context.Context, data []*model.UserIntroductions, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// userIntroductionsCache define a cache struct
type userIntroductionsCache struct {
	cache cache.Cache
}

// NewUserIntroductionsCache new a cache
func NewUserIntroductionsCache(cacheType *model.CacheType) UserIntroductionsCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserIntroductions{}
		})
		return &userIntroductionsCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserIntroductions{}
		})
		return &userIntroductionsCache{cache: c}
	}

	return nil // no cache
}

// GetUserIntroductionsCacheKey cache key
func (c *userIntroductionsCache) GetUserIntroductionsCacheKey(id uint64) string {
	return userIntroductionsCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *userIntroductionsCache) Set(ctx context.Context, id uint64, data *model.UserIntroductions, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetUserIntroductionsCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *userIntroductionsCache) Get(ctx context.Context, id uint64) (*model.UserIntroductions, error) {
	var data *model.UserIntroductions
	cacheKey := c.GetUserIntroductionsCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *userIntroductionsCache) MultiSet(ctx context.Context, data []*model.UserIntroductions, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetUserIntroductionsCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *userIntroductionsCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserIntroductions, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetUserIntroductionsCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.UserIntroductions)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.UserIntroductions)
	for _, id := range ids {
		val, ok := itemMap[c.GetUserIntroductionsCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *userIntroductionsCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserIntroductionsCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *userIntroductionsCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserIntroductionsCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

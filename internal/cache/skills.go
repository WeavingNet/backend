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
	skillsCachePrefixKey = "skills:"
	// SkillsExpireTime expire time
	SkillsExpireTime = 5 * time.Minute
)

var _ SkillsCache = (*skillsCache)(nil)

// SkillsCache cache interface
type SkillsCache interface {
	Set(ctx context.Context, id uint64, data *model.Skills, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Skills, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Skills, error)
	MultiSet(ctx context.Context, data []*model.Skills, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// skillsCache define a cache struct
type skillsCache struct {
	cache cache.Cache
}

// NewSkillsCache new a cache
func NewSkillsCache(cacheType *model.CacheType) SkillsCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Skills{}
		})
		return &skillsCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Skills{}
		})
		return &skillsCache{cache: c}
	}

	return nil // no cache
}

// GetSkillsCacheKey cache key
func (c *skillsCache) GetSkillsCacheKey(id uint64) string {
	return skillsCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *skillsCache) Set(ctx context.Context, id uint64, data *model.Skills, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetSkillsCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *skillsCache) Get(ctx context.Context, id uint64) (*model.Skills, error) {
	var data *model.Skills
	cacheKey := c.GetSkillsCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *skillsCache) MultiSet(ctx context.Context, data []*model.Skills, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetSkillsCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *skillsCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Skills, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetSkillsCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Skills)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Skills)
	for _, id := range ids {
		val, ok := itemMap[c.GetSkillsCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *skillsCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetSkillsCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *skillsCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetSkillsCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

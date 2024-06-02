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
	projectsCachePrefixKey = "projects:"
	// ProjectsExpireTime expire time
	ProjectsExpireTime = 5 * time.Minute
)

var _ ProjectsCache = (*projectsCache)(nil)

// ProjectsCache cache interface
type ProjectsCache interface {
	Set(ctx context.Context, id uint64, data *model.Projects, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Projects, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Projects, error)
	MultiSet(ctx context.Context, data []*model.Projects, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// projectsCache define a cache struct
type projectsCache struct {
	cache cache.Cache
}

// NewProjectsCache new a cache
func NewProjectsCache(cacheType *model.CacheType) ProjectsCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Projects{}
		})
		return &projectsCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Projects{}
		})
		return &projectsCache{cache: c}
	}

	return nil // no cache
}

// GetProjectsCacheKey cache key
func (c *projectsCache) GetProjectsCacheKey(id uint64) string {
	return projectsCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *projectsCache) Set(ctx context.Context, id uint64, data *model.Projects, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetProjectsCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *projectsCache) Get(ctx context.Context, id uint64) (*model.Projects, error) {
	var data *model.Projects
	cacheKey := c.GetProjectsCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *projectsCache) MultiSet(ctx context.Context, data []*model.Projects, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetProjectsCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *projectsCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Projects, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetProjectsCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Projects)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Projects)
	for _, id := range ids {
		val, ok := itemMap[c.GetProjectsCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *projectsCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetProjectsCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *projectsCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetProjectsCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

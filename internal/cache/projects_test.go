package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"weaving_net/internal/model"
)

func newProjectsCache() *gotest.Cache {
	record1 := &model.Projects{}
	record1.ID = 1
	record2 := &model.Projects{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewProjectsCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_projectsCache_Set(t *testing.T) {
	c := newProjectsCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Projects)
	err := c.ICache.(ProjectsCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(ProjectsCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_projectsCache_Get(t *testing.T) {
	c := newProjectsCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Projects)
	err := c.ICache.(ProjectsCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(ProjectsCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(ProjectsCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_projectsCache_MultiGet(t *testing.T) {
	c := newProjectsCache()
	defer c.Close()

	var testData []*model.Projects
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Projects))
	}

	err := c.ICache.(ProjectsCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(ProjectsCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Projects))
	}
}

func Test_projectsCache_MultiSet(t *testing.T) {
	c := newProjectsCache()
	defer c.Close()

	var testData []*model.Projects
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Projects))
	}

	err := c.ICache.(ProjectsCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_projectsCache_Del(t *testing.T) {
	c := newProjectsCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Projects)
	err := c.ICache.(ProjectsCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_projectsCache_SetCacheWithNotFound(t *testing.T) {
	c := newProjectsCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Projects)
	err := c.ICache.(ProjectsCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewProjectsCache(t *testing.T) {
	c := NewProjectsCache(&model.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewProjectsCache(&model.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewProjectsCache(&model.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}

package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"weaving_net/internal/model"
)

func newEducationsCache() *gotest.Cache {
	record1 := &model.Educations{}
	record1.ID = 1
	record2 := &model.Educations{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewEducationsCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_educationsCache_Set(t *testing.T) {
	c := newEducationsCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Educations)
	err := c.ICache.(EducationsCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(EducationsCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_educationsCache_Get(t *testing.T) {
	c := newEducationsCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Educations)
	err := c.ICache.(EducationsCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(EducationsCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(EducationsCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_educationsCache_MultiGet(t *testing.T) {
	c := newEducationsCache()
	defer c.Close()

	var testData []*model.Educations
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Educations))
	}

	err := c.ICache.(EducationsCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(EducationsCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Educations))
	}
}

func Test_educationsCache_MultiSet(t *testing.T) {
	c := newEducationsCache()
	defer c.Close()

	var testData []*model.Educations
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Educations))
	}

	err := c.ICache.(EducationsCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_educationsCache_Del(t *testing.T) {
	c := newEducationsCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Educations)
	err := c.ICache.(EducationsCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_educationsCache_SetCacheWithNotFound(t *testing.T) {
	c := newEducationsCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Educations)
	err := c.ICache.(EducationsCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewEducationsCache(t *testing.T) {
	c := NewEducationsCache(&model.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewEducationsCache(&model.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewEducationsCache(&model.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}

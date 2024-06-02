package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"weaving_net/internal/model"
)

func newUserIntroductionsCache() *gotest.Cache {
	record1 := &model.UserIntroductions{}
	record1.ID = 1
	record2 := &model.UserIntroductions{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewUserIntroductionsCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_userIntroductionsCache_Set(t *testing.T) {
	c := newUserIntroductionsCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserIntroductions)
	err := c.ICache.(UserIntroductionsCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(UserIntroductionsCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_userIntroductionsCache_Get(t *testing.T) {
	c := newUserIntroductionsCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserIntroductions)
	err := c.ICache.(UserIntroductionsCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserIntroductionsCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(UserIntroductionsCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_userIntroductionsCache_MultiGet(t *testing.T) {
	c := newUserIntroductionsCache()
	defer c.Close()

	var testData []*model.UserIntroductions
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserIntroductions))
	}

	err := c.ICache.(UserIntroductionsCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserIntroductionsCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.UserIntroductions))
	}
}

func Test_userIntroductionsCache_MultiSet(t *testing.T) {
	c := newUserIntroductionsCache()
	defer c.Close()

	var testData []*model.UserIntroductions
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserIntroductions))
	}

	err := c.ICache.(UserIntroductionsCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userIntroductionsCache_Del(t *testing.T) {
	c := newUserIntroductionsCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserIntroductions)
	err := c.ICache.(UserIntroductionsCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userIntroductionsCache_SetCacheWithNotFound(t *testing.T) {
	c := newUserIntroductionsCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserIntroductions)
	err := c.ICache.(UserIntroductionsCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewUserIntroductionsCache(t *testing.T) {
	c := NewUserIntroductionsCache(&model.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewUserIntroductionsCache(&model.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewUserIntroductionsCache(&model.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}

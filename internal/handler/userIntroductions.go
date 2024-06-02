package handler

import (
	"errors"
	"math"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"github.com/zhufuyi/sponge/pkg/gin/middleware"
	"github.com/zhufuyi/sponge/pkg/gin/response"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/utils"

	"weaving_net/internal/cache"
	"weaving_net/internal/dao"
	"weaving_net/internal/ecode"
	"weaving_net/internal/model"
	"weaving_net/internal/types"
)

var _ UserIntroductionsHandler = (*userIntroductionsHandler)(nil)

// UserIntroductionsHandler defining the handler interface
type UserIntroductionsHandler interface {
	Create(c *gin.Context)
	DeleteByID(c *gin.Context)
	DeleteByIDs(c *gin.Context)
	UpdateByID(c *gin.Context)
	GetByID(c *gin.Context)
	GetByCondition(c *gin.Context)
	ListByIDs(c *gin.Context)
	ListByLastID(c *gin.Context)
	List(c *gin.Context)
}

type userIntroductionsHandler struct {
	iDao dao.UserIntroductionsDao
}

// NewUserIntroductionsHandler creating the handler interface
func NewUserIntroductionsHandler() UserIntroductionsHandler {
	return &userIntroductionsHandler{
		iDao: dao.NewUserIntroductionsDao(
			model.GetDB(),
			cache.NewUserIntroductionsCache(model.GetCacheType()),
		),
	}
}

// Create a record
// @Summary create userIntroductions
// @Description submit information to create userIntroductions
// @Tags userIntroductions
// @accept json
// @Produce json
// @Param data body types.CreateUserIntroductionsRequest true "userIntroductions information"
// @Success 200 {object} types.CreateUserIntroductionsRespond{}
// @Router /api/v1/userIntroductions [post]
// @Security BearerAuth
func (h *userIntroductionsHandler) Create(c *gin.Context) {
	form := &types.CreateUserIntroductionsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	userIntroductions := &model.UserIntroductions{}
	err = copier.Copy(userIntroductions, form)
	if err != nil {
		response.Error(c, ecode.ErrCreateUserIntroductions)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, userIntroductions)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": userIntroductions.ID})
}

// DeleteByID delete a record by id
// @Summary delete userIntroductions
// @Description delete userIntroductions by id
// @Tags userIntroductions
// @accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.DeleteUserIntroductionsByIDRespond{}
// @Router /api/v1/userIntroductions/{id} [delete]
// @Security BearerAuth
func (h *userIntroductionsHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getUserIntroductionsIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	err := h.iDao.DeleteByID(ctx, id)
	if err != nil {
		logger.Error("DeleteByID error", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// DeleteByIDs delete records by batch id
// @Summary delete userIntroductionss
// @Description delete userIntroductionss by batch id
// @Tags userIntroductions
// @Param data body types.DeleteUserIntroductionssByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.DeleteUserIntroductionssByIDsRespond{}
// @Router /api/v1/userIntroductions/delete/ids [post]
// @Security BearerAuth
func (h *userIntroductionsHandler) DeleteByIDs(c *gin.Context) {
	form := &types.DeleteUserIntroductionssByIDsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.DeleteByIDs(ctx, form.IDs)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// UpdateByID update information by id
// @Summary update userIntroductions
// @Description update userIntroductions information by id
// @Tags userIntroductions
// @accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdateUserIntroductionsByIDRequest true "userIntroductions information"
// @Success 200 {object} types.UpdateUserIntroductionsByIDRespond{}
// @Router /api/v1/userIntroductions/{id} [put]
// @Security BearerAuth
func (h *userIntroductionsHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getUserIntroductionsIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdateUserIntroductionsByIDRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	form.ID = id

	userIntroductions := &model.UserIntroductions{}
	err = copier.Copy(userIntroductions, form)
	if err != nil {
		response.Error(c, ecode.ErrUpdateByIDUserIntroductions)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, userIntroductions)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// GetByID get a record by id
// @Summary get userIntroductions detail
// @Description get userIntroductions detail by id
// @Tags userIntroductions
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetUserIntroductionsByIDRespond{}
// @Router /api/v1/userIntroductions/{id} [get]
// @Security BearerAuth
func (h *userIntroductionsHandler) GetByID(c *gin.Context) {
	idStr, id, isAbort := getUserIntroductionsIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	userIntroductions, err := h.iDao.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("GetByID not found", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
			response.Error(c, ecode.NotFound)
		} else {
			logger.Error("GetByID error", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
			response.Output(c, ecode.InternalServerError.ToHTTPCode())
		}
		return
	}

	data := &types.UserIntroductionsObjDetail{}
	err = copier.Copy(data, userIntroductions)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDUserIntroductions)
		return
	}
	data.ID = idStr

	response.Success(c, gin.H{"userIntroductions": data})
}

// GetByCondition get a record by condition
// @Summary get userIntroductions by condition
// @Description get userIntroductions by condition
// @Tags userIntroductions
// @Param data body types.Conditions true "query condition"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetUserIntroductionsByConditionRespond{}
// @Router /api/v1/userIntroductions/condition [post]
// @Security BearerAuth
func (h *userIntroductionsHandler) GetByCondition(c *gin.Context) {
	form := &types.GetUserIntroductionsByConditionRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	err = form.Conditions.CheckValid()
	if err != nil {
		logger.Warn("Parameters error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	userIntroductions, err := h.iDao.GetByCondition(ctx, &form.Conditions)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("GetByCondition not found", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
			response.Error(c, ecode.NotFound)
		} else {
			logger.Error("GetByCondition error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
			response.Output(c, ecode.InternalServerError.ToHTTPCode())
		}
		return
	}

	data := &types.UserIntroductionsObjDetail{}
	err = copier.Copy(data, userIntroductions)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDUserIntroductions)
		return
	}
	data.ID = utils.Uint64ToStr(userIntroductions.ID)

	response.Success(c, gin.H{"userIntroductions": data})
}

// ListByIDs list of records by batch id
// @Summary list of userIntroductionss by batch id
// @Description list of userIntroductionss by batch id
// @Tags userIntroductions
// @Param data body types.ListUserIntroductionssByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.ListUserIntroductionssByIDsRespond{}
// @Router /api/v1/userIntroductions/list/ids [post]
// @Security BearerAuth
func (h *userIntroductionsHandler) ListByIDs(c *gin.Context) {
	form := &types.ListUserIntroductionssByIDsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	userIntroductionsMap, err := h.iDao.GetByIDs(ctx, form.IDs)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	userIntroductionss := []*types.UserIntroductionsObjDetail{}
	for _, id := range form.IDs {
		if v, ok := userIntroductionsMap[id]; ok {
			record, err := convertUserIntroductions(v)
			if err != nil {
				response.Error(c, ecode.ErrListUserIntroductions)
				return
			}
			userIntroductionss = append(userIntroductionss, record)
		}
	}

	response.Success(c, gin.H{
		"userIntroductionss": userIntroductionss,
	})
}

// ListByLastID get records by last id and limit
// @Summary list of userIntroductionss by last id and limit
// @Description list of userIntroductionss by last id and limit
// @Tags userIntroductions
// @accept json
// @Produce json
// @Param lastID query int true "last id, default is MaxInt32" default(0)
// @Param limit query int false "size in each page" default(10)
// @Param sort query string false "sort by column name of table, and the "-" sign before column name indicates reverse order" default(-id)
// @Success 200 {object} types.ListUserIntroductionssRespond{}
// @Router /api/v1/userIntroductions/list [get]
// @Security BearerAuth
func (h *userIntroductionsHandler) ListByLastID(c *gin.Context) {
	lastID := utils.StrToUint64(c.Query("lastID"))
	if lastID == 0 {
		lastID = math.MaxInt32
	}
	limit := utils.StrToInt(c.Query("limit"))
	if limit == 0 {
		limit = 10
	}
	sort := c.Query("sort")

	ctx := middleware.WrapCtx(c)
	userIntroductionss, err := h.iDao.GetByLastID(ctx, lastID, limit, sort)
	if err != nil {
		logger.Error("GetByLastID error", logger.Err(err), logger.Uint64("latsID", lastID), logger.Int("limit", limit), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertUserIntroductionss(userIntroductionss)
	if err != nil {
		response.Error(c, ecode.ErrListByLastIDUserIntroductions)
		return
	}

	response.Success(c, gin.H{
		"userIntroductionss": data,
	})
}

// List of records by query parameters
// @Summary list of userIntroductionss by query parameters
// @Description list of userIntroductionss by paging and conditions
// @Tags userIntroductions
// @accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.ListUserIntroductionssRespond{}
// @Router /api/v1/userIntroductions/list [post]
// @Security BearerAuth
func (h *userIntroductionsHandler) List(c *gin.Context) {
	form := &types.ListUserIntroductionssRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	userIntroductionss, total, err := h.iDao.GetByColumns(ctx, &form.Params)
	if err != nil {
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertUserIntroductionss(userIntroductionss)
	if err != nil {
		response.Error(c, ecode.ErrListUserIntroductions)
		return
	}

	response.Success(c, gin.H{
		"userIntroductionss": data,
		"total":        total,
	})
}

func getUserIntroductionsIDFromPath(c *gin.Context) (string, uint64, bool) {
	idStr := c.Param("id")
	id, err := utils.StrToUint64E(idStr)
	if err != nil || id == 0 {
		logger.Warn("StrToUint64E error: ", logger.String("idStr", idStr), middleware.GCtxRequestIDField(c))
		return "", 0, true
	}

	return idStr, id, false
}

func convertUserIntroductions(userIntroductions *model.UserIntroductions) (*types.UserIntroductionsObjDetail, error) {
	data := &types.UserIntroductionsObjDetail{}
	err := copier.Copy(data, userIntroductions)
	if err != nil {
		return nil, err
	}
	data.ID = utils.Uint64ToStr(userIntroductions.ID)
	return data, nil
}

func convertUserIntroductionss(fromValues []*model.UserIntroductions) ([]*types.UserIntroductionsObjDetail, error) {
	toValues := []*types.UserIntroductionsObjDetail{}
	for _, v := range fromValues {
		data, err := convertUserIntroductions(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}

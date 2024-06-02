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

var _ EducationsHandler = (*educationsHandler)(nil)

// EducationsHandler defining the handler interface
type EducationsHandler interface {
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

type educationsHandler struct {
	iDao dao.EducationsDao
}

// NewEducationsHandler creating the handler interface
func NewEducationsHandler() EducationsHandler {
	return &educationsHandler{
		iDao: dao.NewEducationsDao(
			model.GetDB(),
			cache.NewEducationsCache(model.GetCacheType()),
		),
	}
}

// Create a record
// @Summary create educations
// @Description submit information to create educations
// @Tags educations
// @accept json
// @Produce json
// @Param data body types.CreateEducationsRequest true "educations information"
// @Success 200 {object} types.CreateEducationsRespond{}
// @Router /api/v1/educations [post]
// @Security BearerAuth
func (h *educationsHandler) Create(c *gin.Context) {
	form := &types.CreateEducationsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	educations := &model.Educations{}
	err = copier.Copy(educations, form)
	if err != nil {
		response.Error(c, ecode.ErrCreateEducations)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, educations)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": educations.ID})
}

// DeleteByID delete a record by id
// @Summary delete educations
// @Description delete educations by id
// @Tags educations
// @accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.DeleteEducationsByIDRespond{}
// @Router /api/v1/educations/{id} [delete]
// @Security BearerAuth
func (h *educationsHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getEducationsIDFromPath(c)
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
// @Summary delete educationss
// @Description delete educationss by batch id
// @Tags educations
// @Param data body types.DeleteEducationssByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.DeleteEducationssByIDsRespond{}
// @Router /api/v1/educations/delete/ids [post]
// @Security BearerAuth
func (h *educationsHandler) DeleteByIDs(c *gin.Context) {
	form := &types.DeleteEducationssByIDsRequest{}
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
// @Summary update educations
// @Description update educations information by id
// @Tags educations
// @accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdateEducationsByIDRequest true "educations information"
// @Success 200 {object} types.UpdateEducationsByIDRespond{}
// @Router /api/v1/educations/{id} [put]
// @Security BearerAuth
func (h *educationsHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getEducationsIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdateEducationsByIDRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	form.ID = id

	educations := &model.Educations{}
	err = copier.Copy(educations, form)
	if err != nil {
		response.Error(c, ecode.ErrUpdateByIDEducations)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, educations)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// GetByID get a record by id
// @Summary get educations detail
// @Description get educations detail by id
// @Tags educations
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetEducationsByIDRespond{}
// @Router /api/v1/educations/{id} [get]
// @Security BearerAuth
func (h *educationsHandler) GetByID(c *gin.Context) {
	idStr, id, isAbort := getEducationsIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	educations, err := h.iDao.GetByID(ctx, id)
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

	data := &types.EducationsObjDetail{}
	err = copier.Copy(data, educations)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDEducations)
		return
	}
	data.ID = idStr

	response.Success(c, gin.H{"educations": data})
}

// GetByCondition get a record by condition
// @Summary get educations by condition
// @Description get educations by condition
// @Tags educations
// @Param data body types.Conditions true "query condition"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetEducationsByConditionRespond{}
// @Router /api/v1/educations/condition [post]
// @Security BearerAuth
func (h *educationsHandler) GetByCondition(c *gin.Context) {
	form := &types.GetEducationsByConditionRequest{}
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
	educations, err := h.iDao.GetByCondition(ctx, &form.Conditions)
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

	data := &types.EducationsObjDetail{}
	err = copier.Copy(data, educations)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDEducations)
		return
	}
	data.ID = utils.Uint64ToStr(educations.ID)

	response.Success(c, gin.H{"educations": data})
}

// ListByIDs list of records by batch id
// @Summary list of educationss by batch id
// @Description list of educationss by batch id
// @Tags educations
// @Param data body types.ListEducationssByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.ListEducationssByIDsRespond{}
// @Router /api/v1/educations/list/ids [post]
// @Security BearerAuth
func (h *educationsHandler) ListByIDs(c *gin.Context) {
	form := &types.ListEducationssByIDsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	educationsMap, err := h.iDao.GetByIDs(ctx, form.IDs)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	educationss := []*types.EducationsObjDetail{}
	for _, id := range form.IDs {
		if v, ok := educationsMap[id]; ok {
			record, err := convertEducations(v)
			if err != nil {
				response.Error(c, ecode.ErrListEducations)
				return
			}
			educationss = append(educationss, record)
		}
	}

	response.Success(c, gin.H{
		"educationss": educationss,
	})
}

// ListByLastID get records by last id and limit
// @Summary list of educationss by last id and limit
// @Description list of educationss by last id and limit
// @Tags educations
// @accept json
// @Produce json
// @Param lastID query int true "last id, default is MaxInt32" default(0)
// @Param limit query int false "size in each page" default(10)
// @Param sort query string false "sort by column name of table, and the "-" sign before column name indicates reverse order" default(-id)
// @Success 200 {object} types.ListEducationssRespond{}
// @Router /api/v1/educations/list [get]
// @Security BearerAuth
func (h *educationsHandler) ListByLastID(c *gin.Context) {
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
	educationss, err := h.iDao.GetByLastID(ctx, lastID, limit, sort)
	if err != nil {
		logger.Error("GetByLastID error", logger.Err(err), logger.Uint64("latsID", lastID), logger.Int("limit", limit), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertEducationss(educationss)
	if err != nil {
		response.Error(c, ecode.ErrListByLastIDEducations)
		return
	}

	response.Success(c, gin.H{
		"educationss": data,
	})
}

// List of records by query parameters
// @Summary list of educationss by query parameters
// @Description list of educationss by paging and conditions
// @Tags educations
// @accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.ListEducationssRespond{}
// @Router /api/v1/educations/list [post]
// @Security BearerAuth
func (h *educationsHandler) List(c *gin.Context) {
	form := &types.ListEducationssRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	educationss, total, err := h.iDao.GetByColumns(ctx, &form.Params)
	if err != nil {
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertEducationss(educationss)
	if err != nil {
		response.Error(c, ecode.ErrListEducations)
		return
	}

	response.Success(c, gin.H{
		"educationss": data,
		"total":        total,
	})
}

func getEducationsIDFromPath(c *gin.Context) (string, uint64, bool) {
	idStr := c.Param("id")
	id, err := utils.StrToUint64E(idStr)
	if err != nil || id == 0 {
		logger.Warn("StrToUint64E error: ", logger.String("idStr", idStr), middleware.GCtxRequestIDField(c))
		return "", 0, true
	}

	return idStr, id, false
}

func convertEducations(educations *model.Educations) (*types.EducationsObjDetail, error) {
	data := &types.EducationsObjDetail{}
	err := copier.Copy(data, educations)
	if err != nil {
		return nil, err
	}
	data.ID = utils.Uint64ToStr(educations.ID)
	return data, nil
}

func convertEducationss(fromValues []*model.Educations) ([]*types.EducationsObjDetail, error) {
	toValues := []*types.EducationsObjDetail{}
	for _, v := range fromValues {
		data, err := convertEducations(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}

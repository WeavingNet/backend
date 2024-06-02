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

var _ WorkexperiencesHandler = (*workexperiencesHandler)(nil)

// WorkexperiencesHandler defining the handler interface
type WorkexperiencesHandler interface {
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

type workexperiencesHandler struct {
	iDao dao.WorkexperiencesDao
}

// NewWorkexperiencesHandler creating the handler interface
func NewWorkexperiencesHandler() WorkexperiencesHandler {
	return &workexperiencesHandler{
		iDao: dao.NewWorkexperiencesDao(
			model.GetDB(),
			cache.NewWorkexperiencesCache(model.GetCacheType()),
		),
	}
}

// Create a record
// @Summary create workexperiences
// @Description submit information to create workexperiences
// @Tags workexperiences
// @accept json
// @Produce json
// @Param data body types.CreateWorkexperiencesRequest true "workexperiences information"
// @Success 200 {object} types.CreateWorkexperiencesRespond{}
// @Router /api/v1/workexperiences [post]
// @Security BearerAuth
func (h *workexperiencesHandler) Create(c *gin.Context) {
	form := &types.CreateWorkexperiencesRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	workexperiences := &model.Workexperiences{}
	err = copier.Copy(workexperiences, form)
	if err != nil {
		response.Error(c, ecode.ErrCreateWorkexperiences)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, workexperiences)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": workexperiences.ID})
}

// DeleteByID delete a record by id
// @Summary delete workexperiences
// @Description delete workexperiences by id
// @Tags workexperiences
// @accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.DeleteWorkexperiencesByIDRespond{}
// @Router /api/v1/workexperiences/{id} [delete]
// @Security BearerAuth
func (h *workexperiencesHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getWorkexperiencesIDFromPath(c)
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
// @Summary delete workexperiencess
// @Description delete workexperiencess by batch id
// @Tags workexperiences
// @Param data body types.DeleteWorkexperiencessByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.DeleteWorkexperiencessByIDsRespond{}
// @Router /api/v1/workexperiences/delete/ids [post]
// @Security BearerAuth
func (h *workexperiencesHandler) DeleteByIDs(c *gin.Context) {
	form := &types.DeleteWorkexperiencessByIDsRequest{}
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
// @Summary update workexperiences
// @Description update workexperiences information by id
// @Tags workexperiences
// @accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdateWorkexperiencesByIDRequest true "workexperiences information"
// @Success 200 {object} types.UpdateWorkexperiencesByIDRespond{}
// @Router /api/v1/workexperiences/{id} [put]
// @Security BearerAuth
func (h *workexperiencesHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getWorkexperiencesIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdateWorkexperiencesByIDRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	form.ID = id

	workexperiences := &model.Workexperiences{}
	err = copier.Copy(workexperiences, form)
	if err != nil {
		response.Error(c, ecode.ErrUpdateByIDWorkexperiences)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, workexperiences)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// GetByID get a record by id
// @Summary get workexperiences detail
// @Description get workexperiences detail by id
// @Tags workexperiences
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetWorkexperiencesByIDRespond{}
// @Router /api/v1/workexperiences/{id} [get]
// @Security BearerAuth
func (h *workexperiencesHandler) GetByID(c *gin.Context) {
	idStr, id, isAbort := getWorkexperiencesIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	workexperiences, err := h.iDao.GetByID(ctx, id)
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

	data := &types.WorkexperiencesObjDetail{}
	err = copier.Copy(data, workexperiences)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDWorkexperiences)
		return
	}
	data.ID = idStr

	response.Success(c, gin.H{"workexperiences": data})
}

// GetByCondition get a record by condition
// @Summary get workexperiences by condition
// @Description get workexperiences by condition
// @Tags workexperiences
// @Param data body types.Conditions true "query condition"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetWorkexperiencesByConditionRespond{}
// @Router /api/v1/workexperiences/condition [post]
// @Security BearerAuth
func (h *workexperiencesHandler) GetByCondition(c *gin.Context) {
	form := &types.GetWorkexperiencesByConditionRequest{}
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
	workexperiences, err := h.iDao.GetByCondition(ctx, &form.Conditions)
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

	data := &types.WorkexperiencesObjDetail{}
	err = copier.Copy(data, workexperiences)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDWorkexperiences)
		return
	}
	data.ID = utils.Uint64ToStr(workexperiences.ID)

	response.Success(c, gin.H{"workexperiences": data})
}

// ListByIDs list of records by batch id
// @Summary list of workexperiencess by batch id
// @Description list of workexperiencess by batch id
// @Tags workexperiences
// @Param data body types.ListWorkexperiencessByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.ListWorkexperiencessByIDsRespond{}
// @Router /api/v1/workexperiences/list/ids [post]
// @Security BearerAuth
func (h *workexperiencesHandler) ListByIDs(c *gin.Context) {
	form := &types.ListWorkexperiencessByIDsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	workexperiencesMap, err := h.iDao.GetByIDs(ctx, form.IDs)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	workexperiencess := []*types.WorkexperiencesObjDetail{}
	for _, id := range form.IDs {
		if v, ok := workexperiencesMap[id]; ok {
			record, err := convertWorkexperiences(v)
			if err != nil {
				response.Error(c, ecode.ErrListWorkexperiences)
				return
			}
			workexperiencess = append(workexperiencess, record)
		}
	}

	response.Success(c, gin.H{
		"workexperiencess": workexperiencess,
	})
}

// ListByLastID get records by last id and limit
// @Summary list of workexperiencess by last id and limit
// @Description list of workexperiencess by last id and limit
// @Tags workexperiences
// @accept json
// @Produce json
// @Param lastID query int true "last id, default is MaxInt32" default(0)
// @Param limit query int false "size in each page" default(10)
// @Param sort query string false "sort by column name of table, and the "-" sign before column name indicates reverse order" default(-id)
// @Success 200 {object} types.ListWorkexperiencessRespond{}
// @Router /api/v1/workexperiences/list [get]
// @Security BearerAuth
func (h *workexperiencesHandler) ListByLastID(c *gin.Context) {
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
	workexperiencess, err := h.iDao.GetByLastID(ctx, lastID, limit, sort)
	if err != nil {
		logger.Error("GetByLastID error", logger.Err(err), logger.Uint64("latsID", lastID), logger.Int("limit", limit), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertWorkexperiencess(workexperiencess)
	if err != nil {
		response.Error(c, ecode.ErrListByLastIDWorkexperiences)
		return
	}

	response.Success(c, gin.H{
		"workexperiencess": data,
	})
}

// List of records by query parameters
// @Summary list of workexperiencess by query parameters
// @Description list of workexperiencess by paging and conditions
// @Tags workexperiences
// @accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.ListWorkexperiencessRespond{}
// @Router /api/v1/workexperiences/list [post]
// @Security BearerAuth
func (h *workexperiencesHandler) List(c *gin.Context) {
	form := &types.ListWorkexperiencessRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	workexperiencess, total, err := h.iDao.GetByColumns(ctx, &form.Params)
	if err != nil {
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertWorkexperiencess(workexperiencess)
	if err != nil {
		response.Error(c, ecode.ErrListWorkexperiences)
		return
	}

	response.Success(c, gin.H{
		"workexperiencess": data,
		"total":        total,
	})
}

func getWorkexperiencesIDFromPath(c *gin.Context) (string, uint64, bool) {
	idStr := c.Param("id")
	id, err := utils.StrToUint64E(idStr)
	if err != nil || id == 0 {
		logger.Warn("StrToUint64E error: ", logger.String("idStr", idStr), middleware.GCtxRequestIDField(c))
		return "", 0, true
	}

	return idStr, id, false
}

func convertWorkexperiences(workexperiences *model.Workexperiences) (*types.WorkexperiencesObjDetail, error) {
	data := &types.WorkexperiencesObjDetail{}
	err := copier.Copy(data, workexperiences)
	if err != nil {
		return nil, err
	}
	data.ID = utils.Uint64ToStr(workexperiences.ID)
	return data, nil
}

func convertWorkexperiencess(fromValues []*model.Workexperiences) ([]*types.WorkexperiencesObjDetail, error) {
	toValues := []*types.WorkexperiencesObjDetail{}
	for _, v := range fromValues {
		data, err := convertWorkexperiences(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}

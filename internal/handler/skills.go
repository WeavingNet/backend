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

var _ SkillsHandler = (*skillsHandler)(nil)

// SkillsHandler defining the handler interface
type SkillsHandler interface {
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

type skillsHandler struct {
	iDao dao.SkillsDao
}

// NewSkillsHandler creating the handler interface
func NewSkillsHandler() SkillsHandler {
	return &skillsHandler{
		iDao: dao.NewSkillsDao(
			model.GetDB(),
			cache.NewSkillsCache(model.GetCacheType()),
		),
	}
}

// Create a record
// @Summary create skills
// @Description submit information to create skills
// @Tags skills
// @accept json
// @Produce json
// @Param data body types.CreateSkillsRequest true "skills information"
// @Success 200 {object} types.CreateSkillsRespond{}
// @Router /api/v1/skills [post]
// @Security BearerAuth
func (h *skillsHandler) Create(c *gin.Context) {
	form := &types.CreateSkillsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	skills := &model.Skills{}
	err = copier.Copy(skills, form)
	if err != nil {
		response.Error(c, ecode.ErrCreateSkills)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, skills)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": skills.ID})
}

// DeleteByID delete a record by id
// @Summary delete skills
// @Description delete skills by id
// @Tags skills
// @accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.DeleteSkillsByIDRespond{}
// @Router /api/v1/skills/{id} [delete]
// @Security BearerAuth
func (h *skillsHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getSkillsIDFromPath(c)
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
// @Summary delete skillss
// @Description delete skillss by batch id
// @Tags skills
// @Param data body types.DeleteSkillssByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.DeleteSkillssByIDsRespond{}
// @Router /api/v1/skills/delete/ids [post]
// @Security BearerAuth
func (h *skillsHandler) DeleteByIDs(c *gin.Context) {
	form := &types.DeleteSkillssByIDsRequest{}
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
// @Summary update skills
// @Description update skills information by id
// @Tags skills
// @accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdateSkillsByIDRequest true "skills information"
// @Success 200 {object} types.UpdateSkillsByIDRespond{}
// @Router /api/v1/skills/{id} [put]
// @Security BearerAuth
func (h *skillsHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getSkillsIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdateSkillsByIDRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	form.ID = id

	skills := &model.Skills{}
	err = copier.Copy(skills, form)
	if err != nil {
		response.Error(c, ecode.ErrUpdateByIDSkills)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, skills)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// GetByID get a record by id
// @Summary get skills detail
// @Description get skills detail by id
// @Tags skills
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetSkillsByIDRespond{}
// @Router /api/v1/skills/{id} [get]
// @Security BearerAuth
func (h *skillsHandler) GetByID(c *gin.Context) {
	idStr, id, isAbort := getSkillsIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	skills, err := h.iDao.GetByID(ctx, id)
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

	data := &types.SkillsObjDetail{}
	err = copier.Copy(data, skills)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDSkills)
		return
	}
	data.ID = idStr

	response.Success(c, gin.H{"skills": data})
}

// GetByCondition get a record by condition
// @Summary get skills by condition
// @Description get skills by condition
// @Tags skills
// @Param data body types.Conditions true "query condition"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetSkillsByConditionRespond{}
// @Router /api/v1/skills/condition [post]
// @Security BearerAuth
func (h *skillsHandler) GetByCondition(c *gin.Context) {
	form := &types.GetSkillsByConditionRequest{}
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
	skills, err := h.iDao.GetByCondition(ctx, &form.Conditions)
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

	data := &types.SkillsObjDetail{}
	err = copier.Copy(data, skills)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDSkills)
		return
	}
	data.ID = utils.Uint64ToStr(skills.ID)

	response.Success(c, gin.H{"skills": data})
}

// ListByIDs list of records by batch id
// @Summary list of skillss by batch id
// @Description list of skillss by batch id
// @Tags skills
// @Param data body types.ListSkillssByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.ListSkillssByIDsRespond{}
// @Router /api/v1/skills/list/ids [post]
// @Security BearerAuth
func (h *skillsHandler) ListByIDs(c *gin.Context) {
	form := &types.ListSkillssByIDsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	skillsMap, err := h.iDao.GetByIDs(ctx, form.IDs)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	skillss := []*types.SkillsObjDetail{}
	for _, id := range form.IDs {
		if v, ok := skillsMap[id]; ok {
			record, err := convertSkills(v)
			if err != nil {
				response.Error(c, ecode.ErrListSkills)
				return
			}
			skillss = append(skillss, record)
		}
	}

	response.Success(c, gin.H{
		"skillss": skillss,
	})
}

// ListByLastID get records by last id and limit
// @Summary list of skillss by last id and limit
// @Description list of skillss by last id and limit
// @Tags skills
// @accept json
// @Produce json
// @Param lastID query int true "last id, default is MaxInt32" default(0)
// @Param limit query int false "size in each page" default(10)
// @Param sort query string false "sort by column name of table, and the "-" sign before column name indicates reverse order" default(-id)
// @Success 200 {object} types.ListSkillssRespond{}
// @Router /api/v1/skills/list [get]
// @Security BearerAuth
func (h *skillsHandler) ListByLastID(c *gin.Context) {
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
	skillss, err := h.iDao.GetByLastID(ctx, lastID, limit, sort)
	if err != nil {
		logger.Error("GetByLastID error", logger.Err(err), logger.Uint64("latsID", lastID), logger.Int("limit", limit), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertSkillss(skillss)
	if err != nil {
		response.Error(c, ecode.ErrListByLastIDSkills)
		return
	}

	response.Success(c, gin.H{
		"skillss": data,
	})
}

// List of records by query parameters
// @Summary list of skillss by query parameters
// @Description list of skillss by paging and conditions
// @Tags skills
// @accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.ListSkillssRespond{}
// @Router /api/v1/skills/list [post]
// @Security BearerAuth
func (h *skillsHandler) List(c *gin.Context) {
	form := &types.ListSkillssRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	skillss, total, err := h.iDao.GetByColumns(ctx, &form.Params)
	if err != nil {
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertSkillss(skillss)
	if err != nil {
		response.Error(c, ecode.ErrListSkills)
		return
	}

	response.Success(c, gin.H{
		"skillss": data,
		"total":        total,
	})
}

func getSkillsIDFromPath(c *gin.Context) (string, uint64, bool) {
	idStr := c.Param("id")
	id, err := utils.StrToUint64E(idStr)
	if err != nil || id == 0 {
		logger.Warn("StrToUint64E error: ", logger.String("idStr", idStr), middleware.GCtxRequestIDField(c))
		return "", 0, true
	}

	return idStr, id, false
}

func convertSkills(skills *model.Skills) (*types.SkillsObjDetail, error) {
	data := &types.SkillsObjDetail{}
	err := copier.Copy(data, skills)
	if err != nil {
		return nil, err
	}
	data.ID = utils.Uint64ToStr(skills.ID)
	return data, nil
}

func convertSkillss(fromValues []*model.Skills) ([]*types.SkillsObjDetail, error) {
	toValues := []*types.SkillsObjDetail{}
	for _, v := range fromValues {
		data, err := convertSkills(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}

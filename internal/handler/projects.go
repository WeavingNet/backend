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

var _ ProjectsHandler = (*projectsHandler)(nil)

// ProjectsHandler defining the handler interface
type ProjectsHandler interface {
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

type projectsHandler struct {
	iDao dao.ProjectsDao
}

// NewProjectsHandler creating the handler interface
func NewProjectsHandler() ProjectsHandler {
	return &projectsHandler{
		iDao: dao.NewProjectsDao(
			model.GetDB(),
			cache.NewProjectsCache(model.GetCacheType()),
		),
	}
}

// Create a record
// @Summary create projects
// @Description submit information to create projects
// @Tags projects
// @accept json
// @Produce json
// @Param data body types.CreateProjectsRequest true "projects information"
// @Success 200 {object} types.CreateProjectsRespond{}
// @Router /api/v1/projects [post]
// @Security BearerAuth
func (h *projectsHandler) Create(c *gin.Context) {
	form := &types.CreateProjectsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	projects := &model.Projects{}
	err = copier.Copy(projects, form)
	if err != nil {
		response.Error(c, ecode.ErrCreateProjects)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, projects)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": projects.ID})
}

// DeleteByID delete a record by id
// @Summary delete projects
// @Description delete projects by id
// @Tags projects
// @accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.DeleteProjectsByIDRespond{}
// @Router /api/v1/projects/{id} [delete]
// @Security BearerAuth
func (h *projectsHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getProjectsIDFromPath(c)
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
// @Summary delete projectss
// @Description delete projectss by batch id
// @Tags projects
// @Param data body types.DeleteProjectssByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.DeleteProjectssByIDsRespond{}
// @Router /api/v1/projects/delete/ids [post]
// @Security BearerAuth
func (h *projectsHandler) DeleteByIDs(c *gin.Context) {
	form := &types.DeleteProjectssByIDsRequest{}
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
// @Summary update projects
// @Description update projects information by id
// @Tags projects
// @accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdateProjectsByIDRequest true "projects information"
// @Success 200 {object} types.UpdateProjectsByIDRespond{}
// @Router /api/v1/projects/{id} [put]
// @Security BearerAuth
func (h *projectsHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getProjectsIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdateProjectsByIDRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	form.ID = id

	projects := &model.Projects{}
	err = copier.Copy(projects, form)
	if err != nil {
		response.Error(c, ecode.ErrUpdateByIDProjects)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, projects)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// GetByID get a record by id
// @Summary get projects detail
// @Description get projects detail by id
// @Tags projects
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetProjectsByIDRespond{}
// @Router /api/v1/projects/{id} [get]
// @Security BearerAuth
func (h *projectsHandler) GetByID(c *gin.Context) {
	idStr, id, isAbort := getProjectsIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	projects, err := h.iDao.GetByID(ctx, id)
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

	data := &types.ProjectsObjDetail{}
	err = copier.Copy(data, projects)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDProjects)
		return
	}
	data.ID = idStr

	response.Success(c, gin.H{"projects": data})
}

// GetByCondition get a record by condition
// @Summary get projects by condition
// @Description get projects by condition
// @Tags projects
// @Param data body types.Conditions true "query condition"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetProjectsByConditionRespond{}
// @Router /api/v1/projects/condition [post]
// @Security BearerAuth
func (h *projectsHandler) GetByCondition(c *gin.Context) {
	form := &types.GetProjectsByConditionRequest{}
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
	projects, err := h.iDao.GetByCondition(ctx, &form.Conditions)
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

	data := &types.ProjectsObjDetail{}
	err = copier.Copy(data, projects)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDProjects)
		return
	}
	data.ID = utils.Uint64ToStr(projects.ID)

	response.Success(c, gin.H{"projects": data})
}

// ListByIDs list of records by batch id
// @Summary list of projectss by batch id
// @Description list of projectss by batch id
// @Tags projects
// @Param data body types.ListProjectssByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.ListProjectssByIDsRespond{}
// @Router /api/v1/projects/list/ids [post]
// @Security BearerAuth
func (h *projectsHandler) ListByIDs(c *gin.Context) {
	form := &types.ListProjectssByIDsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	projectsMap, err := h.iDao.GetByIDs(ctx, form.IDs)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	projectss := []*types.ProjectsObjDetail{}
	for _, id := range form.IDs {
		if v, ok := projectsMap[id]; ok {
			record, err := convertProjects(v)
			if err != nil {
				response.Error(c, ecode.ErrListProjects)
				return
			}
			projectss = append(projectss, record)
		}
	}

	response.Success(c, gin.H{
		"projectss": projectss,
	})
}

// ListByLastID get records by last id and limit
// @Summary list of projectss by last id and limit
// @Description list of projectss by last id and limit
// @Tags projects
// @accept json
// @Produce json
// @Param lastID query int true "last id, default is MaxInt32" default(0)
// @Param limit query int false "size in each page" default(10)
// @Param sort query string false "sort by column name of table, and the "-" sign before column name indicates reverse order" default(-id)
// @Success 200 {object} types.ListProjectssRespond{}
// @Router /api/v1/projects/list [get]
// @Security BearerAuth
func (h *projectsHandler) ListByLastID(c *gin.Context) {
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
	projectss, err := h.iDao.GetByLastID(ctx, lastID, limit, sort)
	if err != nil {
		logger.Error("GetByLastID error", logger.Err(err), logger.Uint64("latsID", lastID), logger.Int("limit", limit), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertProjectss(projectss)
	if err != nil {
		response.Error(c, ecode.ErrListByLastIDProjects)
		return
	}

	response.Success(c, gin.H{
		"projectss": data,
	})
}

// List of records by query parameters
// @Summary list of projectss by query parameters
// @Description list of projectss by paging and conditions
// @Tags projects
// @accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.ListProjectssRespond{}
// @Router /api/v1/projects/list [post]
// @Security BearerAuth
func (h *projectsHandler) List(c *gin.Context) {
	form := &types.ListProjectssRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	projectss, total, err := h.iDao.GetByColumns(ctx, &form.Params)
	if err != nil {
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertProjectss(projectss)
	if err != nil {
		response.Error(c, ecode.ErrListProjects)
		return
	}

	response.Success(c, gin.H{
		"projectss": data,
		"total":        total,
	})
}

func getProjectsIDFromPath(c *gin.Context) (string, uint64, bool) {
	idStr := c.Param("id")
	id, err := utils.StrToUint64E(idStr)
	if err != nil || id == 0 {
		logger.Warn("StrToUint64E error: ", logger.String("idStr", idStr), middleware.GCtxRequestIDField(c))
		return "", 0, true
	}

	return idStr, id, false
}

func convertProjects(projects *model.Projects) (*types.ProjectsObjDetail, error) {
	data := &types.ProjectsObjDetail{}
	err := copier.Copy(data, projects)
	if err != nil {
		return nil, err
	}
	data.ID = utils.Uint64ToStr(projects.ID)
	return data, nil
}

func convertProjectss(fromValues []*model.Projects) ([]*types.ProjectsObjDetail, error) {
	toValues := []*types.ProjectsObjDetail{}
	for _, v := range fromValues {
		data, err := convertProjects(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}

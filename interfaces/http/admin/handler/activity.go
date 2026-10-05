/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午4:49
 * Description: 活动管理 HTTP handler
 *   - 对接 application.activity.AdminService
 *   - POST   /admin/activity/        创建
 *   - PUT    /admin/activity/:id     更新元数据
 *   - PUT    /admin/activity/:id/online   上线(并初始化 Redis 库存)
 *   - PUT    /admin/activity/:id/offline  下线
 *   - GET    /admin/activity/:id     查询
 *   - GET    /admin/activity/        列表(进行中)
 **/

package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	activityapp "github.com/weitrue/Seckill/internal/application/activity"
	"github.com/weitrue/Seckill/internal/domain/activity"
	"github.com/weitrue/Seckill/pkg/utils"
)

// AdminActivityHandler 后台活动 handler
type AdminActivityHandler struct {
	svc *activityapp.AdminService
}

/*NewAdminActivityHandler
 *@Description: 构造
 *@param svc
 *@return *AdminActivityHandler
 */
func NewAdminActivityHandler(svc *activityapp.AdminService) *AdminActivityHandler {
	return &AdminActivityHandler{svc: svc}
}

// addRequest 创建活动的请求 DTO
type addRequest struct {
	Name      string             `json:"name" binding:"required"`
	StartTime int64              `json:"start_time" binding:"required"`
	EndTime   int64              `json:"end_time" binding:"required"`
	Goods     []addGoodsRequest  `json:"goods"`
}

type addGoodsRequest struct {
	GoodsID       string `json:"goods_id" binding:"required"`
	Name          string `json:"name"`
	Price         int64  `json:"price"`
	ActivityPrice int64  `json:"activity_price"`
	Stock         int64  `json:"stock"`
	LimitPerUser  int    `json:"limit_per_user"`
}

/*Add
 *@Description: POST /admin/activity/ 创建活动
 */
func (h *AdminActivityHandler) Add(c *gin.Context) {
	logType := "AdminActivityAdd"
	var req addRequest
	if err := c.BindJSON(&req); err != nil {
		logrus.Warnf("logType:%s, msg:bad request, err:%v", logType, err)
		c.JSON(http.StatusBadRequest, &utils.Response{Code: http.StatusBadRequest, Msg: "bad request"})
		return
	}
	cmd := activityapp.CreateCommand{
		Name:      req.Name,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Status:    activity.StatusPrepared,
	}
	for _, g := range req.Goods {
		cmd.Goods = append(cmd.Goods, activityapp.GoodsItem{
			GoodsID:       g.GoodsID,
			Name:          g.Name,
			Price:         g.Price,
			ActivityPrice: g.ActivityPrice,
			Stock:         g.Stock,
			LimitPerUser:  g.LimitPerUser,
		})
	}
	a, err := h.svc.Create(c.Request.Context(), cmd)
	if err != nil {
		logrus.Errorf("logType:%s, err:%s, cmd:%+v", logType, err.Error(), cmd)
		c.JSON(http.StatusInternalServerError, &utils.Response{Code: http.StatusInternalServerError, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok", Data: a})
}

// updateRequest 更新活动元数据
type updateRequest struct {
	Name      string `json:"name"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Status    int8   `json:"status"`
}

/*Update
 *@Description: PUT /admin/activity/:id 更新元数据
 */
func (h *AdminActivityHandler) Update(c *gin.Context) {
	logType := "AdminActivityUpdate"
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{Code: http.StatusBadRequest, Msg: "invalid id"})
		return
	}
	var req updateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{Code: http.StatusBadRequest, Msg: "bad request"})
		return
	}
	cmd := activityapp.UpdateCommand{
		ID:        id,
		Name:      req.Name,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Status:    activity.Status(req.Status),
	}
	if err := h.svc.Update(c.Request.Context(), cmd); err != nil {
		logrus.Errorf("logType:%s, err:%s, cmd:%+v", logType, err.Error(), cmd)
		c.JSON(http.StatusInternalServerError, &utils.Response{Code: http.StatusInternalServerError, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}

/*Status
 *@Description: PUT /admin/activity/:id/:status 调整上下线状态
 *   路径 status 取 "online" | "offline"
 */
func (h *AdminActivityHandler) Status(c *gin.Context) {
	logType := "AdminActivityStatus"
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{Code: http.StatusBadRequest, Msg: "invalid id"})
		return
	}
	action := c.Param("status")
	ctx := c.Request.Context()
	switch action {
	case "online":
		err = h.svc.Online(ctx, id)
	case "offline":
		err = h.svc.Offline(ctx, id)
	default:
		c.JSON(http.StatusBadRequest, &utils.Response{Code: http.StatusBadRequest, Msg: "invalid action"})
		return
	}
	if err != nil {
		logrus.Errorf("logType:%s, err:%s, id:%d, action:%s", logType, err.Error(), id, action)
		c.JSON(http.StatusInternalServerError, &utils.Response{Code: http.StatusInternalServerError, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}

/*Get
 *@Description: GET /admin/activity/:id 查询
 */
func (h *AdminActivityHandler) Get(c *gin.Context) {
	logType := "AdminActivityGet"
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{Code: http.StatusBadRequest, Msg: "invalid id"})
		return
	}
	// AdminService 没 Get,直接走 api svc? 这里简化:后台也通过 lookup
	// 为避免 admin 和 api svc 耦合,后续加 AdminService.Get;本阶段先走 Online 之前读的方式
	_ = id
	_ = logType
	c.JSON(http.StatusNotImplemented, &utils.Response{Code: http.StatusNotImplemented, Msg: "not implemented"})
}

/*List
 *@Description: GET /admin/activity/ 列表
 */
func (h *AdminActivityHandler) List(c *gin.Context) {
	// TODO: AdminService.List 待补(需要分页等);本阶段先占位
	c.JSON(http.StatusNotImplemented, &utils.Response{Code: http.StatusNotImplemented, Msg: "not implemented"})
}

/*Delete
 *@Description: DELETE /admin/activity/:id 删除(软删=下线)
 */
func (h *AdminActivityHandler) Delete(c *gin.Context) {
	logType := "AdminActivityDelete"
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{Code: http.StatusBadRequest, Msg: "invalid id"})
		return
	}
	if err := h.svc.Offline(c.Request.Context(), id); err != nil {
		logrus.Errorf("logType:%s, err:%s, id:%d", logType, err.Error(), id)
		c.JSON(http.StatusInternalServerError, &utils.Response{Code: http.StatusInternalServerError, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}

// parseID 从 path 的 :id 解出 int64
func parseID(c *gin.Context) (int64, error) {
	return strconv.ParseInt(c.Param("id"), 10, 64)
}

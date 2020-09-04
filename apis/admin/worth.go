package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"net/http"
)

// @Summary 获取当前净值
// @Description 获取当前净值2
// @Tags 企业网站接口
// @Param data body models.NetWorth true "body"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/netWorth/ [get]
// @Security Bearer
func GetNetWorth(c *gin.Context) {
	var res models.NetWorth
	res.ID, _ = tools.StringToInt64(c.Param("id"))
	var resp app.Response
	data, _ := res.Get()
	resp.Data = data
	c.JSON(http.StatusOK, resp.ReturnOK())

}

// @Summary 获取净值列表
// @Description 获取当前净值列表
// @Tags 企业网站接口
// @Param data body models.NetWorth true "body"
// @Success 200 {object} app.Response "{"code": 200, "data": "list":[...]}"
// @Router /api/v1/netWorth/list [get]
// @Security Bearer
func GetNetWorthList(c *gin.Context) {
	var err error

	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}
	var res models.NetWorth
	list, count, _ := res.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)
	app.PageOK(c, list, count, pageIndex, pageSize, "")
}

func UpdateNetWorth(c *gin.Context) {
	var data models.NetWorth
	var resp app.Response
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	if data.ID == 0 {
		err = errors.New("")
		tools.HasError(err, "缺省参数", -1)
		c.JSON(http.StatusOK, resp.ReturnOK())
		return
	}
	data, err = data.Update()
	tools.HasError(err, "数据不存在", -1)
	c.JSON(http.StatusOK, resp.ReturnOK())
}

func InsertNetWorth(c *gin.Context) {
	var data models.NetWorth
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "非法数据格式", 500)
	data.CreateBy = tools.GetUserIdStr(c)
	id, err := data.Insert()
	tools.HasError(err, "添加失败", 500)
	app.OK(c, id, "添加成功")
}

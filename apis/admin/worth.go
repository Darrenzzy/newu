package admin

import (
	"context"
	"encoding/json"
	"errors"
	"go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"net/http"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	res.ID, _ = tools.StringToInt(c.Param("id"))
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
	saveData := models.RkData{}
	bs, _ := redis.String(global.Rdb.Do(context.TODO(), "get", "worth_list"))
	if bs != "" {
		_ = json.Unmarshal([]byte(bs), &saveData)
		if saveData.Count > 0 {
			log.Info("走缓存~")
			app.PageOK(c, saveData.Data, saveData.Count, 1, 10, "")
			return
		}
	}

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

	saveData.Data = list
	saveData.Count = count
	str, _ := json.Marshal(saveData)
	global.Rdb.Do(context.TODO(), "set", "worth_list", string(str))

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
	global.Rdb.Do(context.TODO(), "del", "worth_list")

	c.JSON(http.StatusOK, resp.ReturnOK())
}

func DeleteWorth(c *gin.Context) {
	var data models.NetWorth
	// data.UpdateBy = tools.GetUserIdStr(c)
	IDS := tools.IdsStrToIdsIntGroup("id", c)
	result, err := data.BatchDelete(IDS)
	tools.HasError(err, "删除失败", 500)
	global.Rdb.Do(context.TODO(), "del", "worth_list")
	app.OK(c, result, "删除成功")
}

func InsertNetWorth(c *gin.Context) {
	var data models.NetWorth
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "非法数据格式", 500)

	// data.CreateBy = tools.GetUserIdStr(c)
	// data.CreateBy = data.CreateBy
	id, err := data.Insert()
	tools.HasError(err, "添加失败", 500)
	global.Rdb.Do(context.TODO(), "del", "worth_list")
	app.OK(c, id, "添加成功")
}

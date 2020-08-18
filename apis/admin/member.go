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

func GetMemberList(c *gin.Context) {
	var res models.Member
	var err error

	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}
	if name := c.Request.FormValue("username"); name != "" {
		res.Username = name
	}
	if name := c.Request.FormValue("mobile"); name != "" {
		res.Mobile = name
	}
	result, count, err := res.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)
	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

func GetMember(c *gin.Context) {
	var res models.Member
	res.ID, _ = tools.StringToInt64(c.Param("id"))
	res, err := res.Get()
	tools.HasError(err, " 无该用户", -1)
	var resp app.Response
	resp.Data = res
	c.JSON(http.StatusOK, resp.ReturnOK())
}

func UpdateMember(c *gin.Context) {
	var data models.Member
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

func DeleteMember(c *gin.Context) {
	var data models.Member
	// data.UpdateBy = tools.GetUserIdStr(c)
	IDS := tools.IdsStrToIdsIntGroup("id", c)
	result, err := data.BatchDelete(IDS)
	tools.HasError(err, "删除失败", 500)
	app.OK(c, result, "删除成功")
}

func InsertMember(c *gin.Context) {
	var sysuser models.Member
	err := c.BindWith(&sysuser, binding.JSON)
	tools.HasError(err, "非法数据格式", 500)

	// sysuser.CreateBy = tools.GetUserIdStr(c)
	// id, err := sysuser.Insert()
	// tools.HasError(err, "添加失败", 500)
	// app.OK(c, id, "添加成功")
}

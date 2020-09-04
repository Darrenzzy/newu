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

func GetMemberTransactionList(c *gin.Context) {
	var res models.MemberTransaction
	var err error

	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}
	result, count, err := res.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)
	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

func GetMemberTransaction(c *gin.Context) {
	var res models.MemberTransaction
	res.ID, _ = tools.StringToInt64(c.Param("id"))
	res, err := res.Get()
	tools.HasError(err, " 无该记录", -1)
	var resp app.Response
	resp.Data = res
	c.JSON(http.StatusOK, resp.ReturnOK())
}

func UpdateMemberTransaction(c *gin.Context) {
	var data models.MemberTransaction
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

func DeleteMemberTransaction(c *gin.Context) {
	var data models.MemberTransaction
	// data.UpdateBy = tools.GetUserIdStr(c)
	IDS := tools.IdsStrToIdsIntGroup("id", c)
	result, err := data.BatchDelete(IDS)
	tools.HasError(err, "删除失败", 500)
	app.OK(c, result, "删除成功")
}

func AddMemberTransaction(c *gin.Context) {
	var member *models.MemberTransaction
	err := c.BindWith(&member, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)

	if member.Rate == 0 || member.Type == 0 {
		err = errors.New("")
		tools.HasError(err, "缺省参数", 500)
	}

	if member.Amount == 0 && member.OtherAmount == 0 {
		err = errors.New("")
		tools.HasError(err, "金额缺省参数", 500)
	}

	id, err := member.Insert()
	tools.HasError(err, "添加失败", 500)
	app.OK(c, id, "添加成功")
}

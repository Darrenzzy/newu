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

func GetContacts(c *gin.Context) {
	var res models.Contacts
	res.ID, _ = tools.StringToInt64(c.Param("id"))
	var resp app.Response
	data, _ := res.Get()
	resp.Data = data
	c.JSON(http.StatusOK, resp.ReturnOK())

}

func GetContactsList(c *gin.Context) {
	var err error

	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}
	var res models.Contacts
	list, count, _ := res.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)
	app.PageOK(c, list, count, pageIndex, pageSize, "")
}

func UpdateContacts(c *gin.Context) {
	var data models.Contacts
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

func DeleteContacts(c *gin.Context) {
	var data models.Contacts
	// data.UpdateBy = tools.GetUserIdStr(c)
	IDS := tools.IdsStrToIdsIntGroup("id", c)
	result, err := data.BatchDelete(IDS)
	tools.HasError(err, "删除失败", 500)
	app.OK(c, result, "删除成功")
}

// @Summary 联系我们
// @Description 联系我们表单提交
// @Tags 企业网站接口
// @Param data body models.Contacts true "body"
// @Success 200 {object} app.Response "{"code": 200, "data":{}}"
// @Router /api/v1/contacts [post]
// @Security Bearer
func InsertContacts(c *gin.Context) {
	var data models.Contacts
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "非法数据格式", 500)
	id, err := data.Insert()
	tools.HasError(err, "添加失败", 500)
	app.OK(c, id, "添加成功")
}

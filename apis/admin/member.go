package admin

import (
	"github.com/gin-gonic/gin"
	"go-admin/models"

	// "go-admin/models"
	"go-admin/tools/app"
	"net/http"
)

func GetMemberList(c *gin.Context) {

	var res models.Member

	list, _ := res.GetList()

	var resp app.Response

	resp.Data = list

	c.JSON(http.StatusOK, resp.ReturnOK())
}

func GetMember(c *gin.Context) {

	var res models.Member

	list, _ := res.GetList()

	var resp app.Response

	resp.Data = list

	c.JSON(http.StatusOK, resp.ReturnOK())
}

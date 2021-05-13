package admin

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/log"
	"go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"net/http"
	"strconv"
)

func PushResumeData(c *gin.Context) {
	saveData := models.RkResumeData{}
	saveData.Name = c.Request.FormValue("name")
	saveData.Title = c.Request.FormValue("title")
	saveData.Subtitle = c.Request.FormValue("subtitle")
	saveData.Content = c.Request.FormValue("content")
	saveData.Show, _ = strconv.Atoi(c.Request.FormValue("show"))
	saveData.Errno, _ = strconv.Atoi(c.Request.FormValue("errno"))
	if saveData.Name == "" {
		err := c.ShouldBindWith(&saveData, binding.JSON)
		tools.HasError(err, "数据解析失败", -1)
	}

	if saveData.Content == "" || saveData.Title == "" || saveData.Name == "" {
		err := errors.New("")
		tools.HasError(err, "缺省参数", -1)
		c.JSON(http.StatusOK, saveData)
		return
	}
	// 用name 区分不同人的数据
	key := "resume_data_" + saveData.Name
	bs, _ := json.Marshal(saveData)
	global.Rdb.Do(context.TODO(), "set", key, string(bs))
	c.JSON(http.StatusOK, saveData)
}

func GetResumeData(c *gin.Context) {
	var err error
	saveData := models.RkResumeData{}
	name := c.Request.FormValue("name")
	saveData.Name = name

	if saveData.Name == "" {
		err = errors.New("")
		tools.HasError(err, "缺省参数", -1)
		c.JSON(http.StatusOK, saveData)
		return
	}
	// 用name 区分不同人的数据
	key := "resume_data_" + saveData.Name
	bs, _ := redis.String(global.Rdb.Do(context.TODO(), "get", key))
	if bs != "" {
		err = json.Unmarshal([]byte(bs), &saveData)
		if err == nil {
			log.Info("走缓存~")
		}
	}
	c.JSON(http.StatusOK, saveData)
}

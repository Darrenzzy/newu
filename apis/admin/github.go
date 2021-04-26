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
)

func PushResumeData(c *gin.Context) {
	saveData := models.RkResumeData{}
	err := c.BindWith(&saveData, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	if saveData.Content == "" || saveData.Title == "" {
		err = errors.New("")
		tools.HasError(err, "缺省参数", -1)
		c.JSON(http.StatusOK, saveData)
		return
	}
	bs, _ := json.Marshal(saveData)
	global.Rdb.Do(context.TODO(), "set", "resume_data", string(bs))
	c.JSON(http.StatusOK, saveData)
}

func GetResumeData(c *gin.Context) {
	var err error
	saveData := models.RkResumeData{}
	bs, _ := redis.String(global.Rdb.Do(context.TODO(), "get", "resume_data"))
	if bs != "" {
		err = json.Unmarshal([]byte(bs), &saveData)
		if err == nil {
			log.Info("走缓存~")
		}
	}
	c.JSON(http.StatusOK, saveData)
}

package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/gogf/gf/os/glog"
	"github.com/jinzhu/gorm"
)

var GinEngine *gin.Engine
var CasbinEnforcer *casbin.SyncedEnforcer
var Eloquent *gorm.DB
var Rdb *redis.Redis

var (
	Source string
	Driver string
	DBName string
)

// go-admin Version Info
var Version string

func init() {
	Version = "1.1.2"
}

var (
	Logger        *glog.Logger
	DBLogger      *glog.Logger
	RequestLogger *glog.Logger
)

package database

import (
	"context"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/container/pool"
	"github.com/go-kratos/kratos/pkg/log"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"go-admin/global"
	"time"
)

func Setup(driver string) {
	dbType := driver
	if dbType == "mysql" {
		var db = new(Mysql)
		db.Setup()
	}

	if dbType == "redis" {
		NewRedis()
	}

	// TODO： 如果需要sqlite3请开启下面注释
	// if dbType == "sqlite3" {
	//	var db = new(SqLite)
	//	db.Setup()
	// }

	if dbType == "postgres" {
		var db = new(PgSql)
		db.Setup()
	}
}

func NewRedis() {
	cfg := redis.Config{
		Addr:         "127.0.0.1:6379",
		DialTimeout:  xtime.Duration(90 * time.Second),
		ReadTimeout:  xtime.Duration(90 * time.Second),
		WriteTimeout: xtime.Duration(90 * time.Second),
		SlowLog:      xtime.Duration(90 * time.Second),
		Proto:        "tcp",
	}
	cfg.Config = &pool.Config{
		Active:      10,
		Idle:        5,
		IdleTimeout: xtime.Duration(90 * time.Second),
	}
	global.Rdb = redis.NewRedis(&cfg)
	err := ping(global.Rdb)
	if err != nil {
		panic(err)
	}
	log.Info("Redis 启动成功")

}

func ping(r *redis.Redis) (err error) {
	if _, err = r.Do(context.TODO(), "SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}

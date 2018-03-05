package models

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/reechou/holmes"
	"github.com/reechou/wedding-chatroom/config"
)

var x *xorm.Engine

func InitDB(cfg *config.Config) {
	var err error
	x, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4",
		cfg.DBInfo.User,
		cfg.DBInfo.Pass,
		cfg.DBInfo.Host,
		cfg.DBInfo.DBName))
	if err != nil {
		holmes.Fatal("Fail to init new engine: %v", err)
	}
	//x.SetLogger(nil)
	x.SetMapper(core.GonicMapper{})
	x.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	// if need show raw sql in log
	if cfg.IfShowSql {
		x.ShowSQL(true)
	}

	x.SetMaxIdleConns(0)

	// sync tables
	if err = x.Sync2(new(Chatroom),
		new(ChatroomMember),
		new(BroadcastMessage)); err != nil {
		holmes.Fatal("Fail to sync database: %v", err)
	}

	for i := 0; i < CHATROOM_MESSAGE_NUM; i++ {
		if err = x.Sync2(&ChatroomMessage{ChatroomId: int64(i)}); err != nil {
			holmes.Fatal("Fail to sync database: %v", err)
		}
	}
}

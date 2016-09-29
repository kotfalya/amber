package main

import (
	"github.com/golang/glog"
	"github.com/kotfalya/db"
)

func main() {
	database := db.NewDB(db.DefaultConfig())

	glog.Infoln("1")
	cmd := db.NewGetCmd("test", "level", "stable")
	cmd.Process(database)
	res := cmd.Res().(*db.StrCmdRes)

	if res.Err() != nil {
		glog.Errorln(res.Err())
	} else {
		glog.Infoln(res.Val())
	}
}

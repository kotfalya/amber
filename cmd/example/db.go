package main

import (
	"github.com/kotfalya/db"
	"github.com/golang/glog"
)

func main() {
	database := db.NewDB(db.DefaultConfig())
	
	cmd := db.NewGetCmd("test")
	cmd.Process(database)
	
	res := cmd.Res().(db.StrCmdRes)
	
	return res.Val(), res.Err()
	
	
	res, err := database.Get("test")
	if err != nil {
		glog.Errorln(err)
	} else {
		glog.Infoln(res)
	}
}

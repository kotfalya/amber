package main

import (
	"flag"

	"github.com/golang/glog"
	db "github.com/kotfalya/amber"
)

func main() {
	flag.Parse()
	database := db.NewDB("node1", db.DefaultConfig())

	setCmd := db.NewSetCmd("test", "test val", "level", "local")
	setCmd.Process(database)
	setRes := setCmd.BoolRes()

	if setRes.Err() != nil {
		glog.Errorln(setRes.Err())
	} else {
		glog.Infoln(setRes.Val())
	}

	cmd := db.NewGetCmd("test", "level", "local")
	cmd.Process(database)
	res := cmd.StrRes()

	if res.Err() != nil {
		glog.Errorln(res.Err())
	} else {
		glog.Infoln(res.Val())
	}
}

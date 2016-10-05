package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/kotfalya/db"
)

func main() {
	flag.Parse()
	database := db.NewDB(db.DefaultConfig())

	cmd := db.NewGetCmd("test", "level", "local")
	cmd.Process(database)
	res := cmd.StrRes()

	if res.Err() != nil {
		glog.Errorln(res.Err())
	} else {
		glog.Infoln(res.Val())
	}
}

package main

import (
	"github.com/golang/glog"
	"github.com/kotfalya/store/store"
)

func main() {
	st := store.NewStore()

	if err := st.Set("test_key", "test_value"); err != nil {
		glog.Errorln(err)
		return
	}

	if val, err := st.Get("test_key"); err != nil {
		glog.Errorln(err)
	} else {
		glog.Infoln(val)
	}

	if val, err := st.Get("test_key2"); err != nil {
		glog.Errorln(err)
	} else {
		glog.Infoln(val)
	}
}

package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/kotfalya/store/store"
	"time"
)

var iterations int = 1000
var workers int = 10

func main() {
	st := store.NewStore()

	stop := make(chan struct{})
	req := make(chan int, workers)
	//res := make(chan int, workers * 2)

	go func() {
		for i := 0; i < iterations; i++ {
			req <- i
		}
	}()

	var t1, t2 time.Time

	t1 = time.Now()

	go func() {
		for index := range req {
			//go func(index int) {
			if err := st.Set(fmt.Sprintf("test_key_%d", index), fmt.Sprintf("test_value_%d", index)); err != nil {
				glog.Errorln(err)
			}

			//res <- index
			//}(i)
			if index == iterations-1 {
				break
			}
		}

		close(stop)

	}()

	//go func() {
	//	count := 0
	//	for range res {
	//		count++
	//
	//		if count == iterations {
	//			close(stop)
	//			close(res)
	//			return
	//		}
	//	}
	//}()

	<-stop
	t2 = time.Now()
	glog.Infoln(t1)
	glog.Infoln(t2)

	if val, err := st.Get("test_key_2"); err != nil {
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

package main

import (
	"fmt"

	"github.com/kotfalya/db/example"
)

func main() {
	//stop := make(chan struct{})

	//go func() {
	//	time.Sleep(time.Second)
	//	fmt.Println("1")
	//}()
	//
	//go func() {
	//	time.Sleep(time.Second * 2)
	//	fmt.Println("2")
	//	close(stop)
	//}()

	st := example.NewMyStore()

	if _, err := st.Set("test", "haha"); err != nil {
		fmt.Println(err)
	}

	if str, err := st.Get("test"); err != nil {

		fmt.Println(err)
	} else {
		fmt.Println(str)
	}

}

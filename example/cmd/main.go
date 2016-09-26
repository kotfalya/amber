package main

import (
	"fmt"
	"github.com/kotfalya/store/example"
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
	str, err := st.MyGet("test")
	fmt.Println(str)
	fmt.Println(err)

}

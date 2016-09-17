package main

import (
	"github.com/kotfalya/store/key"
	"github.com/kotfalya/store/store"
)

func main() {
	st := store.NewStore()

	key := key.NewStringKey("hi", "ha")

	st.Get()
}

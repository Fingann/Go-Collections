package main

import (
	"fmt"
	"github.com/Fingann/Go-Collections/collection"
)

func main() {
	lst := Collection.NewList([]string{"hello,", " world"})
	val, _ := lst.Get(0)
	fmt.Print(val)
	val, _ = lst.Get(1)
	fmt.Println(val)

}

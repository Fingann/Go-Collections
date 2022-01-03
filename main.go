package main

import (
	"fmt"

	Collections "github.com/Fingann/Go-Collections/collections"
)

func main() {
	lst2 := Collections.NewList([]string{"!"})
	lst := Collections.NewList([]string{"hello,", " world"})
	lst.AddRange(lst2)
	val, _ := lst.Get(0)
	fmt.Print(val)
	val, _ = lst.Get(1)
	fmt.Print(val)
	val, _ = lst.Get(2)
	fmt.Println(val)

	item, err := lst.Find(func(needle string) bool {
		return needle == "hello,"
	})
	if err != nil {
		fmt.Errorf("error: %s", err)
	}
	fmt.Printf("Found string: \"%s\", in the list \n", item)

}

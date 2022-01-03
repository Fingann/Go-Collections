package main

import (
	"fmt"

	Collections "github.com/Fingann/Go-Collections/collections"
)

func main() {
	// Create a new List and retrieve values
	lst := Collections.NewList([]string{"hello,", " world"})
	val, _ := lst.Get(0)
	fmt.Print(val)
	val, _ = lst.Get(1)
	fmt.Print(val)

	// Add ranges to list
	lst2 := Collections.NewList([]string{"!", "!"})
	lst.AddRange(lst2)
	val, _ = lst.Get(2)
	fmt.Println(val)

	// Search for values within lists using predicate
	item, err := lst.Find(func(needle string) bool {
		return needle == "hello,"
	})
	if err != nil {
		fmt.Errorf("error: %s", err)
	}
	fmt.Printf("Found string: \"%s\", in the list \n", item)

}

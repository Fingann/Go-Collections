package main

import (
	"fmt"

	"github.com/Fingann/Go-Collections/dictionary"
	"github.com/Fingann/Go-Collections/list"
	"github.com/Fingann/Go-Collections/queue"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func PrintItem[T any](item T) error {
	fmt.Println(item)
	return nil
}

func main() {
	// Some Errors are returned, but they are not handled here.
	// This is just a demo.

	list.From(NewPoint(1, 2), NewPoint(3, 4), NewPoint(5, 6)).
		ForEach(PrintItem[Point])

	/// Lists

	slice := []string{"hello", " ", "world"}
	// Create a new List and retrieve values
	list1 := list.From(slice...).Add("!")
	// Get by index
	val, _ := list1.Get(0)
	fmt.Println(val) // "hello"

	// Add range takes an enumerable so a any collection can be usedworld
	list1.AddRange(list.New[string]().Add("!").Add("!"))
	list1.FindAll(func(object string) bool {
		return object == "!"
	}).Count() // 2

	// ForEach lets us loop the list and perform an action on each item
	list1.ForEach(func(item string) error {
		fmt.Print(item)
		return nil // return an error to stop the loop
	}) // "Hello Collections!!!"

	list1.GetEnumerable().
		Where(func(item string) bool { return item == "Hello" }).
		ForEach(PrintItem[string]) // "Hello", returns an error if the action returns an error

	// Search for values within lists using predicate
	item, _ := list1.Find(func(needle string) bool {
		return needle == "hello"
	})
	fmt.Printf("Found string: \"%s\"\n", item)

	/// Queue

	queue := queue.New[string]().
		Enqueue("hello").
		Enqueue("queue")

	queue.Contains("hello") // true
	queue.Count()           // 2
	front, _ := queue.Peek()
	fmt.Printf("Queue Peek: %s\n", front) // "hello"

	val, _ = queue.Dequeue()
	fmt.Println("Queue Dequeue: ", val) // "hello"

	/// Dictionary

	dict := dictionary.New[string, Point]()
	dict.AddKeyValue("key1", NewPoint(1, 2))
	dict.AddKeyValue("key2", NewPoint(3, 4))

	dict.ContainsKey("key1") // true

	point1, _ := dict.Get("key1")
	fmt.Println("Dictionary Get: ", point1) // (1, 2)

	point2, _ := dict.Get("key2")
	fmt.Println("Dictionary Get: ", point2) // (3, 4)

}

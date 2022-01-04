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

func NewPoint(x, y int) *Point {
	return &Point{X: x, Y: y}
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func main() {
	list.From([]*Point{NewPoint(1, 2), NewPoint(3, 4), NewPoint(5, 6)}).
		ForEach(func(p *Point) {
			fmt.Println(p)
		})

	fmt.Println("Lists: ")

	// Create a new List and retrieve values
	list1 := list.From([]string{"hello,"})
	list1.Add(" world")
	val, _ := list1.Get(0)
	fmt.Print(val) // "hello,"
	val, _ = list1.Get(1)
	fmt.Println(val) // " lists"
	val, _ = list1.Get(1)
	fmt.Println(val) // "world"

	// Create a new List
	list2 := list.New[string]()
	list2.Add("!")

	// Add list2 to list 1
	list1.AddRange(list2)
	fmt.Println("Added secound list to first list")

	fmt.Println("Looping and printing list using foreach:")
	// loop the list and print the values
	list1.ForEach(func(item string) { fmt.Printf(item) }) // "Hello, lists!"

	fmt.Println("\nSearching List for \"Hello,\"")
	// Search for values within lists using predicate
	item, _ := list1.Find(func(needle string) bool {
		return needle == "hello,"
	})
	fmt.Printf("Found string: \"%s\"\n", item)

	///////// Queue ////////////
	fmt.Println("\nQueue:")

	queue := queue.New[string]()
	queue.Enqueue("hello")
	queue.Enqueue("queue")

	front, _ := queue.Peek()
	fmt.Printf("Queue Peek: %s\n", front) // "hello"

	val, _ = queue.Dequeue()
	fmt.Println("Queue Dequeue: ", val) // "hello"
	val, _ = queue.Dequeue()
	fmt.Println("Queue Dequeue: ", val) // "queue"

	///////// Dictionary ////////////

	fmt.Println("\nDictionary:")
	dict := dictionary.New[string, *Point]()
	dict.AddKeyValue("key1", NewPoint(1, 2))
	dict.AddKeyValue("key2", NewPoint(3, 4))

	point1, _ := dict.Get("key1")
	fmt.Println("Dictionary Get: ", point1) // (1, 2)

	point2, _ := dict.Get("key2")
	fmt.Println("Dictionary Get: ", point2) // (3, 4)

}

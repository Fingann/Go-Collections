package main

import (
	"fmt"

	List "github.com/Fingann/Go-Collections/collections/list"
	Queue "github.com/Fingann/Go-Collections/collections/queue"
)

func main() {
	// Create a new List and retrieve values
	list1 := List.From([]string{"hello,", " lists"})
	val, _ := list1.Get(0)
	fmt.Print(val) // "hello,"
	val, _ = list1.Get(1)
	fmt.Println(val) // " lists"

	// Create a new List
	list2 := List.New[string]()
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

	queue := Queue.New[string]()
	queue.Enqueue("queue")
	queue.Enqueue("hello")

	front, _ := queue.Peek()
	fmt.Printf("Queue Peek: %s\n", front) // "hello"

	val, _ = queue.Dequeue()
	fmt.Println("Queue Dequeue: ", val) // "hello"
	val, _ = queue.Dequeue()
	fmt.Println("Queue Dequeue: ", val) // "queue"

}

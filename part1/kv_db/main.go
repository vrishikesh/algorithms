package main

import (
	"fmt"

	"kv_db/database"
)

// Create a redis db implementation where
// you can Set key value in constant time
// you can Get key value in constant time
// you must Evict the least accessed key value if buffer is full
// you must make the most recent accessed key value at the top
// you should be able loop in sequence according to the most recent to least recent key value
func main() {
	d := database.NewDatabase[string](3)

	d.Set("hi", "hi")
	d.Set("bye", "bye")
	d.Set("hello", "hello")
	d.Set("hey", "hey")

	fmt.Println(d.Get("hi"))
	fmt.Println(d.Get("bye"))
	fmt.Println()
	fmt.Println()

	iter := d.Iterator()
	for iter.HasNext() {
		fmt.Println(iter.GetNext())
	}
}

package main

import "fmt"

type Stuff struct {
	Thing string
}

func main() {
	var greeting = "HELLO"
	var a string
	fmt.Printf("%s Go Actions v2\n", greeting)
}

func doThings() {
	fmt.Printf("does stuff")
}

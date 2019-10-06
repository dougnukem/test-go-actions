package main

import "fmt"

type Stuff struct {
	Thing string
}

func main() {
	var greeting = "HELLO"
	fmt.Printf("%s Go Actions v2\n", greeting)
}

func stuff() {
	fmt.Printf("does stuff")
}

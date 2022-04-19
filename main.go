package main

import "fmt"

func repeatMe(words ...string) {
	fmt.Printf("%T\n", words)
}

func main() {
	repeatMe("nico", "jc", "me", "coco")
}
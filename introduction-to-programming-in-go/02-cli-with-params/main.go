package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "Gophers", "The name of the person or animal that we are greeting")

	flag.Parse()

	fmt.Println("Hello, "+*name+"!", "")
}

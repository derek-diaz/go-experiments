package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Bob")
	fmt.Println(message)
}

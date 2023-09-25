package main

import (
	"fmt"
	"hello_world/greetings" // Replace "your_package_name" with the actual package name where greetings.go is located
)

func main() {
	name := "Taufik" // Change this to the name you want to greet
	message := greetings.Hello(name)
	fmt.Println(message)
}

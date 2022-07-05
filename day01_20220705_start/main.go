package main

import (
	"day01/service"
	"fmt"
)

func main() {

	fmt.Println("Hello GoLand")

	sayHello := service.SayHello()

	fmt.Println(sayHello)

}

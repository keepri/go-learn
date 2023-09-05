package main

import "fmt"

func main() {
	const name = "Me"
	const rate = 30.5

	msg := fmt.Sprintf("Hi %s, your rate is %.1f", name, rate)

	fmt.Println(msg)
	fmt.Printf("%T value %v\n", 123, 123)
	fmt.Printf("%T value %v\n", "123", "123")
	fmt.Printf("string %s\n", "123")
	fmt.Printf("decimal int %d\n", 123)
	fmt.Printf("float %f\n", 123.42)
	fmt.Printf("float %.2f\n", 123.42069)
}

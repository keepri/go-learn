package main

import "fmt"

func main() {
	s := ""
	i := 10
	f := 0.1
	bool := false
	complex := 0.69 + 0.5i
	multi, line := 123, "cool"
	const constVar = "hello from const"

	fmt.Printf("%T %v\n", s, s)
	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", f, f)
	fmt.Printf("%T %v\n", bool, bool)
	fmt.Printf("%T %v\n", complex, complex)
	fmt.Println(multi, line)
	fmt.Println("convert to int from float", int(f))
	fmt.Println(constVar)
}

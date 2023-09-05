package main

import (
	"fmt"

	"internal/arrays"
	"internal/closures"
	condsfuncs "internal/conds-funcs"
	"internal/datetest"
	"internal/errors"
	helloworld "internal/hello-world"
	"internal/loops"
	"internal/maps"
	"internal/pointers"
	"internal/strings"
	"internal/structs"
	"internal/variables"
)

func main() {
	fmt.Println("\n===================== hello-world =====================")
	helloworld.Run()

	fmt.Println("\n===================== strings =====================")
	strings.Run()

	fmt.Println("\n===================== datetest =====================")
	datetest.Run()

	fmt.Println("\n===================== variables =====================")
	variables.Run()

	fmt.Println("\n===================== conds-funcs =====================")
	condsfuncs.Run()

	fmt.Println("\n===================== closures =====================")
	closures.Run()

	fmt.Println("\n===================== errors =====================")
	errors.Run()

	fmt.Println("\n===================== loops =====================")
	loops.Run()

	fmt.Println("\n===================== arrays =====================")
	arrays.Run()

	fmt.Println("\n===================== structs =====================")
	structs.Run()

	fmt.Println("\n===================== maps =====================")
	maps.Run()

	fmt.Println("\n===================== pointers =====================")
	pointers.Run()
}

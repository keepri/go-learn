package main

import (
	"fmt"
	"strings"
)

func main() {
	len_one := 10
	len_two := 20
	fmt.Printf("len 1 and 2 are %d %d\n", len_one, len_two)

	if len_one == len_two {
		fmt.Println("len one equals len two")
	} else if len_one > len_two {
		fmt.Println("len one bigger")
	} else {
		fmt.Println("len two bigger")
	}

	if val, _, _ := subtract(1, 5); val < 0 {
		fmt.Printf("val %d is less than 0\n", val)
	}

	fmt.Println(comb("Hello", "world!", strings.Join))
	x, y := coords()
	fmt.Println("coords", x, y)

	nums := []int{1, 2, 3, 4}
	fmt.Println("sum of nums", nums, "equals", sum(nums...))
}

func subtract(a int8, b int8) (int8, int8, int8) {
	return a - b, a, b
}

func comb(
	a, b string,
	with func(elems []string, sep string) string,
) string {
	return with([]string{a, b}, ", ")
}

// named return values
func coords() (x, y int) {
	// x, y := 0, 0 - return values get init to 0
	return // no need to specify return values
}

// variatic
func sum(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

package loops

import "fmt"

func Run() {
	sum := 0
	i := 0
	for ; i < 100; i++ {
		sum += i
	}
	fmt.Println("sum", sum)

	sum = 0
	i = 0
	for ; ; i++ {
		if i >= 100 {
			break
		}
		sum += i
	}
	fmt.Println("sum", sum)

	sum = 0
	i = 0
	for i < 100 {
		sum += i
		i++
	}
	fmt.Println("sum", sum)

	n := 1
	var arr [10]int
	for i := range arr {
		arr[i] = n
		n++
	}
	fmt.Println("arr", arr)
}

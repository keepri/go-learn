package arrays

import "fmt"

func Run() {
	// array
	var arr [10]int
	for i := 1; i <= 10; i++ {
		arr[i-1] = i
	}
	fmt.Println("arr", arr)

	// slice
	slice_one := arr[3:7]
	slice_two := arr[:]
	fmt.Printf("slice %v with len %d, capacity %d\n", slice_one, len(slice_one), cap(slice_one))
	fmt.Printf("slice %v with len %d, capacity %d\n", slice_two, len(slice_two), cap(slice_two))

	slice_one = append(slice_one, 69420)
	fmt.Printf("slice %v with len %d, capacity %d\n", slice_one, len(slice_one), cap(slice_one))
	fmt.Println("appending 3 more elements to slice...")
	slice_one = append(slice_one, 69420, 69420, 69420)
	fmt.Printf("slice %v with len %d, capacity %d - capacity got auto incremented\n", slice_one, len(slice_one), cap(slice_one))
}

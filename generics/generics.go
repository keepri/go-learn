package generics

import "fmt"

func Run() {
	s_one, s_two := split[string]([]string{"69", "420"})
	fmt.Printf("split %v %v\n", s_one, s_two)
}

func split[T any](v []T) ([]T, []T) {
	switch len(v) {
	case 0:
		zero_val := []T{}
		return zero_val, zero_val

	case 1:
		return v, v

	default:
		mid := len(v) / 2
		return v[:mid], v[mid:]
	}
}

package maps

import "fmt"

type key struct {
	first_name string
	last_name  string
}

func Run() {
	// var m map[string]int
	// m := make(map[string]int)
	// or
	m := map[string]int{"test": 1}

	m["test"] = 1
	m["test"] = 2
	m["testtest"] = 3
	fmt.Printf("map: %v has %d entries\n", m, len(m))

	if v, ok := m["test"]; ok {
		fmt.Printf("key 'test' has value %d\n", v)
	}

	for k, v := range m {
		fmt.Printf("map value %s: %d\n", k, v)
	}

	clear(m)
	fmt.Printf("map: %v\n", m)

	custom_key := key{first_name: "clfxc", last_name: "mario"}
	hits := make(map[key]int)
	hits[custom_key]++
	hits[custom_key]++
	hits[key{first_name: "luigi", last_name: "yes"}]++
	for k, v := range hits {
		fmt.Printf("hits map %s: %d\n", k, v)
	}
}

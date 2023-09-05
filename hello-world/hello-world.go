package helloworld

import (
	"fmt"

	str "internal/mystrings"
)

func Run() {
	fmt.Println(str.Reverse("Hello, world!"))
}

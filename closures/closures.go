package closures

import "fmt"

func Run() {
	concat := concatter()
	concat("Hello,")
	concat("world!")
	concat("How")
	concat("are")
	fmt.Println(concat("you?"))
}

func concatter() func(string) string {
	sentence := ""
	return func(word string) string {
		sentence += word + " "
		return sentence
	}
}

package structs

import "fmt"

func Run() {
	empty_struct := myStruct{}
	s := myStruct{foo: 7, bar: "baz"}
	o := myOtherStruct{baz: s}
	oo := myOtherStruct{baz: myStruct{bar: "cool", foo: 42}, biq: "test"}
	oo.inline.bla = "bloo"
	emb := embedded{myStruct: myStruct{foo: 3, bar: "z"}}

	me := person{name: "me", message: "phone home"}
	cat := animal{name: "jinx", sound: "MEOW"}
	shout(&me)
	shout(&cat)

	// annon struct
	inline_struct := struct{ foo string }{
		foo: "hello",
	}

	fmt.Println("embedded struct", emb)
	fmt.Println("my empty struct", empty_struct)
	fmt.Println("my struct", s)
	fmt.Println("my other struct", o)
	fmt.Println("my other other struct", oo)
	fmt.Println("inline struct", inline_struct)
}

type myStruct struct {
	foo int
	bar string
}

type myOtherStruct struct {
	biq string
	baz myStruct
	// annon struct
	inline struct {
		bla string
	}
}

type embedded struct{ myStruct }

type rect struct{ width, height int }

// struct methods
func (r *rect) area() int {
	return r.width * r.height
}

type person struct {
	name    string
	message string
}

func (p *person) say_hi() {
	fmt.Printf("%s said %s!\n", p.name, p.message)
}

type animal struct {
	name  string
	sound string
}

func (a *animal) say_hi() {
	fmt.Printf("%s said %s!\n", a.name, a.sound)
}

// interfaces
type interacter interface {
	say_hi()
}

func shout(entity interacter) {
	if _, ok := entity.(*person); ok == true {
		fmt.Println("We have a person!")
	} else {
		fmt.Println("We have an animal!")
	}

	// or
	// switch entity.(type) {
	// case person:
	// 	fmt.Println("We have a person!")
	// case animal:
	// 	fmt.Println("We have something better!")
	// default:
	// 	fmt.Println("Unrecognised entity")
	// }

	entity.say_hi()
	fmt.Println()
}

package pointers

import "fmt"

type circle struct {
	x, y, radius int
}

func (self *circle) grow() {
	self.radius *= 5
}

func (self circle) grow_no_ptr() {
	self.radius *= 5
}

func Run() {
	x := 5
	fmt.Println("x", x)

	y := x
	fmt.Println("y", y)

	z := &x
	*z = 6
	fmt.Println("x", x)
	fmt.Println("z", z, *z)

	c := circle{
		x:      1,
		y:      3,
		radius: 5,
	}
	fmt.Println("circle radius before grow", c.radius)
	fmt.Println("calling grow...")
	c.grow()
	fmt.Println("calling grow_no_ptr...")
	c.grow_no_ptr()
	fmt.Println("circle radius after grow", c.radius)
}

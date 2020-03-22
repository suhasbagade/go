/**
- Interface is a collection of methods.
- Interface is a custom type.
- We can't create an instance of a interface, whereas we can create a variable of an interface type.
- Two types of interface :- Static interface & Dynamic interface.
- Static interface is
**/

package main

import "fmt"

type tank interface {
	TankArea() float64
	TankVolume() float64
}

type tankdim struct {
	radius float64
	height float64
}

func (t tankdim) TankArea() float64 {
	return 2*t.radius*t.height + 2*3.14*t.radius*t.radius
}

func (t tankdim) TankVolume() float64 {
	return 3.14 * t.radius * t.radius * t.height
}

func main() {

	var tan tank

	tan = tankdim{10, 14}

	fmt.Println(tan)
	fmt.Println(tan.TankArea())
	fmt.Println(tan.TankVolume())

	// Type switching

}

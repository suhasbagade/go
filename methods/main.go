/**
- We can two methods with the same name but different receiver types in the same package.
- We can't have two functions withe same name in the same package.
- Method can be of non-struct type, as long as type and the method definitions are present in the same package.
**/

package main

import "fmt"

// struct type receiver
type author struct {
	name     string
	branch   string
	articles int
	salary   int
}

func (a author) show() {
	fmt.Println("Name is :", a.name)
	fmt.Println("Branch is :", a.branch)
	fmt.Println("Articles is :", a.articles)
	fmt.Println("Salary is :", a.salary)

}

// non-struct type receiver
type data int

func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

// pointer type receiver

type engineer struct {
	name     string
	branch   string
	articles int
	salary   int
}

func (a *engineer) showPointer(b string) {
	(*a).branch = b
	fmt.Println("Name is :", a.name)
	fmt.Println("Branch is :", a.branch)
	fmt.Println("Articles is :", a.articles)
	fmt.Println("Salary is :", a.salary)
}

func main() {

	// Receiver of type struct
	res := author{"XYZ", "E&C", 200, 100000}
	res.show()

	// Receiver of type data (non-struct)
	input1 := data(100)
	input2 := data(1000)

	res1 := input1.multiply(input2)
	fmt.Println(res1)

	// Receiver ot type pointer
	res3 := engineer{"ABC", "CIVIL", 300, 100000}

	fmt.Println("BEFORE :: Branch is :", res3.branch)

	p := &res3

	p.showPointer("MECHANICAL")

}

package main

import "fmt"

// interface
type Stringer interface {
	String() string
}

type Person struct {
	Name string
}

func (p *Person) String() string {
	return p.Name
}

// Interface values with nil underlying values - no underlying value
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// NIL interface values - holds neither value nor concrete type.
type Q interface {
	M()
}

// empty interface - An empty interface may hold values of any type. (Every type implements at least zero methods.)
// interface{}

func main() {

	// interface
	var i Stringer = &Person{Name: "John Doe"}
	fmt.Println(i.String())

	// Interface values with nil underlying values
	var j I
	var t *T // null pointer - doesnt hold underlying value
	j = t
	j.M()
	j = &T{"hello"}
	j.M()

	// NIL interface values
	// var k Q
	// k.M()

	// empty interface
	var o interface{}
	
	o = 42
	fmt.Println(o)

	o = "hello"
	describe(o)
}

// The main() function creates an interface value i that holds a reference to a Person value. The fmt.Println() function calls the String() method on the i interface value. The String() method is implemented by the Person struct, so the fmt.Println() function prints the name of the Person value.



func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
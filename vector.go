package main

import (
	"fmt"
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func newVector(x, y float64) Vector {
	return Vector{X: x, Y: y}
}

func (v Vector) Abs(other Vector) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) String() string {
	// ^ If its a pointer its better for performance to use a pointer receiver, but if its a value type, then it is better to use a value receiver.
	return fmt.Sprintf("Vector(%f, %f)", v.X, v.Y)
}

func myPrint(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Printf("%d\n", x.(int))
	case float64:
		fmt.Printf("%f\n", x.(float64))
	case string:
		fmt.Printf("%s\n", x.(string))
	case fmt.Stringer:
		fmt.Printf("%s\n", x.(fmt.Stringer).String())
	default:
		fmt.Printf("%T\n %+v", x, x)
		// usually would print: main.Vector \n {X:3 Y:4}
		// but because of the String() method, it will print: Vector(3.000000, 4.000000)
		// why? because of case fmt.Stringer: fmt.Printf("%s\n", x.(fmt.Stringer).String()) in the switch statement, which calls the String() method of the Vector struct.
	}
}

//func main() {
//	v := newVector(3, 4)
//	myPrint(v)
//}

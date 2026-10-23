package main

/*
func fib() func() int {
	prev, curr := 0, 1
	return func() int {
		next := prev + curr
		prev = curr
		curr = next
		return prev
	}
}

type fibGenerator struct {
	prev int
	curr int
}

func newFibGenerator() *fibGenerator {
	return &fibGenerator{
		prev: 0,
		curr: 1,
	}
}

func (g *fibGenerator) next() int {
	next := g.prev + g.curr
	g.prev = g.curr
	g.curr = next
	return g.prev
}

func main() {
	nextFib := fib()

	//max before overflow is 92

	println("Using closure-based generator:")
	for i := 0; i < 5; i++ {
		println(nextFib())
	}

	println("Using struct-based generator:")
	gen := newFibGenerator()
	for i := 0; i < 5; i++ {
		println(gen.next())
	}
}
*/

//functions - stack
//inside the function - queue

//see code in teacher's git repo

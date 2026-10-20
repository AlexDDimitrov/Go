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

type ExpressionType uint

const (
	ValueType = ExpressionType(iota)
	NegateType
	BinOpType
)

type Expression struct {
	expType ExpressionType
	BinOp   *BinOp
	val     int
	negate  *Expression
}

func (Expression Expression) Eval() int {
	return
}

//creating a new name for byte
type BinOpSymbol byte

const (
	Plus     = BinOpSymbol('+')
	Minus    = BinOpSymbol('-')
	Multyply = BinOpSymbol('*')
)

type BinOp struct {
	left  *Expression
	right *Expression
	op    string
}

func (binOp BinOp) Eval() int {
	switch binOp.op {
	case Plus:
		return binOp.left.Eval() + binOp.right.Eval()
	case Minus:
		return binOp.left.Eval() - binOp.right.Eval()
	case Multyply:
		return binOp.left.Eval() * binOp.right.Eval()
	}
	panic
}

func main() {

}

//see teachers git for the code
//hw model variables, add handle in Eval, finish memory and define in Exp - variable uint

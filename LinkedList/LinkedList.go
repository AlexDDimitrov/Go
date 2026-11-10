package LinkedList

//Double linked list

type Node[Elem any] struct {
	elem Elem
	next *Node
	prev *Node
}

func NewNode[Elem any](elem Elem) * Node[Elem] {
	return &Node {
		elem: elem,
		next: nil,
		prev: nil
	}
}

func (node *Node[Elem]) 

type List[Elem any] struct {

}

//Get it from repo.  
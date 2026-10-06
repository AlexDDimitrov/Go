package main

import "fmt"

func dmemory() {
	p := new(int)
	// p is a pointer to an int, and it points to a newly allocated zero value of type int.
	fmt.Println(*p) // prints 0, the zero value of int

	*p = 5
	fmt.Println(*p)

	p = nil
	// p is now a nil pointer, which means it doesn't point to any valid memory location. Dereferencing a nil pointer will cause a runtime panic.

	// Memory leak isn't happening because there is garbage collection in Go, which automatically frees up memory that is no longer in use. When p is set to nil, the memory that was previously allocated for the int value is no longer referenced and can be reclaimed by the garbage collector.:
	x := 9

	//Leaks can happen if you have a pointer that is still referencing memory that is no longer needed, and you don't set it to nil or otherwise release the reference. In this case, the memory will not be reclaimed by the garbage collector, leading to a memory leak.

	p = &x //redirecting p to point the address of x, the previous memory allocated for the int value is now eligible for garbage collection since there are no references to it.

	fmt.Println(p, *p)

	//a := make([]int, length, capacity)

	a := make([]int, 10)
	// will create a slice of length 10 and capacity 10, with all elements initialized to the zero value of int (which is 0).
	fmt.Println(a)

	a = make([]int, 0, 10) // Allocates a slice of length 0 and capacity 10, with no elements initialized.
	// will create a slice of length 0 and capacity 10, with no elements initialized.
	fmt.Println(a)

	a = nil // freeing the memory allocated for the slice, and setting it to nil allows the garbage collector to reclaim that memory if there are no other references to it.
	fmt.Println(a, len(a), cap(a))

	fmt.Println(8, "abc") // prints 8 abc, because the first argument is an int and the second argument is a string, so they are printed as separate values with a space in between added.
}

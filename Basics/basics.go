// Масив от числа, искаме да извасим поредните числа

package main // папка с файлове, съдържащи фунции; Функции с голяма буква могат да бъдат достъпвани навсякъде

import "fmt"

// Hello, world
func main() {
	fmt.Println("Hello, world")
	variable_declaration()
	output_functions()
	data_types()
	arrays()
	slices()
	logical_ops()
	conditions()
	loops()
	functions("Hola, soi Dora")
	structures()
	interfaces()
}

// Variables
func variable_declaration() {
	// 1. with the keyword var (if we do not supply a type, the compiler does it for us)
	var author string = "Emily Bronte"
	fmt.Println(author)
	// 2. with the operator := but the compiler chooses the type
	book := "Wuthering Heights"
	fmt.Println(book)
	// 3. if not value is assaigned, the variables take the initial values of each type - int = 0, string = "", bool = false
	var a int
	var b string
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// var vs :=
	// - the first inside and outside func, the second only inside
	// - the first does not demand for initialization, whilst it is nesessary in the second

	// 4. Multiple Declaration
	author, book, year := "Suzanne Collins", "The Hunger Games", "2009"
	fmt.Println(author + book + year)

	// 5. Constants -> = and UPPERCASE
	const PAINTER string = "Alexandre Cabanel"
}

// Output functions
func output_functions() {
	// 1. Print() prints without adding any symbol such as \n, etc., in their default format; if we want a specieal symbol -> "text", "\n"
	const PAINTER string = "Alexandre Cabanel"
	sentance := "is a French painter"
	fmt.Print(PAINTER+" "+sentance, ", ")

	// 2. Printft demands for formatted text %v is fot the value, %T is for the type
	sentance2 := "famous for his Romantic painting The Fallen Angel, 1847"
	fmt.Printf("%v.", sentance2)

	// 3. Println - prints with a /n in the end

}

// Data Types
func data_types() {
	// bool (Boolean)
	var are_the_exersices_fun bool = false
	// string (String values)
	var statement string = "When I stepped out, into the bright sunlight, from the darkness of the movie house, I had only two things on my mind: Paul Newman and a ride home."
	// integer numerics - default --- int, int8, int16, int32, int64
	var year int = 1965
	// float numerics - float32, default --- float64
	var rating float64 = 6.0

	fmt.Printf("%v %v %v %v", are_the_exersices_fun, statement, year, rating)
}

// Array
func arrays() {
	// declaration
	// 1. with var -> var name = [length]datatype{values}'with [] it deifnes its length
	var books = [3]string{"The Hunger Games", "Cathcing Fire", "Mockingjay"}
	var series = [...]string{"The Hunger Games", "Cathcing Fire", "Mockingjay"}
	// 2. with :=
	books2 := [3]string{"The Hunger Games", "Cathcing Fire", "Mockingjay"}
	series2 := [...]string{"The Hunger Games", "Cathcing Fire", "Mockingjay"}

	fmt.Println(books)
	fmt.Println(series)
	fmt.Println(books2)
	fmt.Println(series2)

	// arr1 := [5]int{} //not initialized
	// arr2 := [5]int{1:10,2:40} // index:value
	// 0 10 40 0 0
	// len(arr1) -> 5

}

// Slices
func slices() {
	// similar to arrays, but more flexible
	// - stores values but can grow and shrink
	// dynamic array

	// 1. []datatype{values}
	bookshelf := []string{"1", "2"}
	fmt.Println(len(bookshelf)) // size
	fmt.Println(cap(bookshelf)) // capacity

	// 2. from array - adding - [start:end]
	var subjects = [...]string{"Literature", "World History", "Art History"}

	school_subjects := subjects[0:3] // but the capcity becomes from 0 to the end - 3
	fmt.Println(school_subjects)
	// 3. make() -- name := make([]type, length, capacity) -- not giving values though, if  undefined capacity then c = length
	history_eras := make([]string, 5, 8)
	fmt.Println(history_eras)

	// appending - slice = append(slice, el, el2)
	history_eras = append(history_eras, "Ancient times", "Medivial times")
	// copy(destination, sorce) if d < l it does not take all elements

	copy(bookshelf, school_subjects)
	fmt.Println(bookshelf)

}

// Logical Operators
func logical_ops() {
	// AND = &&
	// OR = ||
	// NOT = !
}

//Conditions
func conditions() {
	// if condition {}
	// NO BRACKETS
	/*
		switch expression {
			case x,y:
			// code block
			case z:
			...
			default:
			// code block
			}
	*/

	day := 5

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

}

// For loops
func loops() {
	// for init; cond; step {}
	// continue skips
	// break exits

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// range - easily to go throygh an array, slice, map
	/*
	 	for index, value := range array/slice {
		}
	*/

	fruits := [3]string{"apple", "orange", "banana"}
	for i, val := range fruits {
		fmt.Printf("%v\t%v\n", i, val)
	}
}

// Functions and parameters
func functions(name string) int {
	// parameters -- (var type)
	// func Name(params) type { return output}
	// --- type is the return type
	return 1 + 2

	// func myFunction(x int, y int) (result int) -- here e name the var we will return; in return there is no need to say what we are actually returning

	/*


		func myFunction(x int, y int) (result int) {
			result = x + y
			return
			}
	*/

	// We can return multiple things, even with different types
	/*
		func myFunction(x int, y string) (result int, txt1 string) {
			result = x + x
			txt1 = y + " World!"
			return
			}
	*/

}

type Book struct {
	title  string
	author string
	year   int
}

// Structures
func structures() {
	// for declaring we use the keywords type and struct; we declare them outside funcs

	var book1 Book

	book1.title = "Straving Games"
	book1.author = "Grace Everdeen"
	book1.year = 2025

	// We can pass them as function arguments
	// func printPerson(pers Person)
	//   printPerson(pers2)
}

// Interfaces - declared outside funcs

type Calculate interface {
	Add(a int, b int) int
	Divide(a int, b int) int
}

type Calculator struct{}

func (c Calculator) Add(a int, b int) int {
	value := a + b
	return value
}

func (c Calculator) Divide(a int, b int) int {
	value := a - b
	return value
}

func interfaces() {
	// contains functions but not their code
	// struct but for funcs

	// type NAME inteface
	var calculator Calculate
	calculator = Calculator{}

	fmt.Println(calculator.Add(1, 2))
	fmt.Println(calculator.Divide(1, 2))

}

package main

import "fmt"

/*
A pointer value is the address of a variable ie, this is where a value being assigned to a variable is being stored. WIth pointers, the values assigned to a variable can be manipulated indirectly without using the variable.

Referencing
Referencing is the process of obtaining the memory address of a variable. In Go, you use the & operator to get the address of a variable.

Dereferencing
Dereferencing is the process of accessing the value stored at the memory address held by a pointer. In Go, you use the * operator to dereference a pointer.
*/

func main() {
	var x int  // Declare an integer variable x
	var p *int // Declare a pointer variable p that can hold the address of an integer

	p = &x  // Reference: Assign the address of x to p
	*p = 45 // Dereference: Assign the value 45 to the location p points to (which is x)

	fmt.Print(p) // Print the memory address stored in p
}

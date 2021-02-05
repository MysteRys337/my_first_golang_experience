//to compile any go project...
/*
* go build [name.go] 			  -> create executable
* go run [name.go]   			  -> just runs the application
* go install [application folder] -> creates a executable in the 'bin' folder
 */

//Name for the package
package main

//packages imported to the project
import (
	"fmt"     // Basic golang package ( Standard I/O)
	"strconv" // Conversion to String package
)

//global variable that is only visible to the package(lower case)
var i int = 32

//global variable that is globaly visible(upper case)
var J int = 2

//main function(first to be executed)
func main() {
	//ways of assigning variables(obs: every variable decleared in a function are only visible inside of the function(block scope))

	/* //This was commented because it's not possible to compile codes were values are not used
	var a int = 1
	a = 1

	//creating multiple variables that are releated to one another
	var (
		Gustavo string = "Gustavo"
	)
	*/

	i := 10        //:= -> walrus operator
	n := int8(i)   //Converting to a 8 bit integer
	fmt.Println(n) //Print line

	/* //Direct conversion of int to String(will print the ascii value of said integer)
	a := string(i)
	fmt.Println(a)
	*/

	//String conversion
	//Ex:
	j := strconv.Itoa(i) // Itoa -> Integer to ASCII
	fmt.Println(j)
	fmt.Printf("%v, %T", i, i) // Print Format -> %v = value , %T = type of value

	//Data types
	//
	//Integer
	//int8
	//int16
	//int32
	//int64
	//
	//Unsigned integer
	//
	//uint8
	//uint16
	//uint32
	//
	//Operations(integer)
	//
	// ( + - * / % ) -> Normal operators , ( & | ^ ) -> boolean operators, (>> <<) -> shift left, shift right respectively
	//
	// Obs: You cannot make operations with different data types
	// Ex: int8 + int32
	// You have to convert it first!
	//
	//Float
	//
	//float32
	//float64
	//
	//Ways of declaring floats
	//
	//3.14
	//13.7e72
	//2.1E14
	//
	//Complex numbers(1+1i)
	//
	//complex64
	//complex128
	//
	// Obs: You can convert float to complex with:
	// complex(5,12) => 5 + 12i
	//Ex:
	c := complex(5, 12)
	fmt.Println(c)
	// To convert the real and imaginary part of complex use:
	// real(n) && imag(n)
	//Ex:
	fmt.Println(real(c) + imag(c))
	//
	//String
	//
	//string
	//
	//Obs: it's possible to treat a string as an array
	//Ex:
	s := "A string"
	fmt.Printf("%v", s[2])
	//
	//Obs: Is possible to convert a String into a byte array
	//Ex:
	b := []byte(s)
	fmt.Println(b)
	//
	//Rune
	//
	//Rune represents a int32 char(but it doesn't have to be 32bits long)
	//Ex:
	var r rune = 'a'
	fmt.Printf("%v,%T", r, r) // Notices that the type of this rune is listed as a int32

}

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

//Declaring constants
const k int16 = 1

//declaring multiple constants
const (
	a = iota // set a counter for enumerated constants(starts with 0)
	b
	c
) //Obs: By not assigning a value to 'b' and 'c' the compiler is going to infer these values based on 'a'
/* So...
*  a = 0
*  b = 1
*  c = 2
 */

//main function(first to be executed)
func main() {
	//ways of assigning variables(obs: every variable decleared in a function is only visible inside of the function(block scope))

	/* //This was commented because it's not possible to compile codes were values are not used
	var a int = 1
	a = 1

	//creating multiple variables that are releated to one another
	var (
		Gustavo string = "Gustavo"
	)
	*/

	i := 10 //:= -> walrus operator
	//Also, notices that I declare a variable named "i" inside and outside of this function, this is called "shadowing" and it's the ability of creating same name
	//variables inside and outside of block scopes

	n := int8(i)                                                //Converting to a 8 bit integer
	fmt.Printf("Converting to a 8 bit integer:\nn = %v\n\n", n) //Print line

	/* //Direct conversion of int to String(will print the ascii value of said integer)
	a := string(i)
	fmt.Println(a)
	*/

	//String conversion
	//Ex:
	j := strconv.Itoa(i)                                                                    // Itoa -> Integer to ASCII
	fmt.Printf("Correct way of converting int to string:\n %v, %T\n%v, %T\n\n", i, i, j, j) // Print Format -> %v = value , %T = type of value

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
	fmt.Printf("Complex numbers in go:\n %v,%T\n\n", c, c)
	// To convert the real and imaginary part of complex use:
	// real(n) && imag(n)
	//Ex:
	fmt.Printf("Converting the real -> %v\nConverting the imaginary -> %v\n\n", real(c), imag(c))
	//
	//String
	//
	//string
	//
	//Obs: it's possible to treat a string as an array
	//Ex:
	s := "Printing a string"
	fmt.Printf("Printing a string's third char -> %v\n\n", s[2])
	//
	//Obs: Is possible to convert a String into a byte array
	//Ex:
	b := []byte(s)
	fmt.Printf("Printing a byte array -> %b\n\n", b)
	//
	//Rune
	//
	//Rune represents a int32 char(but it doesn't have to be 32bits long)
	//Ex:
	var r rune = 'a'
	fmt.Printf("Printing a rune:\n%v,%T\n\n", r, r) // Notices that the type of this rune is listed as a int32

	//Arrays
	//Types of declaration
	//1#
	array1 := [3]int{1, 2, 3}
	fmt.Printf("array1: %v\n\n", array1)

	//2#
	array2 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("array2: %v\n\n", array2)

	//3#
	array3 := make([]int, 3)
	fmt.Printf("array3: %v\n\n", array3)

	//4#
	var array4 [4]int
	fmt.Printf("array4: %v\n\n", array4) //will print an empty array

	//Adding elements to the array:
	array3[0] = 1
	fmt.Printf("New element added to array3: %v\n\n", array3)

	//Printing the length of 'array3'
	fmt.Printf("array3's length: %v\n\n", len(array3))
	//Obs: len will print the number of elements an array can hold, and not how many items were added
	//Notices that "fmt.Printf("array3's length: %v\n\n", len(array3))" will print '3' instead of '1'

	//Bidimensional arrays
	var arrayXY [3][3]int
	arrayXY[0] = [3]int{1, 2, 3}
	arrayXY[1] = [3]int{4, 5, 6}
	arrayXY[2] = [3]int{7, 8, 9}
	fmt.Printf("Bidimensional array: %v\n\n", arrayXY)

	//Obs: When you create an array in golang, you don't create a reference to the existing array, instead you create a new one
	//Ex:
	array5 := [...]int{1, 2, 3}
	array6 := array5
	array6[1] = 5
	fmt.Printf("array5: %v", array5)
	fmt.Printf("array6: %v", array6)
}

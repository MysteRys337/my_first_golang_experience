/*
*	Learning about constants in golang
 */
package main

import (
	"fmt"
)

//Declaring constants
//OBS: the same lower case, upper case rule on variables works in here
const k int16 = 1

//declaring multiple constants(enumerated constants)
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

func constants() {
	const k int = 5 //shadowing the block scope k with the global scope k
	fmt.Printf("k = %v\n\n", k)

	//Obs: constants must be values that can be found on compile time
	//Ex:
	//const c = math.Sin(1)

	//Obs2: You can use any basic data type to be constants
	const integer = 1
	const float = 3.4
	const text = "Hi"
	fmt.Printf("%v,%v,%v\n\n", integer, float, text)
}

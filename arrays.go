/* Learning about arrays
*
 */

package main

import "fmt"

func arrays() {

	//Arrays
	//Types of declaration
	//1#
	array1 := [3]int{1, 2, 3}
	fmt.Printf("array1: %v\n\n", array1)

	//2#
	array2 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("array2: %v\n\n", array2)

	//3#
	var array3 [4]int
	fmt.Printf("array4: %v\n\n", array3) //will print an empty array

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

	//Obs: When you create an array through another array, you don't create a reference to the existing array, instead you create a new one
	//Ex:
	array4 := [...]int{1, 2, 3}
	array5 := array4
	array5[1] = 5
	fmt.Printf("array5: %v", array4)
	fmt.Printf("array6: %v", array5)

	//Slice array
	//This is a expandable array(a List if you will)

	//#1
	array6 := make([]int, 3)
	fmt.Printf("array6: %v\n\n", array6)

	//#2
	array7 := []int{1, 2, 3}
	fmt.Printf("array7: %v\n\n", array7)

	//Obs: when you create a slice through another slice, they'll always be a reference
	//Ex:
	array8 := array7
	array8[1] = 5
	fmt.Printf("array7: %v\n\n", array7)
	fmt.Printf("array8: %v\n\n", array8)

	//How to create a copy of another Slice
	array9 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	array10 := array9[:]   //create a slice with all elements of array9
	array11 := array9[3:]  //create a slice from the 4th element to the end
	array12 := array9[:6]  //create a slice from the first 6 elements
	array13 := array9[3:6] //create a slice with the elements between the 3 and 6 element
	fmt.Printf("array9  = %v\n", array9)
	fmt.Printf("array10 = %v\n", array10)
	fmt.Printf("array11 = %v\n", array11)
	fmt.Printf("array12 = %v\n", array12)
	fmt.Printf("array13 = %v\n\n", array13)

}

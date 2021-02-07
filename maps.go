/*
* Learning about maps in Golang
 */
package main

import "fmt"

func maps() {
	//Maps are data structures that allow users to pair two types of values
	//Ex: pairing a String with a integer
	//Obs: #1 type of map declaration
	monthMap := map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}
	fmt.Println(monthMap)

	//#2
	newMap := make(map[string]int)
	fmt.Println(newMap)
	//Obs: You can access information in maps just like in arrays
	//Ex:
	fmt.Println(monthMap["January"])

	//You can delete information inside of maps using 'delete'
	//Ex:
	delete(monthMap, "February")
	fmt.Println(monthMap)
}

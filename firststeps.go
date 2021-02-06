/* HOW TO RUN YOUR FIRST GOLANG PROJECT
*
*  First you have to install Go, which is not that hard, but
*  it can be a little bit frustrating,
*
*  If you use arch linux just like myself, you can just download
*  go from the terminal using: "sudo pacman -S go"
*
*  The other option is to use the golang website
*  https://golang.org/dl/
*
*  I'm not going to explain how to setup path variables and things
*  like that, the objective of this repository is just to show you some
*  golang code and how all of this works
*
*  But basically, after you set the path for Go, that's where you I'll have
*  to code, since it's a static programming language. Setup a src/github.com/yourname folder
*  and then you can start your first project, just like this one
 */

//to compile any go project...
/*
* go build [name.go] 			  -> create executable
* go run [name.go]   			  -> just runs the application
* go install [application folder] -> creates a executable in the 'bin' folder
 */

//Name of the package
package main

//packages imported to the project
import "fmt"

//main function(first to be executed)
func main() {
	fmt.Println("Hello go!") //Executing a print in go

	//Learn more about variables
	//variables()

	//Learn more about constants
	//constants()

	//Learn more about arrays
	//arrays()
}

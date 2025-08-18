// function types
// standard or named function
//anonymous function
// function expression or assign function in variable
// higher order function or first class function
// callback function
// variadic function	
// init function
// closure function
// defer function
// receiver function
// IIFE
// =====
//  standard or named function
package main
import "fmt"
func add(x , y int)  {
 fmt.Println("Sum:", x + y)
}
func main() {
	add(10, 20)
}
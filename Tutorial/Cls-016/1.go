// scope kind is 3
package main
import (
    "fmt"
    "example.com/mathlib" // Import the mathlib package
)
var (
	a int = 1
	b int = 2
)
func main() {
	x := 19
	fmt.Println(a, b, x)
	// a, b, x = 4, 5, 6 accesable
	if x>=18{
		p:=10	
		fmt.Println("x is mature")
		fmt.Println(a, b, x, p)
	}
	fmt.Println("x block theke ber hoyeche")
	fmt.Println(a, b, x)
	// p is not accesable here
	// fmt.Println(p) // this will cause an error
	// add(4,7) // Call the add function to demonstrate its functionality
 fmt.Println("Calling add function from mathlib package:")
	mathlib.Add(4, 7) // Call the add function to demonstrate its functionality
	mathlib.Subtract(10, 3) // Call the subtract function to demonstrate its functionality
 }	
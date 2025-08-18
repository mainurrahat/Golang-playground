// init function
package main
import "fmt"
var a = 10

func main() {
	fmt.Println("Main function called")
	fmt.Println("main er vithor a:", a)
}
func init() {
	fmt.Println("Init function called")
	fmt.Println("Init er vithor 1 a:", a)
	a:= 20
	fmt.Println("Init er vithor 2 a:", a)
}
// Output:
// Init function called
// Main function called

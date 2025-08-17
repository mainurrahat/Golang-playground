package main
import "fmt"
var (
	a int = 1
	b int = 2
)
func printNum(num int) {
	fmt.Println("The number is:", num)
}
func add(x int, y int) {
	result := x + y
	// fmt.Println("The sum of", a, "and", b, "is", result)
	printNum(result)
}
func main() {
	add(a,b) // Call the add function to demonstrate its functionality
	
}
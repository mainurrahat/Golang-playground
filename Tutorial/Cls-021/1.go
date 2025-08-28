package main
import "fmt"

func add(x , y int) {
	fmt.Println(x + y)
}
// eita ke bole standerd /named function
func main() {
	fmt.Println("Main function called")
	add(5, 10)
	func (a int, b int) {
		c:=a+b
		fmt.Println(c)
	}(10, 20) // eita ke bole anonymous function
	// immediate invoked function expression (IIFE)

}
func init() {
	fmt.Println("Init function called")
}
package main
import "fmt"
func sum()  {
	add(1,2)
}
func add(a int, b int) {
	c:=a+b
	fmt.Println("added value is:",c)
}
func main() {
	sum()
	fmt.Println("Main function called")
	add:=func (a int, b int){
		c:=a+b
		fmt.Println("multiplied value is:",c)
	}
	add(10,20)  
}

func init() {
	fmt.Println("Init  function called")
}
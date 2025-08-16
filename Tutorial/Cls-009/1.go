package main
import "fmt"
func main() {
	age:= 14
	if age >= 18 {
		fmt.Println("You are eligible to married.")
	} else if age < 18 && age >= 13 {
		fmt.Println("You are not eligible to married bt you can love someone.")
	}else {
		fmt.Println("You are just teenager.")
		
}
    a:=1
	switch a {
	case 1:
	fmt.Println("a is 1")
	case 2:
	fmt.Println("a is 2")
	case 3:
	fmt.Println("a is 3")
	default:
	fmt.Println("a is something else")
		}

}
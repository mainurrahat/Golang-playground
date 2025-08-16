package main
import "fmt"
func welcomeApp() {
	fmt.Println("Welcome to the Go Application!")	
}
func getUserName() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	return name
}
func getTwoNumbers() (int, int) {
	var num1, num2 int
	fmt.Print("Enter two numbers to add: ")
	fmt.Scanln(&num1, &num2)
	return num1, num2
}
func add(x int, y int) int {
	sum := x + y
	return sum
}
func display(name string, sum int) {
	fmt.Println("Hello,", name, "! Glad to have you here.")
	fmt.Println("The sum , is", sum)
	
}
func goodbye(name string) {
	fmt.Println("Thank you for using the application, ", name)
// goodbye message
	fmt.Println("Goodbye,", name, "! Have a great day!")
	fmt.Println("This is the end of the program.")
	fmt.Println("Thank you for using the application!")
	fmt.Println("Goodbye,", name, "! Have a great day!")
	fmt.Println("This is the end of the program.")}
func main() {
    welcomeApp()
	name := getUserName()
	num1, num2 := getTwoNumbers()
	sum:=add(num1, num2)
	display(name, sum)
	goodbye(name)
  }
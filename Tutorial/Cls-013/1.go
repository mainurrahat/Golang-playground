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
	// print welcome message
	welcomeApp()
	// get user name as input
	name := getUserName()
	// greet the user
	// take the 2 numbers as input
	num1, num2 := getTwoNumbers()
	// sum the two numbers
	sum:=add(num1, num2)
	// display the result
	display(name, sum)
	// thank the user
	goodbye(name)

 }
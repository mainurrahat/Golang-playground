package main
import "fmt"
func main() {
	// print welcome message
	fmt.Println("Welcome to the Application!")
	// get user name as input
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	// greet the user
	fmt.Println("Hello,", name, "! Glad to have you here.")
	// ask for user's age
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	var num1 int
	var num2 int
	fmt.Print("Enter two numbers to add: ")
	fmt.Scanln(&num1, &num2)
	// calculate the sum
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)
	// thank the user
	fmt.Println("Thank you for using the application, ", name, "! You are", age, "years old.")
// goodbye message
	fmt.Println("Goodbye,", name, "! Have a great day!")
	fmt.Println("This is the end of the program.")
	fmt.Println("Thank you for using the application!")
	fmt.Println("Goodbye,", name, "! Have a great day!")
	fmt.Println("This is the end of the program.")
	


 }
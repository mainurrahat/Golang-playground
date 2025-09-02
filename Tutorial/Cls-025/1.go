package main
import "fmt"
const a int = 10
var p=100
func call() {
add:= func(x int,y int) {
	fmt.Println("add fun er modde x-",x)
	fmt.Println("add fun er modde y-",y)
	z:=x+y
	fmt.Println("Sum is:",z)
}
add(5,6)
add(p,a)
}
func main() {
	fmt.Println("main fun er modde a-",a)
	fmt.Println("main fun er modde p-",p)
	call()

}
func init() {
	fmt.Println("Hello from init")
}
/*
phase 1:compilation
// tumi jkn install korcho tkn go compiler install hoye geche
// tumi jkn run korcho tkn binary file create hoye jabe
// binary file mane 0 and 1 er combination
//eikhane 1st phase shesh hoilo
***************compilation phase****************
// code segment
// a=10
// call()=func()
// add()=func()
// main()=func()
// init()=func()

phase 2:execution

// ami jodi go run 1.go kori
// then complie kore then 1 name ekta binary file create kore
// then sei binary file ta ke execute kore (internally kore fele)
// ami jodi go build 1.go kori
// then complie kore then 1 name ekta binary file create kore
// but sei binary file ta ke execute kore na
// ami jodi ./1 kori
// then sei binary file ta ke execute kore
 
***************execution phase****************



/*
package main
import "fmt"
func add1(x int, y int) {
	sum1 := x + y
	fmt.Println("The sum of", x, "and", y, "is", sum1)
}
func add2(x int, y int)int {
	sum2:= x + y
    return sum2
}
func getNumbers(x int, y int)(int, int) {
	sum3 := x + y
	mult := x * y
	return sum3, mult
}
func main() {
	a:= 10
	b:= 20
	add1(a, b)
	add1(10, 20)
	sum3:=add2(30, 40)
	fmt.Println("The sum of 30 and 40 is", sum3)
	p,q:=getNumbers(a,b)
	fmt.Println("The sum of", a, "and", b, "is", p)
	fmt.Println("The multiplication of", a, "and", b, "is", q)

}
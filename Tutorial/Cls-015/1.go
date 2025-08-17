// scope kind is 3
package main
import "fmt"
var (
	a int = 1
	b int = 2
)
func main() {
	x := 19
	fmt.Println(a, b, x)
	// a, b, x = 4, 5, 6 accesable
	if x>=18{
		p:=10	
		fmt.Println("x is mature")
		fmt.Println(a, b, x, p)
	}
	fmt.Println("x block theke ber hoyeche")
	fmt.Println(a, b, x)
	// p is not accesable here
	// fmt.Println(p) // this will cause an error
 }	
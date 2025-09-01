package main
import "fmt"
var a int = 10
var b string = "hello"
func add(x int,y int) {
	fmt.Println("add fun er modde x-",x)
	fmt.Println("add fun er modde y-",y)
	fmt.Println("Sum is:",x+y)
}
func main() {
	fmt.Println("main fun er modde a-",a)
	fmt.Println("main fun er modde b-",b)
	add(5,4)
	add(a,10)

}
func init() {
	// a = 20
	// b = "world"
	fmt.Println("Hello from init")
}
// Internal memory |
//  Code Segment | 
// Data Segment |
//  Stack | 
// Heap | 
// GC
// ekdm first e kichu memory allocate kore,-->code segment, data segment,stack,heap

// heap er upor e GC kaj kore(gc mane garbage collector)
// code segment e rakhe code gulo
// add function, main function, init function
// eigulo code segment e rakhe
// data segment e rakhe global variable gulo
// a=10
// b="hello"
// eigulo data segment e rakhe
// then o first e dekeh init function ache naki,
// then init function ta call kore
// then stack e rakhe init function ta ke
// jothotuku allocate korche,othotuku ke bole stack frame
// then init er exucution korbe,
// shesh hoile,stack theke pop kore dibe
// then abar dekeh main function ache naki
// eita code segment e khuje ber kore
// then call kore main function ta ke
// then main function ta ke stack e rakhe(jothotuku allocate korche,othotuku ke bole stack frame)
// then main er exucution korbe
// then main er modde theke add function ke call korche
// then add function ta ke code segment e khuje ber kore
// then add function ta ke stack e rakhe(jothotuku allocate korche,othotuku ke bole stack frame)
// then stack er modde theke add er parameter gulo ke push kore dibe
// then add er exucution korbe
// then add function er kaj sesh hoile,stack theke pop kore dibe
// then abar main function er kaj korbe
// then abar ekta add function ke call korche
// then add function ta ke code segment e khuje ber kore
// then add function ta ke stack e rakhe(jothotuku allocate korche,othotuku ke bole stack frame)
// then stack er modde theke add er add er argumant gula stack e khujbe,ache naki,
//then patabe,jodi stack e na thake tahole data segment e khujbe
// tkn oparation ektu slow hoye jabe
// then add er exucution korbe
// then add function er kaj sesh hoile,stack theke pop kore dibe
// then abar main function er kaj korbe
// then main function er kaj sesh hoile,stack theke pop kore dibe
// then program ta shesh hoye jabe




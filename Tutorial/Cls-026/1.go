package main
import "fmt"
const a=10
var p=100
func outer()func(){
	money:=100
	age:=20
	fmt.Println("outer func")
	fmt.Println("money=",money)
	fmt.Println("age=",age)
	show:=func(){
		money=money+a+p
		fmt.Println("money=",money)
	}
	return show
 }
 func call(){
	incr1:=outer()
	incr1()
	incr1()
	incr2:=outer()
	incr2()
	incr2()


 }
 func main(){
	call()
 }
func init(){
	fmt.Println("=======Bank init========")
}	

/*
phase
1)compile(compile time)
2)execution(run time)
****compile****
const a=10
func outer()
outer er vitor annonymous1 func ta 
func call()
func main()
func init()



****execution****
-->first e init e asbe 
-->stack er modde init run hobe
-->=======Bank init======== eita print korbe
-->init er kaj sheshe stack theke pop hobe
-->tarpor main e asbe
-->stack er modde main run hobe
-->main er modde call() call hobe
-->then o dekhbe main er modde call() ache naki
-->then gloval section e dekhbe(mane sata segment e)
-->then code segment e chk korbe,
-->then code segment e call() er code ache
-->then stack er modde call() run hobe(er jonno jaiga ni niche ram e)
-->call er modde outer() call hobe
-->then o dekhbe call er modde outer() ache naki
-->then gloval section e dekhbe(mane data segment e)
-->then code segment e chk korbe,
-->then code segment e outer() er code ache
-->then stack er modde outer() run hobe(er jonno jaiga ni niche ram e)
-->outer er modde variable gulo declare hobe -->money:=100-->age:=20
-->then print korbe outer func
-->then print korbe money= 100
-->then print korbe age= 20
-->then annonymous function ta show variable e assign hobe(eita stack e refference hobe)
-->then money,age,show sob outer er stack e thakbe
-->then outer er kaj sheshe stack theke pop hobe
-->then money= 210 hobe karon money=100+a(10)+p(100)
-->then print korbe money= 210
-->then incr1() call hobe
-->kintu incr1 er value outer er show function er reference
-->then stack er modde show function run hobe

-->then money= 320 hobe karon money=210+a(10)+p(100)



output
=======Bank init========
outer func
money= 100
age= 20
money= 210
money= 320
outer func
money= 100
age= 20
money= 210
money= 320
*/
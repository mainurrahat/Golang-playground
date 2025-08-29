// First Order Function Vs Higher Order Function
// Parameter, Argument, First-order Function, Higher-order Function শিখে ফেলি।

// 🔹 1. Parameter (প্যারামিটার)

// 👉 Function define করার সময় যে variable নাম দেওয়া হয়, যেটার মাধ্যমে input ধরা হয়।

// 📌 Example (Go):
// package main
// import "fmt"
// func greet(name string) { // <-- এখানে "name" হলো parameter
//     fmt.Println("Hello,", name)
// }


// এখানে name হলো parameter। Function declare করার সময় আমরা বলছি:
// 👉 "এই function input হিসেবে name নামে একটা string নেবে।"

// 🔹 2. Argument (আর্গুমেন্ট)

// 👉 Function call করার সময় যেই actual value দেওয়া হয়।

// 📌 Example:

// func main() {
//     greet("Rahat") // <-- এখানে "Rahat" হলো argument
// }


// name = parameter

// "Rahat" = argument

// 🔹 3. First-order Function

// 👉 যেই function শুধুই value return করে বা কাজ করে, কিন্তু অন্য function কে না নেয়/না দেয়।
// মানে, সাধারণ function।

// 📌 Example:

// func add(a int, b int) int {
//     return a + b
// }

// func main() {
//     result := add(5, 10) // সাধারণ function call
//     fmt.Println(result)
// }


// add হলো first-order function কারণ এটা শুধু calculation করে, function কে argument হিসেবে নেয় না, function কে return ও করে না।

// 🔹 4. Higher-order Function

// 👉 যেই function হয়

// আরেকটা function কে argument হিসেবে নেয়,

// অথবা function কে return করে।

// 📌 Example 1 (function as argument):

// func apply(fn func(int) int, x int) int {
//     return fn(x) // fn হলো function parameter
// }

// func square(n int) int {
//     return n * n
// }

// func main() {
//     result := apply(square, 5) // square function কে argument দিলাম
//     fmt.Println(result)        // 25
// }


// 📌 Example 2 (function return করা):

// func multiplier(factor int) func(int) int {
//     return func(x int) int {
//         return x * factor
//     }
// }

// func main() {
//     double := multiplier(2) // একটা function return করলো
//     fmt.Println(double(10)) // 20
// }

// ✅ Summary

// Parameter = function definition এ input variable নাম

// Argument = function call করার সময় দেয়া actual value

// First-order Function = সাধারণ function, শুধু value নেয়/ফেরত দেয়

// Higher-order Function = function নেয় বা function ফেরত দেয়
package main
import "fmt"

// Parameter vs Argument demo
func greet(name string) { // "name" হলো parameter
    fmt.Println("Hello,", name)
}

// First-order function
func add(a int, b int) int {
    return a + b
}

// Higher-order function (function as argument)
func apply(fn func(int) int, x int) int {
    return fn(x) // fn হলো function parameter
}

func square(n int) int {
    return n * n
}

// Higher-order function (returning function)
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    // Parameter vs Argument
    greet("Rahat") // "Rahat" হলো argument

    // First-order function
    result1 := add(5, 10)
    fmt.Println("add result:", result1)

    // Higher-order (function as argument)
    result2 := apply(square, 5)
    fmt.Println("apply(square, 5):", result2)

    // Higher-order (function returning function)
    double := multiplier(2)
    fmt.Println("multiplier(2)(10):", double(10))
}


// First Order Function Vs Higher Order Function

// Hello, Rahat
// add result: 15
// apply(square, 5): 25
// multiplier(2)(10): 20
🔹 Step by Step Flow (Simulation)
1. Program শুরু

👉 যখন go run দিয়ে রান করবে, Go runtime প্রথমে তোমার package (main) load করে।
main() function খুঁজে বের করে, এবং execution সেখান থেকে শুরু হয়।

2. Function Definitions

প্রথমে compiler সব function (greet, add, apply, square, multiplier) register করে নেয়।
কিন্তু এগুলো এখনো execute হয় না, শুধু store হয়ে থাকে।

3. func main()

Execution এখানে শুরু হয়।

Line:
greet("Rahat")


Function call হলো greet("Rahat")

এখানে "Rahat" হলো argument

Function definition এ name string হলো parameter

তাই ভেতরে যাবে →

fmt.Println("Hello,", name)


Output:

Hello, Rahat

Next:
result1 := add(5, 10)


Call হলো add(5, 10)

Parameters: a=5, b=10

Function ভেতরে →

return a + b // return 15


তাই result1 এ store হলো 15

Print হবে:

add result: 15

Next:
result2 := apply(square, 5)


Call হলো apply(square, 5)

Parameters: fn = square, x = 5

Function apply ভেতরে যাবে →

return fn(x) // এখানে fn = square


So square(5) call হবে →

return 5 * 5 // 25


Result = 25

Print হবে:

apply(square, 5): 25

Next:
double := multiplier(2)


Call হলো multiplier(2)

Parameter: factor = 2

Function multiplier return করছে একটা নতুন function:

return func(x int) int {
    return x * factor // এখানে factor=2 bind হয়ে গেছে (closure)
}


তাই double এখন একটা function, যা input নিলে সেটা 2 দিয়ে গুন করবে।

Next:
fmt.Println(double(10))


double(10) call হলো →

আসলে এটা হলো সেই returned function:

return 10 * 2 // 20


Output হবে:

multiplier(2)(10): 20

🔹 Final Output (পুরো flow শেষে)
Hello, Rahat
add result: 15
apply(square, 5): 25
multiplier(2)(10): 20

✅ Concept Recap

Parameter = function define করার সময় যে নাম দেওয়া হয় (placeholder)

Argument = function call করার সময় যে actual value দেওয়া হয়

First-order function = add() (শুধু value নেয়, value ফেরত দেয়)

Higher-order function (takes function as arg) = apply(square, 5)

Higher-order function (returns function) = multiplier(2) → return করলো একটা নতুন function
🔹 Callback Function কী?

👉 Callback function হলো এমন একটা ফাংশন যেটাকে অন্য ফাংশনের argument হিসেবে পাঠানো হয়, আর সেই ফাংশন ভেতরে গিয়ে কল করে।

অর্থাৎ একটা ফাংশনের কাজ শেষ হলে অন্য একটা ফাংশনকে কল করা হয় → সেইটাকে Callback বলে।

🔹 Go তে Example:
package main

import "fmt"

// callback function type define
func processData(x int, callback func(int)) {
    fmt.Println("Processing data:", x)
    callback(x) // এখানে callback ফাংশন কল হচ্ছে
}

// একটা callback ফাংশন
func printSquare(n int) {
    fmt.Println("Square:", n*n)
}

func main() {
    // এখানে printSquare ফাংশনকে callback হিসেবে পাঠালাম
    processData(5, printSquare)
}

🔹 Output:
Processing data: 5
Square: 25

🔹 Anonymous function দিয়েও callback করা যায়:
package main

import "fmt"

func processData(x int, callback func(int)) {
    fmt.Println("Processing data:", x)
    callback(x)
}

func main() {
    processData(7, func(n int) {
        fmt.Println("Cube:", n*n*n)
    })
}


Output:

Processing data: 7
Cube: 343


✅ Summary:

Callback হলো function যেটা অন্য function এ পাঠানো হয়।

Higher Order Function callback গ্রহণ করতে পারে।

Go তে আমরা সাধারণত callback ব্যবহার করি event, goroutine, channel handling, data processing ইত্যাদিতে।
/*
1.parameter vs argument
2.first order function 
i.standard function or named function
ii.anonymous function or unnamed function
iii.IIFE (Immediately Invoked Function Expression)
iv.function expression
3.higher order function
4.callback function
5.first class citizen=>variable assign data
*/
📌 Short Notes with Code
1. Parameter vs Argument

Parameter → Function er input variable (placeholder).

Argument → Function call er somoy je real value pass kori.

package main
import "fmt"

func greet(name string) { // name = parameter
    fmt.Println("Hello,", name)
}

func main() {
    greet("Rahat") // "Rahat" = argument
}
Output:

Hello, Rahat
2. First Order Function
(i) Standard Function / Named Function
package main
import "fmt"

func add(a, b int) int { // Named function
    return a + b
}

func main() {
    fmt.Println("Sum:", add(3, 4))
}
Output:

5
(ii) Anonymous Function (Unnamed Function)
package main
import "fmt"

func main() {
    mul := func(a, b int) int { // anonymous function assigned to variable
        return a * b
    }
    fmt.Println("Multiplication:", mul(5, 6))
}
Output:

12

(iii) IIFE (Immediately Invoked Function Expression)
package main
import "fmt"

func main() {
    result := func(x, y int) int { // declare + call immediately
        return x - y
    }(10, 3)
    fmt.Println("IIFE Result:", result)
}
Output:

24

(iv) Function Expression
package main
import "fmt"

func main() {
    var divide func(int, int) int // declare function type
    divide = func(a, b int) int {
        return a / b
    }
    fmt.Println("Division:", divide(20, 5))
}
Output:
12

3. Higher Order Function

👉 Function that takes another function as parameter or returns a function.

package main
import "fmt"

func applyOperation(a, b int, op func(int, int) int) int {
    return op(a, b)
}

func main() {
    add := func(x, y int) int { return x + y }
    mul := func(x, y int) int { return x * y }

    fmt.Println("Add:", applyOperation(3, 4, add))
    fmt.Println("Multiply:", applyOperation(3, 4, mul))
}
Output:
7
12

4. Callback Function

👉 Function ke onno function er parameter hisebe pathano, pore call kora.

package main
import "fmt"

func processMessage(msg string, callback func(string)) {
    fmt.Println("Processing:", msg)
    callback(msg)
}

func main() {
    processMessage("Go is awesome!", func(m string) {
        fmt.Println("Callback received:", m)
    })
}
Output:

Processing: Go is awesome!
Callback received: Go is awesome!


5. First-Class Citizen (Function as Data)

👉 Function ke variable e assign kora, argument hisebe pathano, ba return kora possible.

package main
import "fmt"

func main() {
    add := func(a, b int) int { return a + b } // function as data
    operation := add
    fmt.Println("Result:", operation(7, 8))
}
Output:

Result: 15



✅ Recap

Parameter = placeholder variable, Argument = real value

First Order Function → standard, anonymous, IIFE, function expression

Higher Order Function → function ke input/output hishebe use kora

Callback Function → ekta function ke onno function call er moddhe execute kora

First Class Citizen → function ke variable, parameter, return value hisebe use kora
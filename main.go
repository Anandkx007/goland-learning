package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
)

func main() {

	//printing
	fmt.Println("hello")
	fmt.Println(rand.Int())
	fmt.Printf("Printing with %g specifier.\n", rand.Float32())

	//variables
	var p int = 97
	var j float32 = 0.2
	var k float64 = 19.2
	var s string = "hi"
	var n bool = true
	var w complex128 = cmplx.Sqrt(-5 + 12i)

	fmt.Println(p, j, k, s, n, w)

	//var
	var i int
	i = 97
	i = 27
	fmt.Println(i)

	//automatic values
	var it int
	var f float32
	var b bool
	var s2 string

	fmt.Printf("%v %v %v %q\n", it, f, b, s2)

	//type inference
	var ip int = 99
	fmt.Println(ip)

	io := 97

	fmt.Println(io)

	//declaration type 1

	// when u want to use declared variable much later on

	var var1 int

	var1 = 97

	//when golang will not able to identify the vartype

	var var2 int = 90

	//sort hand

	var3 := 80

	fmt.Println(var1, var2, var3)

	var (
		fn  string = "Anand"
		ls  string = "Kumar"
		eId int    = 3
	)

	fmt.Println(fn, ls, eId)

	// Reassigining variables

	var var4 int = 99
	var4 = 45

	fmt.Println(var4)

}

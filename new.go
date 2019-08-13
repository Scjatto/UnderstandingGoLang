package main // Mark the compiler about the file being an executable instead of a library

import ("fmt"
		"math") // Import multiple packages


type expo struct{ // Struct creation
	base int64
	index int64
	result int64
}

func(c expo) add() int64{ // Method of struct expo , call by value
	return c.base + c.index
}

func(c *expo) power(){ // Method of struct expo , call by reference
	c.result = int64(math.Pow(float64(c.base),float64(c.index)))
}

func main(){
	ex1,ex2 := expo{base:2,index:10},expo{base:10,index:3}
	fmt.Println("Initial value of undefined reference for both structures are:",ex1.result,"and",ex2.result) // Undefined values remain 0
	ex1.power() // Void function power called for ex1
	ex2.power() // Void function power called for ex2
	fmt.Println("Sum of base and index for ex1 is: ",ex1.add())
	fmt.Println("Value of the exponent for ex1 is: ",ex1.result) // Value in the pointer changed, No return type needed
	fmt.Println("Sum of base and index for ex2 is: ",ex2.add())
	fmt.Println("Value of the exponent for ex2 is: ",ex2.result) // Value in the pointer changed, No return type needed
}
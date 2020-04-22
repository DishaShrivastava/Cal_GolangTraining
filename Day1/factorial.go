package main

import "fmt"

func main() {

    number := 6 // Change the number here

    fact := 1

	for i := 1; i <= number; i++ {
	
	    fact = fact * i

	}

    fmt.Printf("Factorial of %#v is %#v",number,fact)
}
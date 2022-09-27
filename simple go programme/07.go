// Golang program to show how
// to take input from the user
package main

import "fmt"

// main function
func main() {

	// Println function is used to
	// display output in the next line
	fmt.Println("Enter Your First Name: ")

	// var then variable name then variable type
	var first string

	// Taking input from user
	fmt.Scanln(&first)
	fmt.Println("Enter Second Last Name: ")
	var second string
	fmt.Scanln(&second)

	// Print function is used to
	// display output in the same line
	fmt.Print("Your Full Name is: ")

	// Addition of two string
	fmt.Print(first + " " + second)
}

/* http request https://www.digitalocean.com/community/tutorials/how-to-make-http-requests-in-go
   https://www.youtube.com/watch?v=ru53LpdVHn4&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=25

*/

// get request https://www.youtube.com/watch?v=V-sxFQ0fWlw&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=28

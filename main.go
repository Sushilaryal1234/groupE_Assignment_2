package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println("This assignment belongs to Group E")

	fmt.Println(" Calculate Area of Ractangle: ")
	fmt.Println(calculateareaofractangle())

	fmt.Println(" Check Palindrome: ")
	palindrome()

	fmt.Println("Find Fibonacci: ")
	fibonacciNumber()

	fmt.Println("Generate Random Number: ")
	fmt.Println("The generated random number is: ", generateRandomNumber())

}

// Function Three
// created by Veerpal Kaur(500226407)
// this function will Calculate the area of rectangle by entering length and width.
func calculateareaofractangle() float64 {

	var length, width float64

	fmt.Println("Enter the length of the rectangle: ")
	fmt.Scan(&length)
	fmt.Println("Enter the Width of the rectangle: ")
	fmt.Scan(&width)

	return length * width
}

// Function Seven
// created by Gurcharan Singh(500228108)
// this function will tells weather the entered string is palidrome or not by using strings
func palindrome() {
	var word string

	fmt.Print("Enter the string without space: ")
	fmt.Scan(&word)

	word = strings.ToLower(word)

	n := ""

	for i := len(word) - 1; i >= 0; i-- {
		n = n + string(word[i])
	}

	if n == word {
		fmt.Println("String is Palindrome")
	} else {
		fmt.Println("String is not Palindrome")
	}
}

// Function Two
// created by Sushil Aryal(500226789)
// this function will generate Fibonacci Number by entering number of terms.
func fibonacciNumber() {

	//Declaring variables
	var input int
	var next int
	var first, second int = 0, 1

	//Ask for user input which user wishes to be shown
	fmt.Println("Enter the number of terms in Fibonacci series you want to show: ")
	//Get a user input
	fmt.Scan(&input)
	fmt.Println("Fibonacci Series: ")
	//Calculate and Print Fibonacci series
	for i := 0; i < input; i++ {
		fmt.Printf("%d ", first)
		next = first + second
		first = second
		second = next
	}
}


// Function one
// created by Durpin Thapa Magar(500217688)
// this function will generate a random number between the input minimum and maximum value entered.
func generateRandomNumber() int {
	//Declaring variables to store minimum number for range
	var min, max int

	//Ask for user input
	fmt.Println("Enter minimum range number: ")

	//store minimum number in variable min
	fmt.Scan(&min)

	fmt.Println("Enter maximum range number: ")

	//store maximum number in variable max
	fmt.Scan(&max)

	//Check if max number is less or equal to min number
	if max <= min {
		fmt.Println("Invalid Numbers: Please enter valid min and max numbers")
		return -1
	}
	//Generate random number between max and min numbers
	randomNumber := rand.Intn(max-min) + min

	//return random number to the caller
	return randomNumber
}
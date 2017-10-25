package main

// shorthand for format
import "fmt"

func main() {
	var name string
	name = "Thiago"
	fmt.Println("Hello World")
	fmt.Println("Hello, my name is " + name)
	fmt.Println("Sum is", add(12,10))
	readInput()

	for i := 1; i <= 5; i++ {
	  if i % 2 == 0 {
        // even
        fmt.Println(i, "is even")
      } else {
        // odd
        fmt.Println(i, "is odd")
      }
  }
}

func add(a int, b int)int {
	return a + b
}

func readInput() {
	var input float64
	
	fmt.Print("Enter a number: ")
	fmt.Scanf("%f", &input)
	fmt.Println("The number is", input)

}
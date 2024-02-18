package greeting

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Print(Hello("John"))
}

func Hello(name string) string {
	// return a greeting that embeds the name in a message
	// := operator is a shortcut for declaring and initializing
	// a variable in one line
	message := fmt.Sprintf("Hi, %v. Welcome", name)
	return message
}

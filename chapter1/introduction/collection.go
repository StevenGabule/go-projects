package main

import (
	"fmt"
)

func main() {
	//var x [5]int
	//x[0] = 10
	//x[1] = 20
	//x[2] = 30
	//x[3] = 40
	//x[4] = 50

	//// Initializing an Array with an Array Literal
	//x := [5]int{10, 20, 30, 40, 50}

	//Listing 2-14. Initializing an Array with ...
	//x := [...]int{10, 20, 30, 40, 50, 60, 30}

	// Listing 2-15. Initializing Values for Specific Elements
	x := [5]int{2: 10, 4: 40}
	fmt.Println(x)

	// creating a slice with the make function
	t := make([]int, 5, 10)
	t[0] = 10
	//t[7] = 13 // return an error
	fmt.Println(t)

	// Listing 2-19. Initializing a Slice with a Slice Literal
	d := []int{10, 20, 30, 40, 50}
	fmt.Println(d)

	// Listing 2-22. Slice with the append Function
	xx := []int{10, 20, 30}
	yy := append(xx, 40, 50)
	fmt.Println(xx, yy) // [10 20 30] [10 20 30 40 50]

	// The copy function creates a new slice by copying elements from an existing slice into another slice.
	// Listing 2-23 shows an example of the copy function.
	xa := []int{10, 20, 30}
	ya := make([]int, 4)
	copy(ya, xa)
	fmt.Println(xa, ya) // [10 20 30] [10 20 30 0]

	// Listing 2-24. Slice Length and Capacity
	xd := make([]int, 2, 5)
	xd[0] = 10
	xd[1] = 20
	fmt.Println("===========================================================")
	fmt.Println(xd)                     // [10 20]
	fmt.Println("Length is", len(xd))   // 	Length is 2
	fmt.Println("Capacity is", cap(xd)) // Capacity is 5

	xd = append(xd, 30, 40, 50)         // 	[10 20 30 40 50]
	fmt.Println(xd)                     // Length is 5
	fmt.Println("Length is", len(xd))   // Capacity is 5
	fmt.Println("Capacity is", cap(xd)) // [10 20 30 40 50]
	fmt.Println(xd)                     // Length is 6

	xd = append(xd, 60)
	fmt.Println("Length is", len(xd))   // Capacity is 10
	fmt.Println("Capacity is", cap(xd)) // [10 20 30 40 50 60]
	fmt.Println(xd)
	fmt.Println("===========================================================")

	// Iterating Over Slices
	xxa := []int{10, 20, 30, 40, 50}
	total := 0
	for k, v := range xxa {
		total += v
		fmt.Printf("Index: %d Value: %d\n", k, v)
	}
	fmt.Println("Total: ", total)
	fmt.Println("===========================================================")

	// Listing 2-26. Creating a Map and Iterating it Over a Collection

	// A map named dict is declared, where the string type is specified for the key (type within the []
	// operator) and value:
	dict := make(map[string]string)

	// Values are assigned to the map with the given key (here the key "go" is for the value "Golang"):
	dict["go"] = "Golang"
	dict["cs"] = "CSharp"
	dict["rb"] = "Ruby"
	dict["py"] = "Python"
	dict["js"] = "JavaScript"

	// Finally, iterate over the collection using the range and print key and value of each element in the collection:
	for k, v := range dict {
		fmt.Printf("Key: %s Value: %s\n", k, v)
	}
	// Key: cs Value: CSharp
	// Key: rb Value: Ruby
	// Key: py Value: Python
	// Key: js Value: JavaScript
	// Key: go Value: Golang

	// Note: The data order will vary every time because a map is an unordered collection.

	// Listing 2-27. Accessing the Value of an Element from a Map
	//lan, ok := dict["go"]

	// Listing 2-28. Accessing the Element Value from a Map in an Idiomatic Way
	if lan, ok := dict["go"]; ok {
		fmt.Println(lan, ok)
	}

	// ** Listing 2-29. Defer Statements for Cleaning up Resources
	// session, err := mgo.Dial("localhost") //MongoDB Session object
	// defer session.Close()
	// c := session.DB("task_db").C("categories")
	// code statements using session object

	// ** Listing 2-30 is the code block that calls panic if there is an error while connecting to a database.
	// session, err := mgo.Dial("localhost") // Create MongoDB Session object
	// if err != nil {
	//   panic(err)
	// }
	// defer session.Close()

	// ** Listing 2-31. Recovering from a Panicking Function Using recover
	//func doPanic() {
	//	defer func() {
	//		if e := recover(); e != nil {
	//			fmt.Println("Recover with: ", e)
	//		}
	//	}()
	//	panic("Just panicking for the sake of demo")
	//	fmt.Println("This will never be called")
	//}
	//func main() {
	//	fmt.Println("Starting to panic")
	//	doPanic()
	//	fmt.Println("Program regains control after panic recover")
	//}

	// ** Listing 2-32. Error Handling in Go
	// f, err := os.Open("readme.ext")
	// if err != nil {
	//	 log.Fatal(err)
	// }

	// ** Listing 2-33 is a custom function that returns multiple values, including an error value.
	//func GetById(id string) (models.Task, error) {
	//	var task models.Task
	//	// Implementation here
	//	return task,nil // multiple return values
	//}

	// ** Listing 2-34 is the code block that demonstrates how to call a function that provides an error value.
	// ** Listing 2-34. A caller Function Checks the Error Value
	//task, err:= GetById (“105”)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//Implementation here if error is nil
}

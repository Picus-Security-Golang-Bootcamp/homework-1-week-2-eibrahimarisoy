package main

import (
	"fmt"
	"os"
	"strings"
)

// Define books variable as a slice of strings
var books = []string{
	"alice",
	"wonderland",
	"moby dick",
	"peter pan",
}

// define console information
const info = `
> list : lists all the books
> search <bookName> : queries a book by name
`

func main() {

	// get the arguments from the command line
	args := os.Args[1:]

	// if the user has not provided any arguments
	if len(args) < 1 {
		return
	}

	switch args[0] {

	case "list":
		// list all the books
		list(books)

	case "search":
		// if the user has not provided <bookName>
		if len(args) < 2 {
			return
		}

		// checks bookName is in the books slice and get index
		book := strings.Join(args[1:], " ")
		if i, e := contains(books, book); e {
			fmt.Printf("Book name => %s \n", books[i])
			return
		}
		fmt.Println("We don't have that book")
		fmt.Printf(info)
	}
	// terminates the program
	return
}

// list prints all the books
func list(books []string) {
	for _, v := range books {
		fmt.Println(v)
	}
}

// contains checks if a book is in the books slice if exists return its index
func contains(books []string, book string) (int, bool) {
	for i, v := range books {
		if strings.ToLower(v) == strings.ToLower(book) {
			return i, true
		}
	}
	return 0, false
}

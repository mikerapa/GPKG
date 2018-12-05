package main

import (
	"fmt"
	"glib"
)

func main() {
	fmt.Println("Running")
	firstName, lastName := glib.GetName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}

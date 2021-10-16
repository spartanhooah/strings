package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()

	name := "Hello world"
	fmt.Println("String:", name)

	fmt.Println("\nBytes")

	for i:= 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}

	fmt.Println()

	fmt.Println("\nIndex\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}

	name = "Γειά σου Κόσμε"
	fmt.Println("\nIndex\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}

	fmt.Println()
	fmt.Println("Three ways to concatenate strings")

	h := "Hello, "
	w := "world."

	// concatenation; not super efficient
	myString := h + w
	fmt.Println(myString)

	// using fmt; more efficient
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString)

	// using stringbuilder; most efficient
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())

	fmt.Println()

	// substring
	name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println("Getting a substring")
	fmt.Println(name[0:13])
}
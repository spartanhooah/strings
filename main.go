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

	courseName := "Learn Go for Beginners Crash Course"

	// Need to cast to String to get the letter rather than the rune
	fmt.Println(string(courseName[0]))
	fmt.Println(string(courseName[6]))

	for i := 0; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}

	fmt.Println("Length of courseName is", len(courseName))

	var mySlice []string
	mySlice = append(mySlice, "one")
	mySlice = append(mySlice, "two")
	mySlice = append(mySlice, "three")

	fmt.Println("mySlice has", len(mySlice), "elements")

	fmt.Println("The last element in mySlice is", mySlice[len(mySlice) - 1])

	courses := []string {
		"Learn Go for Beginners Crash Course",
		"Learn Java for Beginners Crash Course",
		"Learn Python for Beginners Crash Course",
		"Learn C for Beginners Crash Course",
	}

	for _, x := range courses {
		if strings.Contains(x, "Go") {
			fmt.Println("Found it in", x, "and index is", strings.Index(x, "Go"))
		}
	}

	newString := "Go is a great programming language. Go for it!"

	fmt.Println(strings.HasPrefix(newString, "Go"))
	fmt.Println(strings.HasPrefix(newString, "Python"))
	fmt.Println(strings.HasSuffix(newString, "!"))
	fmt.Println(strings.Count(newString, "Go"))
	fmt.Println(strings.Count(newString, "fish"))
	fmt.Println(strings.Index(newString, "Go"))
	fmt.Println(strings.LastIndex(newString, "Go"))

	// there's also a ReplaceAll method
	newString = strings.Replace(newString, "Go", "Golang", 1)

	fmt.Println(newString)

	if "Apha" > "Absolute" {
		fmt.Println("A is greater than B")
	} else {
		fmt.Println("A is not greater than B")
	}

	badEmail := " me@here.com  "
	badEmail = strings.TrimSpace(badEmail)
	fmt.Printf("=%s=\n", badEmail)

	str := "alpha alpha alpha alpha alpha"
	fmt.Println(replaceNth(str, "alpha", "beta", 3))

	myString = "This is a clear EXAMPLE of why we search in one case only."

	if strings.Contains(strings.ToLower(myString), "this") {
		fmt.Println("Found it!")
	} else {
		fmt.Println("Didn't find it!")
	}

	fmt.Println(strings.ToLower(myString))
	fmt.Println(strings.ToUpper(myString))
	// EXAMPLE stays all caps
	fmt.Println(strings.Title(myString))

	// EXAMPLE => Example
	fmt.Println(strings.Title(strings.ToLower(myString)))
}

func replaceNth(s, old, new string, index int) string {
	i := 0

	for j := 1; j <= index; j++ {
		x := strings.Index(s[i:], old)

		if x < 0 {
			// didn't find the target string
			break
		}

		// found the target string
		i += x

		if j == index {
			return s[:i] + new + s[i + len(old):]
		}
		
		i += len(old)
	}

	return s
}
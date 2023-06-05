package main

import "fmt"

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated = repeated + character
	}

	return repeated
}

func main() {
	result := Repeat("a")

	fmt.Println(result)
}

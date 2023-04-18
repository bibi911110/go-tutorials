package main

import "fmt"

func myFunction1(p1 string) {
	fmt.Println(p1)
}

func myFunction2(p1 string) string {
	msg := fmt.Sprintf("%s function", p1)

	return msg
}

func myFunction3(p1 string) (string, int) {
	msg := fmt.Sprintf("%s function", p1)
	return msg, 10
}

func main() {
	myFunction1("hello")

	s := myFunction2("Hello")
	fmt.Println(s)
}

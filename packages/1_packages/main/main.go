/*
Fix the bug in the code
*/

// package mailio
package main

import (
	"fmt"
)

func test(text string) {
	fmt.Println(text)
}

func main() {
	test("starting Mailio server")
	test("stopping Mailio server")
}

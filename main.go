package main

import "fmt"

func main() {
	// fmt.Println("hello")

	// fmt.Println("Im in as a collaborator")

	var shazu int
	var althaf int
	sum(input(althaf), input(shazu))
	sum(althaf, shazu)

}
func input(t int) int {
	fmt.Scanf("%d", &t)
	return t
}
func sum(a int, b int) {
	println(a + b)
}

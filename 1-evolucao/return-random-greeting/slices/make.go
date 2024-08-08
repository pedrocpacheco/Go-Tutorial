package main

import "fmt"

var slice []int
var sliceMake = make([]byte, 5, 5)

func main() {
	var slice2 []int
	sliceMake2 := make([]byte, 5, 8)
	makeSucinto := make([]byte, 5)
	fmt.Println(sliceMake2, slice2, makeSucinto)

	fmt.Println(len(sliceMake2), cap(sliceMake2), len(makeSucinto), cap(makeSucinto))

}

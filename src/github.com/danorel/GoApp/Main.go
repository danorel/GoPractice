package main

import "fmt"

const (
	isRoot = 1 << iota
	hasAccessToDelete
	hasAccessToEdit
	hasAccessToView
)

func runConstants() {
	const userRights byte = hasAccessToView | hasAccessToEdit
	fmt.Printf("User roles: %b\n", userRights)
	fmt.Printf("Is user root? Result: %v\n", userRights & isRoot == isRoot)
	fmt.Printf("Has access to view? Result: %v\n", userRights & hasAccessToView == hasAccessToView)
	fmt.Printf("Has access to delete? Result: %v\n", userRights & hasAccessToDelete == hasAccessToDelete)
}

func runArray() {
	// Manipulations with arrays: without & pointing to different cells.
	array := []float32 { 5, 4, 3, 1 }
	var arrayCopy []float32 = array
	fmt.Println("Arrays:")
	fmt.Println(array)
	fmt.Println(arrayCopy)

	sliceCopy := array[:]
	sliceCopy0t2 := array[0:2]
	sliceCopy2tLen := array[2:len(array)]
	fmt.Println("Slice & Array:")
	fmt.Println(array)
	fmt.Println(sliceCopy)
	fmt.Println(sliceCopy0t2)
	fmt.Println(sliceCopy2tLen)

	// Manipulations. Be careful, array and slice pointing to the same data.
	sliceCopy[0] = 3
	fmt.Println("Slice & Array. Same cells:")
	fmt.Println(array)
	fmt.Println(sliceCopy)
}

func main() {
	runConstants()
	runArray()
}
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

func runArraysAndSlices() {
	// Manipulations with arrays: without & pointing to different cells.
	array := []float32 { 5, 4, 3, 1 }
	var arrayCopy []float32 = array
	fmt.Println("Arrays:")
	fmt.Println(array)
	fmt.Println(arrayCopy)

	// Manipulations with slices: cutting the slices with indices as in Python
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

func runMaps() {
	// Map declaration. Way #1.
	someMap := make(map[string]int)

	// Map initialization.
	someMap = map[string]int {"Key": 12313}
	fmt.Println(someMap)

	// Map declaration. Way #2.
	anotherMap := map[string]int {"Key": 123123}
	fmt.Println(anotherMap)

	// Map value extraction.
	fmt.Println(anotherMap["Key"])

	// Map addition.
	anotherMap["AnotherKey"] = 14
	fmt.Println(anotherMap["AnotherKey"])

	// Check Map key order.
	fmt.Println(anotherMap)

	// Delete the key.
	pop, ok := anotherMap["AnotherKey"]
	fmt.Println(pop, ok)
	if ok {
		delete(anotherMap, "AnotherKey")
	}
	fmt.Println(anotherMap)

	// Copying another map
	anotherMapCopy := anotherMap
	fmt.Println(anotherMapCopy)
}

type Doctor struct {
	number int
	actorName string
	companions []Doctor
}

func runStructs() {
	assistantDoctor := Doctor{
		number: 1,
		actorName: "Your Assistant",
		companions: []Doctor {},
	}

	fmt.Println(assistantDoctor)

	aDoctor := Doctor{
		number: 5,
		actorName: "Danyil Orel",
		companions: []Doctor{ assistantDoctor },
	}

	fmt.Println(aDoctor)
}

func main() {
	runConstants()
	runArraysAndSlices()
	runMaps()
	runStructs()
}
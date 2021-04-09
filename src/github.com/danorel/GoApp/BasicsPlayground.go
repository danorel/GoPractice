package main

import (
	"fmt"
	"reflect"
)

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

type Human struct {
	Name string `required max:"100"`
	Surname string
}

type Actor struct {
	Human
	Identifier int
	Pseudonym  string
	Companions []Actor
}

func runStructs() {
	assistantDoctor := Actor{
		Human:		Human{
			Name: "Your",
			Surname: "Assistant",
		},
		Identifier: 1,
		Pseudonym:  "Funny Boy",
		Companions: []Actor{},
	}

	fmt.Println(assistantDoctor)

	aDoctor := Actor{
		Human:		Human{
			Name: "Danyil",
			Surname: "Orel",
		},
		Identifier: 5,
		Pseudonym:  "Pyrokinesis",
		Companions: []Actor{assistantDoctor },
	}

	fmt.Println(aDoctor)

	// Using reflection to get field.
	t := reflect.TypeOf(Human{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

func runLoops() {
	s := []int64 {1, 2, 3}
	for i, v := range s {
		fmt.Println(i, v)
	}
}

func runControlFlow() {
	// Outputs "Middle" after function runControlFlow ends.
	// But before any results will be returned to another function
	fmt.Println("[1]Start")
	defer fmt.Println("[1]Middle")
	fmt.Println("[1]End")

	// LIFO principal of working with defer function calls.
	defer fmt.Println("[2]Start")
	defer fmt.Println("[2]Middle")
	defer fmt.Println("[2]End")
}

func main() {
	runConstants()
	runArraysAndSlices()
	runMaps()
	runStructs()
	runLoops()
	runControlFlow()
}
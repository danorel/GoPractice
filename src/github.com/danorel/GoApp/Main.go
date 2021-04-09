package main

import "fmt"

const (
	isRoot = 1 << iota
	hasAccessToDelete
	hasAccessToEdit
	hasAccessToView
)

func main() {
	const userRights byte = hasAccessToView | hasAccessToEdit
	fmt.Printf("User roles: %b\n", userRights)
	fmt.Printf("Is user root? Result: %v\n", userRights & isRoot == isRoot)
	fmt.Printf("Has access to view? Result: %v\n", userRights & hasAccessToView == hasAccessToView)
	fmt.Printf("Has access to delete? Result: %v\n", userRights & hasAccessToDelete == hasAccessToDelete)
}
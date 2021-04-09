package main

import "fmt"

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func main() {
	var w Writer = ConsoleWriter{}
	n, err := w.Write([]byte{1, 2, 3, 4})
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}
	defer fmt.Println(n)
}
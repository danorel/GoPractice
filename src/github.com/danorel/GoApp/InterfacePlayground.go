package main

import "fmt"

const CharsPerLine = 8

type Writer interface {
	Write([]byte) error
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) error {
	for i, v := range data {
		if i % CharsPerLine == 0 {
			fmt.Println()
		}
		_, err := fmt.Print(string(v))
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var w Writer = ConsoleWriter{}
	err := w.Write([]byte("Hello, Go!"))
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}
}
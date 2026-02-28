package main

import "fmt"

type Epson struct {
}

func (e *Epson) PrintFile() {
	fmt.Println("Printing file using Epson printer")
}

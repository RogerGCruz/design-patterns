package main

import "fmt"

type Hp struct {
}

func (h *Hp) PrintFile() {
	fmt.Println("Printing file using HP printer")
}

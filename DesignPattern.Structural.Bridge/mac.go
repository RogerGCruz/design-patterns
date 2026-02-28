package main

import "fmt"

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Print("Mac printing: ")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(printer Printer) {
	m.printer = printer
}

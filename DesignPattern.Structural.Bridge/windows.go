package main

import "fmt"

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Print("Windows printing: ")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(printer Printer) {
	w.printer = printer
}

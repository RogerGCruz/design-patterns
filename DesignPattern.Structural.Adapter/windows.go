package main

import "fmt"

type Windows struct {
}

func (w *Windows) InsertIntoUsbPort() {
	fmt.Println("Windows: Inserted into USB Port")
}

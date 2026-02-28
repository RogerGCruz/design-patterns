package main

import "fmt"

type WindowsAdapter struct {
	windows *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("WindowsAdapter: Adapting USB to Lightning Port")
	w.windows.InsertIntoUsbPort()
}

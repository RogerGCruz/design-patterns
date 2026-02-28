package main

import (
	"fmt"
)

func main() {
	ak47, _ := getGun("AK47")
	musket, _ := getGun("Musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(gun IGun) {
	fmt.Printf("Gun: %s\n", gun.getName())
	println()
	fmt.Printf("Power: %d\n", gun.getPower())
	println()
}

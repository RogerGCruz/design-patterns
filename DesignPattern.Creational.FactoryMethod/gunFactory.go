package main

import "fmt"

func getGun(gunType string) (IGun, error) {
	switch gunType {
	case "AK47":
		return NewAk47(), nil
	case "Musket":
		return NewMusket(), nil
	default:
		return nil, fmt.Errorf("gun type %s not recognized", gunType)
	}
}

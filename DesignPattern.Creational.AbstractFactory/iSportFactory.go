package main

import "fmt"

type ISportFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportFactory(brand string) (ISportFactory, error) {
	if brand == "Adidas" {
		return &Adidas{}, nil
	}

	if brand == "Nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}

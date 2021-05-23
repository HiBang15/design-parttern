package models

import (
	"fmt"
)

type IsSportsFactory interface {
	MakeShoe() IsShoe
	MakeShort() IsShort
}

func GetSportsFactory(brand string) (IsSportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}

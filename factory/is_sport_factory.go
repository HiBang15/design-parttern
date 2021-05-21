package factory

import (
	"fmt"
	"github.com/HiBang15/design-parttern.git/models"
)

type iSportsFactory interface {
	makeShoe() models.IsShoes
	makeShort() models.IShort
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}
	if brand == "nike" {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
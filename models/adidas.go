package models

type Adidas struct {
}

type AdidasShoe struct {
	Shoe
}

type AdidasShort struct {
	Short
}

func (a *Adidas) makeShoe() IsShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			Logo: "Adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) makeShort() IsShoe {
	return &AdidasShort{
		Short: Short{
			Logo: "Adidas",
			Size: 14,
		},
	}
}

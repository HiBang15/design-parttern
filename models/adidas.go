package models

type Adidas struct {
}

type AdidasShoe struct {
	Shoe
}

type AdidasShort struct {
	Short
}

func (a *Adidas) MakeShoe() IsShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			Logo: "Adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) MakeShort() IsShort {
	return &AdidasShort{
		Short: Short{
			Logo: "Adidas",
			Size: 14,
		},
	}
}

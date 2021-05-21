package models

type Nike struct {
}

type NikeShoe struct {
	Shoe
}

type NikeShort struct {
	Short
}

func (n *Nike) makeShoe() IsShoe {
	return &NikeShoe{
		Shoe: Shoe{
			Logo: "Adidas",
			Size: 14,
		},
	}
}

func (n *Nike) makeShort() IsShoe {
	return &NikeShort{
		Short: Short{
			Logo: "Adidas",
			Size: 14,
		},
	}
}

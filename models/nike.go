package models

type Nike struct {
}

type NikeShoe struct {
	Shoe
}

type NikeShort struct {
	Short
}

func (n *Nike) MakeShoe() IsShoe {
	return &NikeShoe{
		Shoe: Shoe{
			Logo: "Adidas",
			Size: 14,
		},
	}
}

func (n *Nike) MakeShort() IsShort {
	return &NikeShort{
		Short: Short{
			Logo: "Adidas",
			Size: 14,
		},
	}
}

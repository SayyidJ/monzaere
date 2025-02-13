package widget

type BoxConstraint struct {
	MinWidth  float64
	MinHeight float64
	MaxWidth  float64
	MaxHeight float64
}

func (box *BoxConstraint) Loose() BoxConstraint {
	return BoxConstraint{
		MinWidth:  0,
		MinHeight: 0,
		MaxWidth:  0,
		MaxHeight: 0,
	}
}

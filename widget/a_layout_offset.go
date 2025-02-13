package widget

type Offset struct {
	X float64
	Y float64
}

func (o Offset) Add(offset Offset) Offset {
	return Offset{X: o.X + offset.X, Y: o.Y + offset.Y}
}

func (o Offset) Subtract(offset Offset) Offset {
	return Offset{X: o.X - offset.X, Y: o.Y - offset.Y}
}

func (o Offset) Multiply(offset Offset) Offset {
	return Offset{X: o.X * offset.X, Y: o.Y * offset.Y}
}

func (o Offset) Divide(offset Offset) Offset {
	return Offset{X: o.X / offset.X, Y: o.Y / offset.Y}
}

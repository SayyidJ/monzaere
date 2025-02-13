package widget

type Rect struct {
	X float64
	Y float64
	W float64
	H float64
}

func CreateRectFromSizeOffset(size Size, offset Offset) Rect {
	return Rect{
		X: offset.X,
		Y: offset.Y,
		W: size.Width,
		H: size.Height,
	}
}

func (rect *Rect) GetSize() Size {
	return Size{
		Width:  rect.W,
		Height: rect.H,
	}
}

func (rect *Rect) GetOffset() Offset {
	return Offset{
		X: rect.X,
		Y: rect.Y,
	}
}

func (rect *Rect) TopLeft() Offset {
	return rect.GetOffset()
}

func (rect *Rect) BottomLeft() Offset {
	return Offset{
		X: rect.X,
		Y: rect.Y + rect.H,
	}
}

func (rect *Rect) TopRight() Offset {
	return Offset{
		X: rect.X + rect.W,
		Y: rect.Y,
	}
}

func (rect *Rect) BottomRight() Offset {
	return Offset{
		X: rect.X + rect.W,
		Y: rect.Y + rect.H,
	}
}

func (rect *Rect) HitOffset(offset Offset) bool {
	return offset.X >= rect.X && offset.X < (rect.X+rect.W) &&
		offset.Y >= rect.Y && offset.Y < (rect.Y+rect.H)
}

func (r *Rect) HitRect(rect Rect) bool {
	return r.X < rect.X+rect.W &&
		r.X+r.W > rect.X &&
		r.Y < rect.Y+rect.H &&
		r.Y+r.H > rect.Y
}

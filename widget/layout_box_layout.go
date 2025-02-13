package widget

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

// a leaf on widget tree that has renderObject
type RenderBox struct {
	rect         Rect
	renderObject RenderObject
	needsLayout  bool
	needsPaint   bool
	Hitable      bool
}

func NewRenderBox(renderObject RenderObject, offset Offset) (*RenderBox, error) {
	if renderObject == nil {
		return nil, errors.New("renderObject cannot be nil")
	}

	rect := CreateRectFromSizeOffset(Size{Width: 0, Height: 0}, offset)
	renderBox := &RenderBox{
		rect:         rect,
		renderObject: renderObject,
		needsLayout:  true,
		needsPaint:   false,
		Hitable:      false,
	}
	renderObject.Attach(renderBox)
	return renderBox, nil
}

// hit event usually a mouse or tap
func (renderBox *RenderBox) HitTest(offset Offset) bool {
	if !renderBox.Hitable {
		return false
	}
	return renderBox.rect.HitOffset(offset)
}

// called by RenderObject or Parent
func (renderBox *RenderBox) MarkNeedsLayout() {
	renderBox.needsLayout = true
}

func (renderBox *RenderBox) MarkNeedsPaint() {
	renderBox.needsPaint = true
}

func (renderBox *RenderBox) Layout(constraint BoxConstraint) {
	if renderBox.needsLayout {
		size := renderBox.renderObject.Layout(constraint)
		renderBox.rect.W = size.Width
		renderBox.rect.H = size.Height
		renderBox.needsLayout = false
		renderBox.needsPaint = true
	}
}

func (renderBox *RenderBox) Paint(parent *ebiten.Image) {
	if renderBox.needsPaint {
		renderBox.renderObject.Paint(parent)
		renderBox.needsPaint = false
	}
}

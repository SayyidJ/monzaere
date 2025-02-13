package widget

import "github.com/hajimehoshi/ebiten/v2"

// Strategy how to determined area of region
// not always a box, can be rect, can be polygon and circle
// this communcate with widget
type RenderLayout interface {

	// all region can be tested
	Hit(Offset) bool

	// no matter what the shape is, others need to know its box like region
	// this must layouting whatever content it has
	Layout(constraint BoxConstraint) Size

	// althoug this dont care about painting, its simply pass
	// the parent (a image) to where to draw
	// this get called by parent
	Paint(parent ebiten.Image)

	// will layout on the next frame
	MarkNeedLayout()
	IsNeedLayout() bool

	// will paint on the next frame
	MarkNeedPaint()
	IsNeedPaint() bool
}

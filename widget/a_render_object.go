package widget

import "github.com/hajimehoshi/ebiten/v2"

type RenderObject interface {

	// render object must cache this somewhere renderbox to access MarksNeed.Paint | .Layout
	Attach(renderBox *RenderBox)

	// called by Render Box
	Layout(constraint BoxConstraint) Size

	// called if necessary
	Paint(parent *ebiten.Image)
}

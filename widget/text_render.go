package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	et "github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TextRenderObject struct {
	box        *RenderBox
	text       *Text
	renderable bool
	views      []textView
	size       Size
}

func createTextRenderObject(text *Text) *TextRenderObject {
	return &TextRenderObject{
		text:       text,
		renderable: true,
	}
}

func (render *TextRenderObject) Attach(box *RenderBox) {
	render.box = box

}

func (render *TextRenderObject) Layout(constraint BoxConstraint) Size {

	if constraint.MaxWidth <= 0 {
		render.renderable = false
		return Size{}
	}

	if !render.renderable {
		render.renderable = true
	}

	render.views = []textView{}
	start := 0
	for i := 0; i < len(render.text.runes); i++ {
		end := i + 1
		sub := render.text.runes[start:end]
		w, _ := et.Measure(string(sub), render.text.fontFace, 0)
		if w > constraint.MaxWidth {
			if start == i {
				render.views = append(render.views, textView{start: start, end: end})
				start = end
			} else {
				render.views = append(render.views, textView{start: start, end: i})
				start = i
			}
		}

		if i == len(render.text.runes)-1 {
			render.views = append(render.views, textView{start: start, end: len(render.text.runes)})
		}
	}

	size := Size{Width: constraint.MaxWidth, Height: constraint.MaxHeight}
	render.size = size
	return size
}

func (render *TextRenderObject) Paint(parent *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	c := render.text.Option.Color
	r := float32(c.R) / 255 * 2
	g := float32(c.G) / 255 * 2
	b := float32(c.B) / 255 * 2
	a := float32(c.A) / 255 * 2

	op.ColorScale.Scale(r, g, b, a)
	ol := et.LayoutOptions{
		LineSpacing:    1.0,
		PrimaryAlign:   et.AlignStart,
		SecondaryAlign: et.AlignStart,
	}
	opt := &et.DrawOptions{
		DrawImageOptions: op,
		LayoutOptions:    ol,
	}

	for i := 0; i < len(render.views); i++ {
		view := render.views[i]
		part := render.text.runes[view.start:view.end]
		opt.GeoM.Translate(0, (render.text.Option.Size*1.5)*float64(i))
		et.Draw(parent, string(part), render.text.fontFace, opt)
	}

}

func (render *TextRenderObject) Draw(offset Offset) {

}

func (render *TextRenderObject) GetSize() Size {
	return render.size
}

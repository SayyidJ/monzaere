package widget

import (
	"image/color"

	et "github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TextAlign int

const (
	TextAlignLeft TextAlign = iota
	TextAlignRight
	TextAlignCenter
	TextAlignJustify
)

type TextOption struct {
	Size  float64
	Style FontStyle
	Align TextAlign
	Color color.RGBA
}

type textView struct {
	start int
	end   int
}

type Text struct {
	runes     []rune
	Option    TextOption
	renderBox *RenderBox
	fontFace  *et.GoTextFace
}

func NewText(text string, opt TextOption) *Text {
	var size float64 = 14.0
	if opt.Size > 0 {
		size = opt.Size
	}
	face := GetFont(opt.Style, size)
	return &Text{
		runes: []rune(text),
		Option: TextOption{
			Size:  size,
			Style: opt.Style,
			Align: opt.Align,
			Color: opt.Color,
		},
		fontFace: face,
	}
}

func (text *Text) CreateRenderBox(offset Offset) {
	renderObj := createTextRenderObject(text)
	renderBox, err := NewRenderBox(renderObj, offset)
	if err != nil {
		panic(err)
	}

	text.renderBox = renderBox
}

func (text *Text) GetRenderBox() *RenderBox {
	if text.renderBox == nil {
		panic("Text Render Box is used before CreateRenderBox being Called")
	}
	return text.renderBox
}

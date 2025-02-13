package widget

import (
	"bytes"
	"os"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var cmuBold []byte
var cmuItalic []byte
var cmuRegular []byte

var ffCmuBold *text.GoTextFaceSource
var ffCmuItalic *text.GoTextFaceSource
var ffCmuRegular *text.GoTextFaceSource

type FontStyle int

const (
	FontRegular FontStyle = iota
	FontBold
	FontItalic
)

func loadFont() error {
	var err error
	cmuBold, err = os.ReadFile("assets/fonts/cmu_bold.ttf")
	if err != nil {
		return err
	}
	cmuItalic, err = os.ReadFile("assets/fonts/cmu_italic.ttf")
	if err != nil {
		return err
	}
	cmuRegular, err = os.ReadFile("assets/fonts/cmu_normal.ttf")
	if err != nil {
		return err
	}
	return nil
}

func createFontFace() error {
	var err error
	ffCmuBold, err = text.NewGoTextFaceSource(bytes.NewReader(cmuBold))
	if err != nil {
		return err
	}
	ffCmuItalic, err = text.NewGoTextFaceSource(bytes.NewReader(cmuItalic))
	if err != nil {
		return err
	}
	ffCmuRegular, err = text.NewGoTextFaceSource(bytes.NewReader(cmuRegular))
	if err != nil {
		return err
	}
	return nil
}

func InitFont() error {
	var err error
	err = loadFont()
	if err != nil {
		return err
	}
	err = createFontFace()
	if err != nil {
		return err
	}
	return nil
}

func GetFont(fontStyle FontStyle, size float64) *text.GoTextFace {

	var source *text.GoTextFaceSource = ffCmuRegular
	switch fontStyle {
	case FontItalic:
		source = ffCmuItalic
	case FontBold:
		source = ffCmuBold
	}

	return &text.GoTextFace{
		Source: source,
		Size:   size,
	}
}

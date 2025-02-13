package main

import (
	"image/color"
	"log"

	"github.com/SayyidJ/Monzaere/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

var test *widget.Text
var parent *ebiten.Image
var popt *ebiten.DrawImageOptions

func init() {

	err := widget.InitFont()
	if err != nil {
		panic("Gagal init font")
	}
	size := widget.Size{Width: 50, Height: 100}
	parent = ebiten.NewImage(int(size.Width), int(size.Height))
	parent.Fill(color.RGBA{255, 0, 0, 255})
	popt = &ebiten.DrawImageOptions{}

	offset := widget.Offset{X: 100, Y: 100}

	popt.GeoM.Translate(offset.X, offset.Y)
	test = widget.NewText("Hello World", widget.TextOption{
		Color: color.RGBA{255, 255, 255, 255},
	})
	test.CreateRenderBox(offset)
	test.GetRenderBox().Layout(widget.BoxConstraint{MaxWidth: size.Width, MaxHeight: size.Height})
	test.GetRenderBox().Paint(parent)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	screen.DrawImage(parent, popt)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}

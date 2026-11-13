package component

import (
	"math"

	guihandler "github.com/KONshougun/K-game/gui/guiHandler"
	"github.com/hajimehoshi/ebiten/v2"
)

type Image struct {
	*ComponentData

	Img  *ebiten.Image
	Opts *ebiten.DrawImageOptions
}

func NewImage(img *ebiten.Image, area guihandler.Area) *Image {
	i := &Image{
		ComponentData: NewComponentData(area),
		Img:           img,
		Opts:          &ebiten.DrawImageOptions{},
	}

	imgBounds := i.Img.Bounds()
	scaleX := area.Width / float64(imgBounds.Dx())
	scaleY := area.Height /**0.10*/ / float64(imgBounds.Dy())
	scale := math.Min(scaleX, scaleY)
	i.Opts.GeoM.Scale(scale, scale)
	i.Opts.GeoM.Translate(area.X, area.Y)

	return i
}

func (b *Image) GetComponent() *ComponentData {
	return b.ComponentData
}
func (b *Image) Update() error {
	return nil
}

func (b *Image) Draw(screen *ebiten.Image) {
	screen.DrawImage(b.Img, b.Opts)
}

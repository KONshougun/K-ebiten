package container

import (
	"image/color"

	"github.com/KONshougun/K-ebiten/component"
	"github.com/KONshougun/K-ebiten/handlers"

	"github.com/hajimehoshi/ebiten/v2"
)

type QoS int8

const (
	Fast   QoS = -1
	Normal QoS = 0
	Slow   QoS = 1

	QoSCount QoS = 3
)

type Child struct {
	C   component.Component
	QoS QoS
}

type Container interface {
	GetChildren() []component.Component
	AddChild(c ...Child) bool
	RemoveChild(ids ...uint8) bool
}

type ContainerData struct {
	*component.ComponentData
	*component.Background
	*Fade
}

func NewContainerData(area handlers.Area, img *ebiten.Image) *ContainerData {
	return &ContainerData{
		ComponentData: component.NewComponentData(area),
		Background:    component.NewBackground(img, area),
		Fade: &Fade{
			Alpha:  0,
			Speed:  0,
			Active: false,
		},
	}
}

type Fade struct {
	Alpha  float32
	Speed  float32
	Active bool
}

func (f *Fade) FadeIn(speed float32) {
	f.Speed = -speed
	f.Active = true
}
func (f *Fade) FadeOut(speed float32) {
	f.Speed = speed
	f.Active = true
}

func (f *Fade) Update() {
	if !f.Active {
		return
	}

	f.Alpha += f.Speed

	if f.Alpha >= 1 {
		f.Alpha = 1
		f.Active = false
	}

	if f.Alpha <= 0 {
		f.Alpha = 0
		f.Active = false
	}
}
func (f *Fade) Draw(screen *ebiten.Image) {
	if f.Alpha <= 0 {
		return
	}

	overlay := ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())
	overlay.Fill(color.Black)

	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(f.Alpha)

	screen.DrawImage(overlay, op)
}

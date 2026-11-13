package component

import (
	"image/color"
	"math"

	"github.com/KONshougun/K-ebiten/handlers"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type BorderOptions struct {
	BorderColor color.Color
	BorderWidth float64
}

type ProgressBar struct {
	*ComponentData

	MaxValue  *float64
	CurrValue map[*float64][]color.NRGBA

	Border          BorderOptions
	BackgroundColor color.Color
	//BarColor        []color.NRGBA
}

func NewProgressBar(maxValue *float64, currValue map[*float64][]color.NRGBA, border BorderOptions, area handlers.Area) *ProgressBar {
	p := &ProgressBar{
		ComponentData:   NewComponentData(area),
		BackgroundColor: color.Black,
		Border:          border,
		MaxValue:        maxValue,
		CurrValue:       currValue,
	}
	return p
}

func (b *ProgressBar) GetComponent() *ComponentData {
	return b.ComponentData
}
func (b *ProgressBar) Update() error {
	return nil
}
func (b *ProgressBar) Draw(screen *ebiten.Image) {
	// Sfondo
	vector.FillRect(
		screen,
		float32(b.X),
		float32(b.Y),
		float32(b.Width),
		float32(b.Height),
		b.BackgroundColor,
		false,
	)

	// Parte piena
	for value, color := range b.CurrValue {
		rate := *value / *b.MaxValue
		if rate > 1 {
			rate = 1
		}
		vector.FillRect(
			screen,
			float32(b.X),
			float32(b.Y),
			float32(b.Width*rate),
			float32(b.Height),
			b.GetBarColor(value, color),
			false,
		)
	}

	//Bordo
	vector.StrokeRect(
		screen,
		float32(b.X),
		float32(b.Y),
		float32(b.Width),
		float32(b.Height),
		float32(b.Border.BorderWidth),
		b.Border.BorderColor,
		false,
	)
}

func (b *ProgressBar) GetBarColor(value *float64, color []color.NRGBA) color.Color {
	if len(color) == 1 {
		return color[0]
	}
	if len(color) == 3 {
		return b.ThreeColorLerp(value)
	}
	panic("dajjeee")
}

// COLOR
func LerpColor(a, b color.NRGBA, t float64) color.NRGBA {
	return color.NRGBA{
		R: uint8(float64(a.R) + (float64(b.R)-float64(a.R))*t),
		G: uint8(float64(a.G) + (float64(b.G)-float64(a.G))*t),
		B: uint8(float64(a.B) + (float64(b.B)-float64(a.B))*t),
		A: (a.A + b.A) / 2,
	}
}
func (b *ProgressBar) ThreeColorLerp(value *float64) color.NRGBA {
	rate := 1 - *b.MaxValue / *value
	if rate > 1 {
		rate = 1
	}
	// Curva esponenziale
	rate = math.Pow(rate, 2.0)

	if rate <= 0.5 {
		return LerpColor(b.CurrValue[value][0], b.CurrValue[value][1], rate/0.5)
	}

	return LerpColor(b.CurrValue[value][1], b.CurrValue[value][2], (rate-0.5)/0.5)
}

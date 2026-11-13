package handlers

import (
	"bytes"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/text/language"
)

type TextSize byte

const (
	VerySmallSize TextSize = 8
	SmallSize     TextSize = 12
	NormalSize    TextSize = 16
	LargeSize     TextSize = 24
	VeryLargeSize TextSize = 32
)

type TextLineSpacing byte

const (
	NoneSpace      TextLineSpacing = 0
	SmallSpace     TextLineSpacing = 8
	NormalSpace    TextLineSpacing = 16
	LargeSpace     TextLineSpacing = 24
	VeryLargeSpace TextLineSpacing = 32
)

// FONT
type Font string

const (
	NotoSerif = "asset/ui/font/NotoSerif.ttf"
)

func LoadFont(font Font, size float64) text.Face {
	data, err := os.ReadFile(string(font))
	if err != nil {
		panic(err)
	}

	source, err := text.NewGoTextFaceSource(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	return &text.GoTextFace{
		Source:   source,
		Language: language.Italian,
		Size:     size,
	}
}

type Text struct {
	Content string
	Face    text.Face
	Layout  *text.DrawOptions
}

func NewText(str string, c color.Color, size TextSize, space TextLineSpacing, horizontalAlign, verticalAlign text.Align) Text {
	return Text{
		Content: str,
		Face:    LoadFont(NotoSerif, float64(size)),
		Layout: &text.DrawOptions{
			LayoutOptions: text.LayoutOptions{
				LineSpacing:    float64(space),
				PrimaryAlign:   horizontalAlign,
				SecondaryAlign: verticalAlign,
			},
			DrawImageOptions: ebiten.DrawImageOptions{
				ColorScale: colorScale(c),
			},
		},
	}
}
func colorScale(c color.Color) ebiten.ColorScale {
	var scale ebiten.ColorScale
	scale.ScaleWithColor(c)
	return scale
}

func (t Text) Fit(area Area) bool {
	w, h := text.Measure(
		t.Content,
		t.Face,
		t.Layout.LineSpacing,
	)

	return w <= area.Width && h <= area.Height
}

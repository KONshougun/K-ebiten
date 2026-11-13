package component

import (
	"image/color"

	"github.com/KONshougun/K-ebiten/handlers"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Label struct {
	*ComponentData

	handlers.Text
}

func NewLabel(str string, c color.Color, size handlers.TextSize, space handlers.TextLineSpacing,
	horizontalAlign, verticalAlign text.Align, area handlers.Area) *Label {
	return &Label{
		ComponentData: NewComponentData(area),

		Text: handlers.NewText(str, c, size, space, horizontalAlign, verticalAlign),
	}
}

func (l *Label) GetComponent() *ComponentData {
	return l.ComponentData
}
func (l *Label) Update() error {
	return nil
}
func (l *Label) Draw(screen *ebiten.Image) {
	str := l.fitText()
	l.Layout.GeoM.Reset()
	l.Layout.GeoM.Translate(l.X+l.Width/2*float64(l.Layout.PrimaryAlign), l.Y+l.Height/2*float64(l.Layout.SecondaryAlign))
	text.Draw(screen, str, l.Face, l.Layout)
}

func (l *Label) fitText() string {
	const suffix = "..."

	if l.Text.Fit(l.Area) {
		return l.Text.Content
	}

	str := []rune(l.Text.Content)

	for len(str) > 0 {
		// Evita che il testo finisca con un newline prima di "..."
		if str[len(str)-1] == '\n' {
			str = str[:len(str)-1]
		}

		l.Content = string(str) + suffix

		if l.Text.Fit(l.Area) {
			return l.Content
		}

		str = str[:len(str)-1]
	}

	l.Content = suffix
	if l.Text.Fit(l.Area) {
		return l.Content
	}

	return ""
}

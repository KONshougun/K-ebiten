package component

import (
	"image/color"
	"strings"

	"github.com/KONshougun/K-ebiten/handlers"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type LogConsole struct {
	*ComponentData
	*Background
	text handlers.Text
}

func NewLogConsole(background *ebiten.Image, c color.Color, size handlers.TextSize, space handlers.TextLineSpacing,
	horizontalAlign, verticalAlign text.Align, area handlers.Area) *LogConsole {
	l := &LogConsole{
		ComponentData: NewComponentData(area),
		Background:    NewBackground(background, area),
		text:          handlers.NewText("", c, size, space, horizontalAlign, verticalAlign),
	}
	return l
}
func (l *LogConsole) GetComponent() *ComponentData {
	return l.ComponentData
}
func (l *LogConsole) Update() error {
	return l.Background.Update()
}
func (l *LogConsole) Draw(screen *ebiten.Image) {
	l.Background.Draw(screen)
	l.text.Layout.GeoM.Reset()
	l.text.Layout.GeoM.Translate(l.X+l.Width/2*float64(l.text.Layout.PrimaryAlign), l.Y+l.Height/2*float64(l.text.Layout.SecondaryAlign))
	text.Draw(screen, l.text.Content, l.text.Face, l.text.Layout)
}
func (l *LogConsole) AddText(str string) {
	l.text.Content += str + "\n"
	for !l.text.Fit(l.Area) {

		l.RemoveText()
	}
}
func (l *LogConsole) RemoveText() {
	i := strings.Index(l.text.Content, "\n")
	l.text.Content = l.text.Content[i+2:]

	/*f := false
	row := strings.Split(l.text.Content, "\n")[0]
	if row == "" {
		fmt.Printf("l.text.Content: %v\n", l.text.Content)
	}
	l.text.Content, f = strings.CutPrefix(l.text.Content, row+"\n")
	if !f {
		panic("cazxxzzc")
	}*/
}

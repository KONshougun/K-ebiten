package container

import (
	"github.com/KONshougun/K-ebiten/handlers"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GameWidth  = 1280
	GameHeight = 720
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (int, int)

	GetSceneData() *SceneData
}

type SceneData struct {
	*Panel
}

func NewScene(img *ebiten.Image) *SceneData {
	if img == nil {
		panic("oh coglio l'immagine va messa")
	}
	return &SceneData{
		Panel: NewPanel(handlers.Area{
			X:      0,
			Y:      0,
			Width:  GameWidth,
			Height: GameHeight,
		}, img),
	}
}

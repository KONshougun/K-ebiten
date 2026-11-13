package component

import (
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"
	"github.com/hajimehoshi/ebiten/v2"
)

/*type Controller interface {
	Update()
	Draw(screen *ebiten.Image)
}*/

type Component interface {
	Update() error
	Draw(screen *ebiten.Image)

	GetComponent() *ComponentData
}

type ComponentData struct {
	Id uint8
	guihandler.Area
}

func NewComponentData( /*controller Controller,*/ area guihandler.Area) *ComponentData {
	return &ComponentData{
		//Controller: controller,
		Area: area,
	}
}

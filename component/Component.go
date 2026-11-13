package component

import (
	"github.com/KONshougun/K-ebiten/handlers"
	"github.com/hajimehoshi/ebiten/v2"
)

type Component interface {
	Update() error
	Draw(screen *ebiten.Image)

	GetComponent() *ComponentData
}

type ComponentData struct {
	Id uint8
	handlers.Area
}

func NewComponentData(area handlers.Area) *ComponentData {
	return &ComponentData{
		//Controller: controller,
		Area: area,
	}
}

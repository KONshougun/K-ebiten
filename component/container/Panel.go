package container

import (
	"github.com/KONshougun/K-game/gui/component"
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"
	"github.com/hajimehoshi/ebiten/v2"
)

type Panel struct {
	*ContainerData
	children [QoSCount][]component.Component
}

func NewPanel(area guihandler.Area, img *ebiten.Image, children ...Child) *Panel {
	p := &Panel{
		ContainerData: NewContainerData(area, img),
	}
	p.AddChild(children...)
	return p
}
func (p *Panel) AddChild(children ...Child) bool {
	success := true
	for i := range children {
		if p.GetComponent().Area.Contains(children[i].C.GetComponent().Area) {

			//Setto l'id
			var app uint8 = 0
			for _, child := range p.GetChildren() {
				if child.GetComponent().Id >= app {
					app = child.GetComponent().Id + 1
				}
			}

			p.children[children[i].QoS] = append(p.children[children[i].QoS], children[i].C)
		} else {
			success = false
		}
	}
	return success
}

func (p *Panel) RemoveChild(ids ...uint8) bool {
	success := true
	for _, id := range ids {
		app := false
		for qos := range p.children {
			for i := range p.children[qos] {
				if p.children[qos][i].GetComponent().Id == id {
					p.children[qos] = append(p.children[qos][:i], p.children[qos][i+1:]...)
					app = true
					break
				}
			}
		}
		if !app {
			success = false
		}
	}
	return success
}

func (p *Panel) GetChildren() []component.Component {
	c := make([]component.Component, 0)
	for _, children := range p.children {
		c = append(c, children...)
	}
	return c
}
func (p *Panel) Update() error {
	p.Background.Update()
	for qos := range p.children {
		for _, child := range p.children[qos] {
			if err := child.Update(); err != nil {
				return err
			}
		}
	}
	p.Fade.Update()
	return nil
}
func (p *Panel) Draw(screen *ebiten.Image) {
	p.Background.Draw(screen)
	for qos := range p.children {
		for _, child := range p.children[qos] {
			child.Draw(screen)
		}
	}
	p.Fade.Draw(screen)
}

func (p *Panel) GetComponent() *component.ComponentData {
	return p.ComponentData
}

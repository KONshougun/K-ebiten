package container

import (
	"fmt"

	"github.com/KONshougun/K-ebiten/component"
	"github.com/KONshougun/K-ebiten/handlers"

	"github.com/hajimehoshi/ebiten/v2"
)

type ButtonGroup struct {
	*ContainerData
	btns        []*component.Button
	selected    []bool
	MaxSelected int
}

func NewButtonGroup(area handlers.Area, maxSelected int, btns ...*component.Button) *ButtonGroup {
	bGroup := &ButtonGroup{
		ContainerData: NewContainerData(area, nil),
		MaxSelected:   maxSelected,
	}
	for _, btn := range btns {
		bGroup.AddChild(Child{
			C: btn,
		})
	}
	for i := range bGroup.btns {
		bGroup.btns[i].OnClick = func() {
			bGroup.selected[i] = !bGroup.selected[i]
		}
	}

	return bGroup
}
func (btnGroup *ButtonGroup) AddChild(children ...Child) bool {
	success := true
	for _, child := range children {
		btn, ok := child.C.(*component.Button)

		if ok && btnGroup.Area.Contains(btn.GetComponent().Area) {
			if len(btnGroup.btns) > 0 {
				btn.GetComponent().Id = btnGroup.btns[len(btnGroup.btns)-1].GetComponent().Id + 1
			} else {
				btn.GetComponent().Id = 0
			}
			btnGroup.btns = append(btnGroup.btns, btn)
			btnGroup.selected = append(btnGroup.selected, false)
			i := len(btnGroup.btns) - 1
			btnGroup.btns[i].OnClick = func() {
				btnGroup.selected[i] = !btnGroup.selected[i]
			}
		} else {
			fmt.Println("Che hai passato??")
			success = false
		}
	}
	return success
}
func (btnGroup *ButtonGroup) RemoveChild(ids ...uint8) bool {
	success := true
	for _, id := range ids {
		app := false
		for i := range btnGroup.btns {
			if btnGroup.btns[i].GetComponent().Id == id {
				btnGroup.btns = append(btnGroup.btns[:i], btnGroup.btns[i+1:]...)
				btnGroup.selected = append(btnGroup.selected[:i], btnGroup.selected[i+1:]...)
				app = true
				break
			}
		}
		if !app {
			success = false
		}
	}
	return success
}
func (btnGroup *ButtonGroup) GetSelected() []int {
	s := make([]int, 0, len(btnGroup.btns))
	for i := range btnGroup.selected {
		if btnGroup.selected[i] {
			s = append(s, i)
		}
	}

	return s
}

func (g *ButtonGroup) Update() error {
	for _, c := range g.btns {
		if err := c.Update(); err != nil {
			return err
		}
	}
	return nil
}
func (g *ButtonGroup) Draw(screen *ebiten.Image) {
	for _, c := range g.btns {
		c.Draw(screen)
	}
}

func (g *ButtonGroup) GetComponent() *component.ComponentData {
	return g.ComponentData
}

func (g *ButtonGroup) GetChildren() []component.Component {
	children := make([]component.Component, len(g.btns))

	for i, btn := range g.btns {
		children[i] = btn
	}

	return children
}

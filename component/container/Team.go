package container

import (
	"fmt"

	"github.com/KONshougun/K-game/gui/component"
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"
	"github.com/KONshougun/K-game/model/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type Team struct {
	*ContainerData
	Members []*EntityView
}

func NewTeam(team []*entity.EntityData, area guihandler.Area) *Team {
	if area.Height < area.Width/2 {
		panic("area non corretta")
	}
	t := &Team{
		ContainerData: NewContainerData(area, nil),
		Members:       make([]*EntityView, 0, 4),
	}
	t.AddMember(team...)

	return t
}
func (t *Team) AddMember(members ...*entity.EntityData) bool {
	for i := range members {
		if len(t.Members) == 4 {
			return false
		}
		area := guihandler.Area{
			X:      t.Area.X,
			Y:      t.Area.Y + t.Area.Height*0.25,
			Width:  t.Area.Width * 0.25,
			Height: t.Area.Height * 0.75,
		}
		offSet := t.Area.Width * 0.25
		if members[i].EntityType == entity.Character ||
			members[i].EntityType == entity.Character {
			area.X = t.Area.Width - t.Area.Width*0.25
			offSet = -offSet
		}
		//Cerco lo spazio dove inserirlo
		for j := range 4 {

			found := false
			for _, m := range t.Members {
				if m.X == area.X+offSet*float64(j) {
					found = true
					break
				}
			}
			if !found {
				area.X += offSet * float64(j)
				if !t.AddChild(Child{C: NewEntityView(members[i], area)}) {
					fmt.Println("daje")
				}
				break
			}
		}
	}
	return true
}
func (t *Team) AddChild(children ...Child) bool {
	success := true
	for i := range children {
		e, ok := children[i].C.(*EntityView)
		if ok && len(t.Members) < 4 && t.GetComponent().Area.Contains(children[i].C.GetComponent().Area) {

			//Setto l'id
			var app uint8 = 0
			if len(t.Members) > 0 {
				app = t.Members[len(t.Members)-1].GetComponent().Id + 1
			}
			e.GetComponent().Id = app

			t.Members = append(t.Members, e)
		} else {
			success = false
		}
	}
	return success
}

func (t *Team) RemoveChild(ids ...uint8) bool {
	success := true
	for _, id := range ids {
		app := false
		for i := range t.Members {
			if t.Members[i].GetComponent().Id == id {
				t.Members = append(t.Members[:i], t.Members[i+1:]...)
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

func (t *Team) Update() error {
	for _, a := range t.Members {
		a.Update()
	}
	return nil
}
func (t *Team) Draw(screen *ebiten.Image) {
	for _, a := range t.Members {
		a.Draw(screen)
	}
}
func (t *Team) GetComponent() *component.ComponentData {
	return t.ComponentData
}
func (t *Team) GetChildren() []*entity.EntityData {
	team := make([]*entity.EntityData, len(t.Members))
	for i, e := range t.Members {
		team[i] = e.Data
	}
	return team
}

// RAPPORTO BASE:ALTEZZA 1:2
type EntityView struct {
	*Panel
	Data *entity.EntityData
}

func NewEntityView(data *entity.EntityData, area guihandler.Area) *EntityView {

	e := &EntityView{
		Panel: NewPanel(area, nil),
		Data:  data,
	}

	avatar := component.NewImage(guihandler.GetAvatar(data.Name, data.EntityType), guihandler.Area{
		X:      float64(e.Area.X + e.Area.Width*0.08),
		Y:      float64(e.Area.Y + e.Area.Height*0.15),
		Width:  float64(e.Area.Width * 0.85),
		Height: float64(e.Area.Height * 0.81),
	})
	healthBar := component.NewProgressBar(component.HealthBar, data, component.BorderOptions{},
		guihandler.Area{
			X:      area.X + area.Width*0.05,
			Y:      area.Y + area.Height*0.05,
			Width:  area.Width * 0.90,
			Height: area.Height * 0.04,
		},
	)
	staminaBar := component.NewProgressBar(component.StaminaBar, data, component.BorderOptions{},
		guihandler.Area{
			X:      area.X + area.Width*0.05,
			Y:      area.Y + area.Height*0.09,
			Width:  area.Width * 0.90,
			Height: area.Height * 0.02,
		},
	)
	manaBar := component.NewProgressBar(component.ManaBar, data, component.BorderOptions{},
		guihandler.Area{
			X:      area.X + area.Width*0.05,
			Y:      area.Y + area.Height*0.11,
			Width:  area.Width * 0.90,
			Height: area.Height * 0.02,
		},
	)
	e.AddChild([]Child{{healthBar, Normal}, {staminaBar, Normal}, {manaBar, Normal}, {avatar, Normal}}...)

	return e
}

func (e *EntityView) Update() error {
	return e.Panel.Update()
}
func (e *EntityView) Draw(screen *ebiten.Image) {
	e.Panel.Draw(screen)
}

func (e *EntityView) GetComponent() *component.ComponentData {
	return e.ComponentData
}

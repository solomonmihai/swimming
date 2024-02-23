package components

import (
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Model struct {
	Model *rl.Model
	Color *rl.Color
}

func (model Model) Id() kanye.ComponentId {
	return MODEL_ID
}

func NewModel(model *rl.Model) *Model {
	return &Model{
		Model: model,
		Color: &rl.White,
	}
}

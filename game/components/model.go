package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Model struct {
	Model *rl.Model
}

func NewModel(model *rl.Model) *Model {
	return &Model{
		Model: model,
	}
}

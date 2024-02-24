package components

import (
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Model struct {
	Model *rl.Model
	Texture *rl.Texture2D

	Color *rl.Color
	Visible bool
	Debug bool
	DebugColor *rl.Color
}

func (model Model) Id() kanye.ComponentId {
	return MODEL_ID
}

func (model *Model) SetTexture(texture *rl.Texture2D) {
	model.Texture = texture
	rl.SetMaterialTexture(model.Model.Materials, rl.MapDiffuse, *texture)
}

func NewModel(model *rl.Model) *Model {
	return &Model{
		Model: model,
		Color: &rl.White,
		Visible: true,
		Debug: false,
		DebugColor: &rl.Red,
	}
}

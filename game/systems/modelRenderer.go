package systems

import (
	"main/game/components"
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var VEC_UP = rl.NewVector3(0, 1, 0)

type ModelRenderer struct {}

func (mr *ModelRenderer) Update(scene *kanye.BaseScene) {
	rl.BeginMode3D(*scene.Camera)

	for _, ent := range scene.Entities() {

		if !ent.HasComponents(components.TRANSFORM_ID, components.MODEL_ID) {
			continue
		}
	
		model := (*ent.GetComponent(components.MODEL_ID)).(*components.Model)
		transform := (*ent.GetComponent(components.TRANSFORM_ID)).(*components.Transform)

		rl.DrawModelEx(*model.Model, transform.Pos, VEC_UP, transform.Rotation, transform.Scale, *model.Color)
	}	

	rl.EndMode3D()
}

package systems

import (
	"main/game/components"
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ModelRenderer struct {}

func (mr *ModelRenderer) Update(scene *kanye.BaseScene) {
	rl.BeginMode3D(scene.Camera)
	for _, ent := range scene.Entities() {

		model := kanye.GetComponent[components.Model](ent)
		transform := kanye.GetComponent[components.Transform](ent)

		if model == nil || transform == nil {
			continue
		}

		rl.DrawModel(*model.Model, transform.Pos, 20, rl.Red)
	}	
	rl.EndMode3D()
}

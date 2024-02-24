package systems

import (
	"main/game/components"
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ModelRenderer struct {
	kanye.BaseSystem
}

func (mr *ModelRenderer) Update(scene *kanye.BaseScene) {
	rl.BeginMode3D(*scene.Camera)

	mr.IterateEntities(scene, func (ent *kanye.Entity) {
		model := (*ent.GetComponent(components.MODEL_ID)).(*components.Model)
		transform := (*ent.GetComponent(components.TRANSFORM_ID)).(*components.Transform)

		if model.Visible {
			rl.DrawModelEx(*model.Model, transform.Pos, kanye.VEC_UP, transform.Rotation, transform.Scale, *model.Color)
		}

		if model.Debug {
			rl.DrawModelWiresEx(*model.Model, transform.Pos, kanye.VEC_UP, transform.Rotation, transform.Scale, *model.DebugColor)
		}

	}, components.TRANSFORM_ID, components.MODEL_ID)

	rl.EndMode3D()
}

package systems

import (
				"main/game/components"
				"main/kanye"

				rl "github.com/gen2brain/raylib-go/raylib"
)

type DebugSystem struct {
				kanye.BaseSystem
				debug bool
}

func NewDebugSystem() *DebugSystem {
				return &DebugSystem{
								debug: false,
				}
}

func (ds *DebugSystem) Update(scene *kanye.BaseScene) {
				if rl.IsKeyPressed(rl.KeyP) {
								ds.debug = !ds.debug
								ds.IterateEntities(scene, func (ent *kanye.Entity) {
												model := (*ent.GetComponent(components.MODEL_ID)).(*components.Model)
												model.Debug = ds.debug
								}, components.MODEL_ID)
				}

}

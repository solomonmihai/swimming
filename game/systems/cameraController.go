package systems

import (
	"main/game/components"
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CameraController struct {
  kanye.BaseSystem
}

func NewCameraController(scene *kanye.BaseScene) *CameraController {
	cc := &CameraController{}

	pos := rl.NewVector3(0, 30, 30)

	cam := rl.NewCamera3D(pos, rl.Vector3Zero(), kanye.VEC_UP, 45, rl.CameraPerspective)
	scene.Camera = &cam

  rl.DisableCursor()

	return cc
}

func (cc *CameraController) Update(scene *kanye.BaseScene) {
	// rl.UpdateCamera(scene.Camera, rl.CameraFirstPerson)

	cc.IterateEntities(scene, func (ent *kanye.Entity) {
		transform := (*ent.GetComponent(components.TRANSFORM_ID)).(*components.Transform)

		scene.Camera.Position.X = transform.Pos.X
		scene.Camera.Position.Z = transform.Pos.Z - 30 * -1

		scene.Camera.Target.X = transform.Pos.X
		scene.Camera.Target.Z = transform.Pos.Z
	}, components.PLAYER_DATA_ID)

}

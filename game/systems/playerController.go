package systems

import (
	"main/game/components"
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerController struct {
	kanye.BaseSystem
}

func handleControls(camera *rl.Camera) rl.Vector3 {
	dir := rl.Vector3Zero()

	camForward := rl.GetCameraForward(camera)
	camRight := rl.GetCameraRight(camera)

	if rl.IsKeyDown(rl.KeyW) {
		dir = rl.Vector3Add(dir, rl.NewVector3(camForward.X, 0, camForward.Z))
	}

	if rl.IsKeyDown(rl.KeyS) {
		dir = rl.Vector3Add(dir, rl.NewVector3(-camForward.X, 0, -camForward.Z))
	}

	if rl.IsKeyDown(rl.KeyA) {
		dir = rl.Vector3Add(dir, rl.NewVector3(-camRight.X, 0, -camRight.Z))
	}

	if rl.IsKeyDown(rl.KeyD) {
		dir = rl.Vector3Add(dir, rl.NewVector3(camRight.X, 0, camRight.Z))
	}

	dir = rl.Vector3Normalize(dir)

	return dir
}

func handlePhysics(transform *components.Transform, data *components.PlayerData, dir rl.Vector3) {
	dt := rl.GetFrameTime()

	vel := rl.Vector3Scale(dir, data.Speed*dt)

	transform.Pos = rl.Vector3Add(transform.Pos, vel)
}

func (pc *PlayerController) Update(scene *kanye.BaseScene) {
	pc.IterateEntities(scene, func(player *kanye.Entity) {
		transform := (*player.GetComponent(components.TRANSFORM_ID)).(*components.Transform)
		data := (*player.GetComponent(components.PLAYER_DATA_ID)).(*components.PlayerData)

		dir := handleControls(scene.Camera)
		handlePhysics(transform, data, dir)
	}, components.PLAYER_DATA_ID)
}

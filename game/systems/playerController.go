package systems

import (
	"main/game/components"
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerController struct {
	kanye.BaseSystem
}

func handleControls() rl.Vector3 {
	dir := rl.Vector3Zero()

	if rl.IsKeyDown(rl.KeyW) {
		dir.Z += -1
	}

	if rl.IsKeyDown(rl.KeyS) {
		dir.Z += 1
	}

	if rl.IsKeyDown(rl.KeyA) {
		dir.X += -1
	}

	if rl.IsKeyDown(rl.KeyD) {
		dir.X += 1
	}

  dir = rl.Vector3Normalize(dir)

  return dir
}

func handlePhysics(transform *components.Transform, data *components.PlayerData, dir rl.Vector3) {
	dt := rl.GetFrameTime()

  vel := rl.Vector3Scale(dir, data.Speed * dt)

  transform.Pos = rl.Vector3Add(transform.Pos, vel)
}

func (pc *PlayerController) Update(scene *kanye.BaseScene) {
	pc.IterateEntities(scene, func (player *kanye.Entity) {
    transform := (*player.GetComponent(components.TRANSFORM_ID)).(*components.Transform)
    data := (*player.GetComponent(components.PLAYER_DATA_ID)).(*components.PlayerData)

    dir := handleControls()
    handlePhysics(transform, data, dir)
	}, components.PLAYER_DATA_ID)
}

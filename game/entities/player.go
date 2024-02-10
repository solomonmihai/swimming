package entities

import (
	"main/game/components"
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewPlayer() *kanye.Entity {
	player := kanye.NewEntity()

	player.AddComponent(
		components.NewTransform(rl.Vector3Zero(), rl.NewVector3(100, 100, 100)),
	)

	return player
}

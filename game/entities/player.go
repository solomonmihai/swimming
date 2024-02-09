package entities

import (
	"main/game/components"
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewPlayer() *kanye.Entity {
	player := kanye.NewEntity()

	player.AddComponent(
		components.NewTransform(rl.Vector2Zero(), rl.NewVector2(100, 100)),
	)

	player.AddComponent(&components.SquareRender{
		Color: rl.Red,
	})

	return player
}

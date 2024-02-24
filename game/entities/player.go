package entities

import (
	"main/game"
	"main/game/components"
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewPlayer() *kanye.Entity {
	player := kanye.NewEntity()

	transform := components.DefaultTransform()

	model := components.NewModel(game.Game.Assets.Model("cube"))
	model.Color = &rl.White

	player.AddComponent(
		transform,
		model,
		components.NewPlayer(),
	)

	return player
}

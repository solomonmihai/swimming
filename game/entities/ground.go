package entities

import (
	"main/game"
	"main/game/components"
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewGround(pos rl.Vector3) *kanye.Entity {
  ground := kanye.NewEntity()

  transform := components.DefaultTransform()
  transform.Pos = pos

  model := components.NewModel(game.Game.Assets.Model("cube"))
  model.Color = &rl.Gray

  ground.AddComponent(
    transform,
    model,
  )

  return ground
}

package entities

import (
	"main/game/components"
	"main/kanye"
)

func NewBox(transform *components.Transform, model *components.Model) *kanye.Entity {
  ground := kanye.NewEntity()

  ground.AddComponent(
    transform,
    model,
  )

  return ground
}

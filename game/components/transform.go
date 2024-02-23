package components

import (
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Transform struct {
	Pos, Scale rl.Vector3
	Rotation float32
}

func (transform Transform) Id() kanye.ComponentId {
	return TRANSFORM_ID
}

func DefaultTransform() *Transform {
	return &Transform{
		Pos:   rl.Vector3Zero(),
		Scale: rl.Vector3One(),
		Rotation: 0,
	}
}

func NewTransform(pos, scale rl.Vector3) *Transform {
	return &Transform{
		Pos:   pos,
		Scale: scale,
	}
}

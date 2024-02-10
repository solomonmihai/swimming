package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Transform struct {
	Pos, Scale rl.Vector3
}

func DefaultTransform() *Transform {
	return &Transform{
		Pos:   rl.Vector3Zero(),
		Scale: rl.Vector3One(),
	}
}

func NewTransform(pos, scale rl.Vector3) *Transform {
	return &Transform{
		Pos:   pos,
		Scale: scale,
	}
}

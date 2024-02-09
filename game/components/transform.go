package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Transform struct {
	Pos, Scale rl.Vector2
}

func DefaultTransform() *Transform {
	return &Transform{
		Pos:   rl.Vector2Zero(),
		Scale: rl.Vector2One(),
	}
}

func NewTransform(pos, scale rl.Vector2) *Transform {
	return &Transform{
		Pos:   pos,
		Scale: scale,
	}
}

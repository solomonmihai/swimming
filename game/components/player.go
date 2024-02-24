package components

import (
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerData struct {
	Gravity, Speed, JumpForce float32
	Acceleration rl.Vector3
}

func (player PlayerData) Id() kanye.ComponentId {
	return PLAYER_DATA_ID
}

func NewPlayer() *PlayerData {
	return &PlayerData{
		Gravity: 0,
		Speed: 20,
		JumpForce: 0,
		Acceleration: rl.Vector3Zero(),
	}
}


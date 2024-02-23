package components

import "main/kanye"

type Player struct {
	Gravity, Speed, JumpForce float32
}

func (player Player) Id() kanye.ComponentId {
	return PLAYER_ID
}

func NewPlayer() *Player {
	return &Player{
		Gravity: 0,
		Speed: 0,
		JumpForce: 0,
	}
}


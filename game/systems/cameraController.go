package systems

import (
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CameraController struct {}

func NewCameraController(scene *kanye.BaseScene) *CameraController {
	cc := &CameraController{}

	cam := rl.NewCamera3D(rl.NewVector3(50, 50, 50), rl.NewVector3(0, 10, 0), rl.NewVector3(0, 1, 0), 45, rl.CameraPerspective)

	scene.Camera = &cam

	return cc
}

func (cc *CameraController) Update(scene *kanye.BaseScene) {
	rl.UpdateCamera(scene.Camera, rl.CameraFirstPerson)	
}

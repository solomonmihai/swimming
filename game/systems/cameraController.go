package systems

import (
	"main/game/components"
	"main/kanye"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CameraController struct {
	kanye.BaseSystem

	camDist, xRot, yRot float32
	sensitivity         float32
}

func NewCameraController(scene *kanye.BaseScene) *CameraController {
	cc := &CameraController{
		camDist:     50,
		xRot:        -rl.Pi / 2,
		yRot:        0,
		sensitivity: 0.1,
	}

	pos := rl.NewVector3(0, cc.camDist, cc.camDist)

	cam := rl.NewCamera3D(pos, rl.Vector3Zero(), kanye.VEC_UP, 45, rl.CameraPerspective)
	scene.Camera = &cam

	rl.DisableCursor()

	return cc
}

func (cc *CameraController) Update(scene *kanye.BaseScene) {
	dt := rl.GetFrameTime()

	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	// Get mouse movement
	mouseX := float32(rl.GetMouseX())
	mouseY := float32(rl.GetMouseY())

	cc.xRot -= (mouseX - screenWidth/2) * cc.sensitivity * dt
	cc.yRot -= (mouseY - screenHeight/2) * cc.sensitivity * dt

	cc.yRot = rl.Clamp(cc.yRot, rl.Pi/9, rl.Pi/3)

	rl.SetMousePosition(int(screenWidth/2), int(screenHeight/2))

	cc.IterateEntities(scene, func(ent *kanye.Entity) {
		transform := (*ent.GetComponent(components.TRANSFORM_ID)).(*components.Transform)

		x := float32(math.Cos(float64(cc.xRot))) * cc.camDist
		y := float32(math.Sin(float64(cc.yRot))) * cc.camDist
		z := float32(math.Sin(float64(cc.xRot))) * cc.camDist * -1

		scene.Camera.Position.X = transform.Pos.X + x
		scene.Camera.Position.Y = transform.Pos.Y + y
		scene.Camera.Position.Z = transform.Pos.Z + z

		scene.Camera.Target = transform.Pos
	}, components.PLAYER_DATA_ID)
}

package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type Camera struct {
	Position, Front, Up, Right, WorldUp mgl32.Vec3
	Yaw, Pitch                          float32
	MovementSpeed, MouseSensitivity     float32
}

func NewCamera(position, up mgl32.Vec3, yaw, pitch float32) *Camera {
	front := mgl32.Vec3{0.0, 0.0, -1.0}
	camera := &Camera{
		Position:         position,
		Front:            front,
		Up:               up,
		Right:            up.Cross(front).Normalize(),
		WorldUp:          up,
		Yaw:              yaw,
		Pitch:            pitch,
		MovementSpeed:    2.5,
		MouseSensitivity: 0.1,
	}
	camera.updateCameraVectors()
	return camera
}

func (camera *Camera) GetViewMatrix() mgl32.Mat4 {
	return mgl32.LookAtV(camera.Position, camera.Position.Add(camera.Front), camera.Up)
}

func (camera *Camera) updateCameraVectors() {
	front := mgl32.Vec3{
		float32(math.Cos(float64(mgl32.DegToRad(camera.Yaw))) * math.Cos(float64(mgl32.DegToRad(camera.Pitch)))),
		float32(math.Sin(float64(mgl32.DegToRad(camera.Pitch)))),
		float32(math.Sin(float64(mgl32.DegToRad(camera.Yaw))) * math.Cos(float64(mgl32.DegToRad(camera.Pitch)))),
	}
	camera.Front = front.Normalize()
	camera.Right = camera.Front.Cross(camera.WorldUp).Normalize()
	camera.Up = camera.Right.Cross(camera.Front).Normalize()
}

func (camera *Camera) ProcessKeyboard(direction string, deltaTime float32) {
	velocity := camera.MovementSpeed * deltaTime
	if direction == "FORWARD" {
		camera.Position = camera.Position.Add(camera.Front.Mul(velocity))
	}
	if direction == "BACKWARD" {
		camera.Position = camera.Position.Sub(camera.Front.Mul(velocity))
	}
	if direction == "LEFT" {
		camera.Position = camera.Position.Sub(camera.Right.Mul(velocity))
	}
	if direction == "RIGHT" {
		camera.Position = camera.Position.Add(camera.Right.Mul(velocity))
	}
}

func (camera *Camera) ProcessMouseMovement(xoffset, yoffset float32) {
	xoffset *= camera.MouseSensitivity
	yoffset *= camera.MouseSensitivity

	camera.Yaw += xoffset
	camera.Pitch += yoffset

	if camera.Pitch > 89.0 {
		camera.Pitch = 89.0
	}
	if camera.Pitch < -89.0 {
		camera.Pitch = -89.0
	}

	camera.updateCameraVectors()
}

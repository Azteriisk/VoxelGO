package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Scene struct {
	Lights []*Light
	Cube   *Cube
	Camera *Camera
}

func NewScene() *Scene {
	// Initialize lights
	lights := []*Light{
		NewLight(mgl32.Vec3{1.2, 1.0, 2.0}, mgl32.Vec3{1.0, 1.0, 1.0}, 1.0),
		NewLight(mgl32.Vec3{-1.2, 1.0, 2.0}, mgl32.Vec3{1.0, 0.0, 0.0}, 0.5),
	}

	// Initialize the cube with specific scale and color
	cube := NewCube(1.0, mgl32.Vec3{0.5, 0.7, 0.3}) // Scale and Color defined here

	// Initialize the camera with default properties
	camera := NewCamera(mgl32.Vec3{0.0, 0.0, 3.0}, mgl32.Vec3{0.0, 1.0, 0.0}, -90.0, 0.0)

	return &Scene{
		Lights: lights,
		Cube:   cube,
		Camera: camera,
	}
}

func (scene *Scene) SetUpShaderUniforms(shaderProgram uint32) {
	// Set the light position
	lightPos := mgl32.Vec3{1.2, 1.0, 2.0}
	lightPosUniform := gl.GetUniformLocation(shaderProgram, gl.Str("lightPos\x00"))
	gl.Uniform3fv(lightPosUniform, 1, &lightPos[0])

	// Set the view position (camera position)
	viewPosUniform := gl.GetUniformLocation(shaderProgram, gl.Str("viewPos\x00"))
	gl.Uniform3fv(viewPosUniform, 1, &scene.Camera.Position[0])

	// Set the light color
	lightColor := mgl32.Vec3{1.0, 1.0, 1.0} // white light
	lightColorUniform := gl.GetUniformLocation(shaderProgram, gl.Str("lightColor\x00"))
	gl.Uniform3fv(lightColorUniform, 1, &lightColor[0])

	// Set the object color (cube color)
	objectColorUniform := gl.GetUniformLocation(shaderProgram, gl.Str("objectColor\x00"))
	gl.Uniform3fv(objectColorUniform, 1, &scene.Cube.Color[0])

	// Set the projection matrix once
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), 800.0/600.0, 0.1, 100.0)
	projectionUniform := gl.GetUniformLocation(shaderProgram, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])
}

func (scene *Scene) UpdateCameraView(shaderProgram uint32) {
	view := scene.Camera.GetViewMatrix()
	viewUniform := gl.GetUniformLocation(shaderProgram, gl.Str("view\x00"))
	gl.UniformMatrix4fv(viewUniform, 1, false, &view[0])
}

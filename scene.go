package main

import (
	"fmt"
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
	// Set the number of lights
	numLightsUniform := gl.GetUniformLocation(shaderProgram, gl.Str("numLights\x00"))
	gl.Uniform1i(numLightsUniform, int32(len(scene.Lights)))

	for i, light := range scene.Lights {
		// Set the light position
		lightPosUniform := gl.GetUniformLocation(shaderProgram, gl.Str(fmt.Sprintf("lights[%d].position\x00", i)))
		gl.Uniform3fv(lightPosUniform, 1, &light.Position[0])

		// Set the light color
		lightColorUniform := gl.GetUniformLocation(shaderProgram, gl.Str(fmt.Sprintf("lights[%d].color\x00", i)))
		gl.Uniform3fv(lightColorUniform, 1, &light.Color[0])

		// Set the light intensity
		lightIntensityUniform := gl.GetUniformLocation(shaderProgram, gl.Str(fmt.Sprintf("lights[%d].intensity\x00", i)))
		gl.Uniform1f(lightIntensityUniform, light.Intensity)
	}

	// Set the view position (camera position)
	viewPosUniform := gl.GetUniformLocation(shaderProgram, gl.Str("viewPos\x00"))
	gl.Uniform3fv(viewPosUniform, 1, &scene.Camera.Position[0])

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

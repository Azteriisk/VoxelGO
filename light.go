package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"strconv"
)

type Light struct {
	Position  mgl32.Vec3
	Color     mgl32.Vec3
	Intensity float32
}

func NewLight(position, color mgl32.Vec3, intensity float32) *Light {
	return &Light{
		Position:  position,
		Color:     color,
		Intensity: intensity,
	}
}

func (light *Light) SetUniforms(shaderProgram uint32, index int) {
	// Build uniform name strings for this light
	lightPosStr := "lights[" + strconv.Itoa(index) + "].position\x00"
	lightColorStr := "lights[" + strconv.Itoa(index) + "].color\x00"
	lightIntensityStr := "lights[" + strconv.Itoa(index) + "].intensity\x00"

	// Set the uniform values in the shader
	gl.Uniform3fv(gl.GetUniformLocation(shaderProgram, gl.Str(lightPosStr)), 1, &light.Position[0])
	gl.Uniform3fv(gl.GetUniformLocation(shaderProgram, gl.Str(lightColorStr)), 1, &light.Color[0])
	gl.Uniform1f(gl.GetUniformLocation(shaderProgram, gl.Str(lightIntensityStr)), light.Intensity)
}

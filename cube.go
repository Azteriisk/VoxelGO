package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Cube struct {
	VAO   uint32
	VBO   uint32
	Color mgl32.Vec3
}

func NewCube(scale float32, color mgl32.Vec3) *Cube {
	var VAO, VBO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.GenBuffers(1, &VBO)

	gl.BindVertexArray(VAO)

	vertices := []float32{
		// positions          // normals
		// Front face
		-0.5 * scale, -0.5 * scale, 0.5 * scale, 0.0, 0.0, 1.0,
		0.5 * scale, -0.5 * scale, 0.5 * scale, 0.0, 0.0, 1.0,
		0.5 * scale, 0.5 * scale, 0.5 * scale, 0.0, 0.0, 1.0,
		0.5 * scale, 0.5 * scale, 0.5 * scale, 0.0, 0.0, 1.0,
		-0.5 * scale, 0.5 * scale, 0.5 * scale, 0.0, 0.0, 1.0,
		-0.5 * scale, -0.5 * scale, 0.5 * scale, 0.0, 0.0, 1.0,

		// Back face
		-0.5 * scale, -0.5 * scale, -0.5 * scale, 0.0, 0.0, -1.0,
		0.5 * scale, -0.5 * scale, -0.5 * scale, 0.0, 0.0, -1.0,
		0.5 * scale, 0.5 * scale, -0.5 * scale, 0.0, 0.0, -1.0,
		0.5 * scale, 0.5 * scale, -0.5 * scale, 0.0, 0.0, -1.0,
		-0.5 * scale, 0.5 * scale, -0.5 * scale, 0.0, 0.0, -1.0,
		-0.5 * scale, -0.5 * scale, -0.5 * scale, 0.0, 0.0, -1.0,

		// Left face
		-0.5 * scale, 0.5 * scale, 0.5 * scale, -1.0, 0.0, 0.0,
		-0.5 * scale, 0.5 * scale, -0.5 * scale, -1.0, 0.0, 0.0,
		-0.5 * scale, -0.5 * scale, -0.5 * scale, -1.0, 0.0, 0.0,
		-0.5 * scale, -0.5 * scale, -0.5 * scale, -1.0, 0.0, 0.0,
		-0.5 * scale, -0.5 * scale, 0.5 * scale, -1.0, 0.0, 0.0,
		-0.5 * scale, 0.5 * scale, 0.5 * scale, -1.0, 0.0, 0.0,

		// Right face
		0.5 * scale, 0.5 * scale, 0.5 * scale, 1.0, 0.0, 0.0,
		0.5 * scale, 0.5 * scale, -0.5 * scale, 1.0, 0.0, 0.0,
		0.5 * scale, -0.5 * scale, -0.5 * scale, 1.0, 0.0, 0.0,
		0.5 * scale, -0.5 * scale, -0.5 * scale, 1.0, 0.0, 0.0,
		0.5 * scale, -0.5 * scale, 0.5 * scale, 1.0, 0.0, 0.0,
		0.5 * scale, 0.5 * scale, 0.5 * scale, 1.0, 0.0, 0.0,

		// Top face
		-0.5 * scale, 0.5 * scale, -0.5 * scale, 0.0, 1.0, 0.0,
		0.5 * scale, 0.5 * scale, -0.5 * scale, 0.0, 1.0, 0.0,
		0.5 * scale, 0.5 * scale, 0.5 * scale, 0.0, 1.0, 0.0,
		0.5 * scale, 0.5 * scale, 0.5 * scale, 0.0, 1.0, 0.0,
		-0.5 * scale, 0.5 * scale, 0.5 * scale, 0.0, 1.0, 0.0,
		-0.5 * scale, 0.5 * scale, -0.5 * scale, 0.0, 1.0, 0.0,

		// Bottom face
		-0.5 * scale, -0.5 * scale, -0.5 * scale, 0.0, -1.0, 0.0,
		0.5 * scale, -0.5 * scale, -0.5 * scale, 0.0, -1.0, 0.0,
		0.5 * scale, -0.5 * scale, 0.5 * scale, 0.0, -1.0, 0.0,
		0.5 * scale, -0.5 * scale, 0.5 * scale, 0.0, -1.0, 0.0,
		-0.5 * scale, -0.5 * scale, 0.5 * scale, 0.0, -1.0, 0.0,
		-0.5 * scale, -0.5 * scale, -0.5 * scale, 0.0, -1.0, 0.0,
	}

	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, gl.Ptr(uintptr(0)))
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6*4, gl.Ptr(uintptr(3*4)))
	gl.EnableVertexAttribArray(1)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return &Cube{
		VAO:   VAO,
		VBO:   VBO,
		Color: color,
	}
}

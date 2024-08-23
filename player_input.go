package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type PlayerInput struct {
	Window       *glfw.Window
	Camera       *Camera
	LastX, LastY float64
	FirstMouse   bool
}

// NewPlayerInput Constructor to create a new PlayerInput instance
func NewPlayerInput(window *glfw.Window, camera *Camera) *PlayerInput {
	return &PlayerInput{
		Window:     window,
		Camera:     camera,
		LastX:      400, // Initial mouse position (center of the window)
		LastY:      300,
		FirstMouse: true,
	}
}

// ProcessInput handles keyboard input for camera movement
func (input *PlayerInput) ProcessInput(deltaTime float32) {
	if input.Window.GetKey(glfw.KeyW) == glfw.Press {
		input.Camera.ProcessKeyboard("FORWARD", deltaTime)
	}
	if input.Window.GetKey(glfw.KeyS) == glfw.Press {
		input.Camera.ProcessKeyboard("BACKWARD", deltaTime)
	}
	if input.Window.GetKey(glfw.KeyA) == glfw.Press {
		input.Camera.ProcessKeyboard("LEFT", deltaTime)
	}
	if input.Window.GetKey(glfw.KeyD) == glfw.Press {
		input.Camera.ProcessKeyboard("RIGHT", deltaTime)
	}
}

// MouseCallback handles mouse movement and updates camera orientation
func (input *PlayerInput) MouseCallback(xpos, ypos float64) {
	if input.FirstMouse {
		input.LastX = xpos
		input.LastY = ypos
		input.FirstMouse = false
	}

	xoffset := float32(xpos - input.LastX)
	yoffset := float32(input.LastY - ypos) // Reversed since y-coordinates range from bottom to top

	input.LastX = xpos
	input.LastY = ypos

	input.Camera.ProcessMouseMovement(xoffset, yoffset)
}

// SetupCallbacks sets the necessary callbacks to capture input
func (input *PlayerInput) SetupCallbacks() {
	input.Window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled) // Hide the cursor and capture it
	input.Window.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		input.MouseCallback(xpos, ypos)
	})
}

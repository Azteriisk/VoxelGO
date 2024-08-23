package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"log"
	"os"
	"runtime"
	"strings"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalf("failed to initialize glfw: %v", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	window, err := glfw.CreateWindow(800, 600, "Voxel Engine", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	gl.Viewport(0, 0, 800, 600)
	gl.ClearColor(0.2, 0.3, 0.3, 1.0)
	gl.Enable(gl.DEPTH_TEST)

	// Create the scene
	scene := NewScene()

	// Initialize the shader program
	shaderProgram := CreateShaderProgram("vertex_shader.glsl", "fragment_shader.glsl")
	gl.UseProgram(shaderProgram)

	// Set up scene uniforms (lights, cube color, projection)
	scene.SetUpShaderUniforms(shaderProgram)

	// Set up player input for the camera
	playerInput := NewPlayerInput(window, scene.Camera)
	playerInput.SetupCallbacks()

	// Model matrix
	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(shaderProgram, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	// Main render loop
	for !window.ShouldClose() {
		deltaTime := float32(0.016) // Replace with actual frame time calculation

		// Process input
		playerInput.ProcessInput(deltaTime)

		// Clear the screen and draw
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Update the view matrix from the camera
		scene.UpdateCameraView(shaderProgram)

		// Bind the VAO and draw the cube
		gl.BindVertexArray(scene.Cube.VAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 36)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

// CreateShaderProgram loads, compiles, and links vertex and fragment shaders
func CreateShaderProgram(vertexPath, fragmentPath string) uint32 {
	vertexShaderSource, err := os.ReadFile(vertexPath)
	if err != nil {
		log.Fatalf("failed to load vertex shader: %v", err)
	}
	fragmentShaderSource, err := os.ReadFile(fragmentPath)
	if err != nil {
		log.Fatalf("failed to load fragment shader: %v", err)
	}

	vertexShader := compileShader(string(vertexShaderSource)+"\x00", gl.VERTEX_SHADER)
	fragmentShader := compileShader(string(fragmentShaderSource)+"\x00", gl.FRAGMENT_SHADER)

	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)

	var success int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)
		logInfo := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(logInfo))
		log.Fatalf("failed to link shader program: %v", logInfo)
	}

	// Clean up shaders (no longer needed after linking into a program)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return shaderProgram
}

// compileShader compiles individual shaders (vertex or fragment)
func compileShader(source string, shaderType uint32) uint32 {
	shader := gl.CreateShader(shaderType)
	csource, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csource, nil)
	free()
	gl.CompileShader(shader)

	var success int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		logInfo := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logInfo))
		log.Fatalf("failed to compile shader: %v", logInfo)
	}

	return shader
}

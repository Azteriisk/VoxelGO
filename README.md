### **SO FAR...**
- This uses OpenGL to create a window and generate a single voxel
- You can give the voxel scale, color, translate, and rotate properties that override the defaults.
- There's a camera defined that acts as the player with WASD movement and mouse-move look functionality (TODO: Space Bar + L-Ctrl)
- The player then has properties for speed and look sensitivity that can be altered so far. (TODO: FOV, Depth of Field)
- Support for multiple lights and blending between them along with some other basic lighting/render features
- Shader files are automatically embedded into the final exe for portable execution. (TODO: Find a way to not trigger Windows Defender when I use -ldflags="-H windowsgui"

### **Next Main Features:**
- Multiple voxels
- Simplified Game Config file to centralize certain settings about items in the scene.go file.
- Voxel Interaction (Delete + Place + Push + Pull)

package main

import (
	"github.com/vfrazao-ns1/raytracing1weekend/objects"
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

type config struct {
	FileName        string          // Name of file to save to render
	ImgWidth        int             // Resolution width
	Aspect          float64         // Aspect ratio float e.g., 16:9 equals 1.7777777
	SamplesPerPixel int             // How many rays to simulate hitting a given pixel (higher is better quality)
	MaxDepth        int             // Max distance a ray will fly
	Camera          cameraConfig    // Camera config
	Animation       animationConfig // Whether to make an animation
}

type cameraConfig struct {
	LookFrom  vec3.Point // Initial camera position
	LookAt    vec3.Point // Point at which the camera is initially looking at
	Vup       vec3.Vec3  // ViewUp vector, where is up
	VFOV      float64    // Vertical field of view
	Aperture  float64    // Camera aperture
	FocusDist float64    // Camera focus distance
}

type worldConfig struct {
	Random bool              // Whether to generate a random scene, overides the Static attribute
	Static objects.Hittables // List of "Hittable" shapes
}

type animationConfig struct {
	Enabled  bool // Whether to run in animation mode
	Fps      int  // Frames per second
	Duration int  // How long to animate for in seconds
}

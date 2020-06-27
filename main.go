package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/vfrazao-ns1/raytracing1weekend/renderer"

	"github.com/vfrazao-ns1/raytracing1weekend/camera"
	"github.com/vfrazao-ns1/raytracing1weekend/objects"
	"github.com/vfrazao-ns1/raytracing1weekend/ray"
	"github.com/vfrazao-ns1/raytracing1weekend/utils"
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

const (
	maxColor = 255
)

func main() {
	fileName := os.Args[1]

	imgWidth := 380
	var err error
	if len(os.Args) == 3 {
		imgWidth, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse image size quitting! - %v\n", err)
			os.Exit(1)
		}
	}

	aspect := 16.0 / 9.0
	imgHeight := int(float64(imgWidth) / aspect)
	samplesPerPixel := 100
	maxDepth := 50

	numPixels := (imgHeight * imgWidth)

	lookfrom := vec3.Point{X: 13, Y: 2, Z: 3}
	lookat := vec3.Point{X: 0, Y: 0, Z: 0}
	vup := vec3.Vec3{X: 0, Y: 1, Z: 0}
	distToFocus := 10.0
	aperture := 0.1
	cam := camera.InitCamera(lookfrom, lookat, vup, 20, float64(imgWidth)/float64(imgHeight), aperture, distToFocus)

	world := RandomWorld()
	numWorkers := runtime.NumCPU()
	jobs := make(chan job, numWorkers*10)
	results := make(chan renderer.Pixel, numPixels)
	start := time.Now()

	s := workerState{
		jobs:     jobs,
		results:  results,
		height:   imgHeight,
		width:    imgWidth,
		spp:      samplesPerPixel,
		world:    world,
		maxDepth: maxDepth,
		cam:      cam,
	}

	go fillJobsQueue(imgHeight, imgWidth, jobs)
	for i := 0; i < numWorkers; i++ {
		go worker(s)
	}

	pixels := make([]renderer.Pixel, numPixels)
	for i := 0; i < numPixels; i++ {
		if i%1000 == 0 && i > 0 {
			progress(i, numPixels, start)
		}
		pixels[i] = <-results
	}
	progress(numPixels, numPixels, start)
	close(jobs)

	pngRenderer := renderer.PNGRenderer{
		ImageWidth:      imgWidth,
		ImageHeight:     imgHeight,
		ImagePixels:     pixels,
		SamplesPerPixel: samplesPerPixel,
	}
	pngRenderer.Render(fileName)
	fmt.Fprintf(os.Stderr, "\nDone\n")
}

// RayColor returns the ray color
func RayColor(r ray.Ray, world objects.Hittable, depth int) vec3.Color {
	hitRec := new(objects.HitRecord)

	if depth <= 0 {
		return vec3.Color{X: 0, Y: 0, Z: 0}
	}

	tmin := 0.001
	tmax := math.Inf(1)
	if world.Hit(r, tmin, tmax, hitRec) {
		scattered := new(ray.Ray)
		attenuation := new(vec3.Color)
		if hitRec.Material.Scatter(r, *hitRec, attenuation, scattered) {
			return attenuation.Mul(RayColor(*scattered, world, depth-1))
		}
		return vec3.Color{X: 0, Y: 0, Z: 0}
	}

	// If no hits then the color == background
	unitDirection := r.Direction.Unit()
	t := 0.5 * (unitDirection.Y + 1.0)
	return vec3.Color{X: 1, Y: 1, Z: 1}.ScalarMul(1 - t).Add(vec3.Color{X: 0.5, Y: 0.7, Z: 1}.ScalarMul(t))
}

type workerState struct {
	jobs     <-chan job
	results  chan<- renderer.Pixel
	height   int
	width    int
	spp      int // samples per pixel
	world    objects.HittableList
	maxDepth int
	cam      *camera.Camera
}

type job struct {
	i int
	j int
}

func fillJobsQueue(height, width int, jobs chan<- job) {
	for j := (height - 1); j >= 0; j-- {
		for i := 0; i < width; i++ {
			jobs <- job{i: i, j: j}
		}
	}
}

func worker(state workerState) {
	for job := range state.jobs {
		pixel := renderer.Pixel{
			Color:    vec3.Color{X: 0, Y: 0, Z: 0},
			Position: vec3.Point{X: float64(job.i), Y: float64(state.height - 1 - job.j), Z: 0},
		}
		for s := 0; s < state.spp; s++ {
			u := (float64(job.i) + utils.RandomDouble()) / float64(state.width-1)
			v := (float64(job.j) + utils.RandomDouble()) / float64(state.height-1)
			ray := state.cam.GetRay(u, v)
			pixel.Color = pixel.Color.Add(RayColor(ray, state.world, state.maxDepth))
		}
		state.results <- pixel
	}
}

func progress(done, total int, start time.Time) {
	barSize := 70
	pctComplete := float64(done) / float64(total)
	doneBars := int(pctComplete * float64(barSize))
	elapsed := time.Since(start).Seconds()
	rate := float64(done) / elapsed
	eta := float64(total-done) / rate

	fmt.Fprintf(os.Stderr, "\r[")
	for i := 0; i < doneBars; i++ {
		fmt.Fprintf(os.Stderr, "#")
	}
	for i := 0; i < (barSize - doneBars); i++ {
		fmt.Fprintf(os.Stderr, " ")
	}
	fmt.Fprintf(os.Stderr, "] (%.2f%%) Rate: %.2f - Elapsed: %6d - ETA: %6ds", pctComplete*100, rate, int(elapsed), int(eta))

	for i := 0; i < 5; i++ {
		fmt.Fprintf(os.Stderr, " ")
	}
}

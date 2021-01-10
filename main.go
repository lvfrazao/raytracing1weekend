package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
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

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
var configFile = flag.String("config", "config.json", "Location of config file")
var worldFile = flag.String("world", "world.json", "Location of world file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// Load tracer configuration
	var tracerConfig config
	c, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to open config: %s\n", err))
	}
	if err = json.Unmarshal(c, &tracerConfig); err != nil {
		log.Fatal(fmt.Sprintf("Unable to decode config: %s\n+", err))
	}

	// Load world configuration
	var worldConf worldConfig
	w, err := ioutil.ReadFile(*worldFile)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to open world file: %s\n", err))
	}
	if err = json.Unmarshal(w, &worldConf); err != nil {
		log.Fatal(fmt.Sprintf("Unable to decode world config: %s\n+", err))
	}

	imgHeight := int(float64(tracerConfig.ImgWidth) / tracerConfig.Aspect)

	cam := camera.InitCamera(tracerConfig.Camera.LookFrom, tracerConfig.Camera.LookAt, tracerConfig.Camera.Vup, tracerConfig.Camera.VFOV, float64(tracerConfig.ImgWidth)/float64(imgHeight), tracerConfig.Camera.Aperture, tracerConfig.Camera.FocusDist)

	var world objects.HittableList
	if worldConf.Random == true {
		world = RandomWorld()
	} else {
		world = worldFromConfig(worldConf)
	}

	if tracerConfig.Animation.Enabled {
		framesPerSecond := tracerConfig.Animation.Fps
		numFrames := framesPerSecond * tracerConfig.Animation.Duration
		fileExt := filepath.Ext(tracerConfig.FileName)
		baseFileName := tracerConfig.FileName[:len(tracerConfig.FileName)-len(fileExt)]
		// radius of our circle is the distance from the origin in the XZ plane
		r := math.Sqrt(math.Pow(tracerConfig.Camera.LookFrom.X, 2) + math.Pow(tracerConfig.Camera.LookFrom.Z, 2))
		for i := 0; i < numFrames+1; i++ {
			// Move in circle of radius r
			// x**2 + z**2 = r**2
			// z = sqrt(r**2 - x**2)
			// x max is r
			// x min is -r
			tracerConfig.Camera.LookFrom.X = float64(int((r*2/float64(numFrames))*float64(i*2)*1000000)%int((2*r+0.0001)*1000000))/1000000 - r
			tracerConfig.Camera.LookFrom.Z = math.Sqrt(math.Pow(r, 2) - math.Pow(tracerConfig.Camera.LookFrom.X, 2))
			if float64(i) > float64((numFrames+1)/2.0) {
				tracerConfig.Camera.LookFrom.X = tracerConfig.Camera.LookFrom.X * -1
				tracerConfig.Camera.LookFrom.Z = tracerConfig.Camera.LookFrom.Z * -1
			}
			cam = camera.InitCamera(tracerConfig.Camera.LookFrom, tracerConfig.Camera.LookAt, tracerConfig.Camera.Vup, tracerConfig.Camera.VFOV, float64(tracerConfig.ImgWidth)/float64(imgHeight), tracerConfig.Camera.Aperture, tracerConfig.Camera.FocusDist)
			// Format specifier hardcoded to number of zero padding, might want to do this more dynamically some other time
			tracerConfig.FileName = fmt.Sprintf("%s%05d%s", baseFileName, i, fileExt)
			renderFrame(tracerConfig, world, cam)
		}

	} else {
		renderFrame(tracerConfig, world, cam)
	}
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

func renderFrame(c config, world objects.HittableList, cam *camera.Camera) {
	imgHeight := int(float64(c.ImgWidth) / c.Aspect)
	numPixels := (imgHeight * c.ImgWidth)
	numWorkers := runtime.NumCPU()
	jobs := make(chan job, numWorkers*10)
	results := make(chan renderer.Pixel, numWorkers*10)
	start := time.Now()

	s := workerState{
		jobs:     jobs,
		results:  results,
		height:   imgHeight,
		width:    c.ImgWidth,
		spp:      c.SamplesPerPixel,
		world:    world,
		maxDepth: c.MaxDepth,
		cam:      cam,
	}
	go fillJobsQueue(imgHeight, c.ImgWidth, jobs)
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
		ImageWidth:      c.ImgWidth,
		ImageHeight:     imgHeight,
		ImagePixels:     pixels,
		SamplesPerPixel: c.SamplesPerPixel,
	}
	pngRenderer.Render(c.FileName)
	fmt.Fprintf(os.Stderr, "\n")
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
	fmt.Fprintf(os.Stderr, "] (%.2f%%) Rate: %.0f - Elapsed: %6.2f - ETA: %6ds", pctComplete*100, rate, elapsed, int(eta))

	for i := 0; i < 5; i++ {
		fmt.Fprintf(os.Stderr, " ")
	}
}

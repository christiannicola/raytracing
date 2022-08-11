package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	var (
		imageFile     *os.File
		cpuProfile    *os.File
		memProfile    *os.File
		imageFileName *string
		// Image
		imageAspectRatio     = 3.0 / 2.0
		imageWidth           = 1200
		imageHeight          = int(float64(imageWidth) / imageAspectRatio)
		imageSamplesPerPixel = 500
		imageMaxDepth        = 50
		// Camera
		lookFrom    = newVec3(13, 2, 3)
		lookAt      = newVec3(0, 0, 0)
		vup         = newVec3(0, 1, 0)
		distToFocus = 10.0
		aperture    = 0.1
		cam         = newCamera(lookFrom, lookAt, vup, 20, imageAspectRatio, aperture, distToFocus)
		// World
		world = randomScene()
		err   error
	)

	if cpuProfile, err = os.Create("cpu_profile"); err != nil {
		log.Fatalf("unable to open cpu profile: %v", err)
	}

	imageFileName = flag.String("f", "image.ppm", "output path of the resulting imageFile")
	flag.Parse()

	if imageFile, err = os.Create(*imageFileName); err != nil {
		log.Fatalf("unable to open image file: %v\n", err)
	}

	if err = pprof.StartCPUProfile(cpuProfile); err != nil {
		log.Fatalf("unable to start cpu profile: %v", err)
	}

	defer pprof.StopCPUProfile()

	if _, err = fmt.Fprintf(imageFile, "P3\n%d %d\n255\n", imageWidth, imageHeight); err != nil {
		log.Fatalf("unable to writer PPM header: %v", err)
	}

	for j := float64(imageHeight - 1); j >= 0; j-- {
		log.Printf("scanlines remaining: %d", int(j))
		for i := float64(0); i < float64(imageWidth); i++ {
			pixelColor := newVec3(0, 0, 0)

			for s := 0; s < imageSamplesPerPixel; s++ {
				u := (i + randomFloat64()) / float64(imageWidth-1)
				v := (j + randomFloat64()) / float64(imageHeight-1)
				r := cam.getRay(u, v)
				pixelColor = addVec3(rayColor(&r, &world, imageMaxDepth), pixelColor)
			}

			if err = writeColor(imageFile, pixelColor, imageSamplesPerPixel); err != nil {
				log.Fatalf("unable to write scan line to imageFile: %v\n", err)
			}
		}
	}

	if memProfile, err = os.Create("mem_profile"); err != nil {
		log.Fatalf("unable to open mem profile: %v", err)
	}

	if err = pprof.WriteHeapProfile(memProfile); err != nil {
		log.Fatalf("unable to write mem profile: %v", err)
	}
}

func randomScene() hittableList {
	world := newHittableList()

	groundMaterial := newLambertian(newVec3(0.5, 0.5, 0.5))

	world.add(newSphere(newVec3(0, -1000, 0), 1000, groundMaterial))

	for a := float64(-11); a < 11; a++ {
		for b := float64(-11); b < 11; b++ {
			chooseMat := randomFloat64()
			center := newVec3(a+0.9*randomFloat64(), 0.2, b+0.9*randomFloat64())

			if subtractVec3(center, newVec3(4, 0.2, 0)).length() > 0.9 {
				var sphereMaterial material

				if chooseMat < 0.8 {
					albedo := multiplyVec3(random(), random())
					sphereMaterial = newLambertian(albedo)
				} else if chooseMat < 0.95 {
					albedo := randomMinMax(0.5, 1)
					fuzz := randomFloat64MinMax(0, 0.5)
					sphereMaterial = newMetal(albedo, fuzz)
				} else {
					sphereMaterial = newDielectric(1.5)
				}

				world.add(newSphere(center, 0.2, sphereMaterial))
			}
		}
	}

	material1 := newDielectric(1.5)
	world.add(newSphere(newVec3(0, 1, 0), 1.0, material1))

	material2 := newLambertian(newVec3(0.4, 0.2, 0.1))
	world.add(newSphere(newVec3(-4, 1, 0), 1.0, material2))

	material3 := newMetal(newVec3(0.7, 0.6, 0.5), 0.0)
	world.add(newSphere(newVec3(4, 1, 0), 1.0, material3))

	return world
}

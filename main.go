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
		imageAspectRatio     = 16.0 / 9.0
		imageWidth           = 800
		imageHeight          = int(float64(imageWidth) / imageAspectRatio)
		imageSamplesPerPixel = 100
		imageMaxDepth        = 50
		// Camera
		cam = newCamera(newVec3(-2, 2, 1), newVec3(0, 0, -1), newVec3(0, 1, 0), 90, imageAspectRatio)
		// World
		world = newHittableList()
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

	materialGround := newLambertian(newVec3(0.8, 0.8, 0.0))
	materialCenter := newLambertian(newVec3(0.1, 0.2, 0.5))
	materialLeft := newDielectric(1.5)
	materialRight := newMetal(newVec3(0.8, 0.6, 0.2), 0.0)

	world.add(newSphere(newVec3(0.0, -100.5, -1.0), 100.0, materialGround))
	world.add(newSphere(newVec3(0.0, 0.0, -1.0), 0.5, materialCenter))
	world.add(newSphere(newVec3(-1.0, 0.0, -1.0), 0.5, materialLeft))
	world.add(newSphere(newVec3(-1.0, 0.0, -1.0), -0.4, materialLeft))
	world.add(newSphere(newVec3(1.0, 0.0, -1.0), 0.5, materialRight))

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

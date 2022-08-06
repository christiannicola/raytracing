package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		file     *os.File
		fileName *string
		// Image
		imageAspectRatio     = 16.0 / 9.0
		imageWidth           = 800
		imageHeight          = int(float64(imageWidth) / imageAspectRatio)
		imageSamplesPerPixel = 100
		// Camera
		cam = newCamera()
		// World
		world = newHittableList()
		err   error
	)

	fileName = flag.String("f", "image.ppm", "output path of the resulting file")
	flag.Parse()

	if file, err = os.Create(*fileName); err != nil {
		log.Fatalf("unable to open file: %v\n", err)
	}

	world.add(newSphere(newVec3(0, 0, -1), 0.5))
	world.add(newSphere(newVec3(0, -100.5, -1), 100))

	if _, err = fmt.Fprintf(file, "P3\n%d %d\n255\n", imageWidth, imageHeight); err != nil {
		log.Fatalf("unable to writer PPM header: %v", err)
	}

	for j := float64(imageHeight - 1); j >= 0; j-- {
		for i := float64(0); i < float64(imageWidth); i++ {
			pixelColor := newVec3(0, 0, 0)

			for s := 0; s < imageSamplesPerPixel; s++ {
				u := (i + randomFloat64()) / float64(imageWidth-1)
				v := (j + randomFloat64()) / float64(imageHeight-1)
				r := cam.getRay(u, v)
				pixelColor.add(rayColor(&r, &world))
			}

			if err = writeColor(file, pixelColor, imageSamplesPerPixel); err != nil {
				log.Fatalf("unable to write scan line to file: %v\n", err)
			}
		}
	}
}

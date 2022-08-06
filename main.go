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
		imageAspectRatio = 16.0 / 9.0
		imageWidth       = 800
		imageHeight      = int(float64(imageWidth) / imageAspectRatio)
		// Camera
		cameraViewportHeight = 2.0
		cameraViewportWidth  = imageAspectRatio * cameraViewportHeight
		cameraFocalLength    = 1.0
		cameraOrigin         = emptyVec3()
		cameraHorizontal     = newVec3(cameraViewportWidth, 0, 0)
		cameraVertical       = newVec3(0, cameraViewportHeight, 0)
		// World
		world = newHittableList()
		err   error
	)

	fileName = flag.String("f", "image.ppm", "output path of the resulting file")
	flag.Parse()

	if file, err = os.Create(*fileName); err != nil {
		log.Fatalf("unable to open file: %v\n", err)
	}

	// Position
	lowerLeftCorner := subtractVec3(subtractVec3(subtractVec3(cameraOrigin, divideVec3(cameraHorizontal, 2.0)), divideVec3(cameraVertical, 2.0)), newVec3(0, 0, cameraFocalLength))

	world.add(newSphere(newVec3(0, 0, -1), 0.5))
	world.add(newSphere(newVec3(0, -100.5, -1), 100))

	if _, err = fmt.Fprintf(file, "P3\n%d %d\n255\n", imageWidth, imageHeight); err != nil {
		log.Fatalf("unable to writer PPM header: %v", err)
	}

	for j := float64(imageHeight - 1); j >= 0; j-- {
		// log.Printf("scan lines remaining: %d", int(j))
		for i := float64(0); i < float64(imageWidth); i++ {
			u := i / float64(imageWidth-1)
			v := j / float64(imageHeight-1)

			direction := subtractVec3(addVec3(addVec3(lowerLeftCorner, multiplyVec3ByFactor(cameraHorizontal, u)), multiplyVec3ByFactor(cameraVertical, v)), cameraOrigin)

			r := newRay(cameraOrigin, direction)

			pixelColor := rayColor(&r, &world)

			if err = writeColor(file, pixelColor); err != nil {
				log.Fatalf("unable to write scan line to file: %v\n", err)
			}
		}
	}

	// log.Println("done")
}

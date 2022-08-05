package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	imageWidth  int = 256
	imageHeight     = 256
)

func main() {
	var (
		file     *os.File
		fileName *string
		err      error
	)

	fileName = flag.String("f", "image.ppm", "output path of the resulting file")
	flag.Parse()

	if file, err = os.Create(*fileName); err != nil {
		log.Fatalf("unable to open file: %v\n", err)
	}

	defer file.Close()

	if _, err = fmt.Fprintf(file, "P3\n%d %d\n255\n", imageWidth, imageHeight); err != nil {
		log.Fatalf("unable to writer PPM header: %v", err)
	}

	for j := float64(imageHeight - 1); j >= 0; j-- {
		log.Printf("scan lines remaining: %d", int(j))
		for i := float64(0); i < float64(imageWidth); i++ {
			pixelColor := newVec3(i/float64(imageWidth-1), j/float64(imageHeight-1), 0.25)

			if err = writeColor(file, pixelColor); err != nil {
				log.Fatalf("unable to write scan line to file: %v\n", err)
			}
		}
	}

	log.Println("done")
}

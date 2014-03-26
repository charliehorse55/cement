package main

import (
	"os"
	"fmt"
	"image"
	_ "image/png"
	"image/jpeg"
	"image/color"
)

func getImageRes(filename string) (int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()
	
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}
	
	return config.Width, config.Height, nil
}

func loadImage(filename string, n *image.NRGBA) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer r.Close()
	
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	
	expectSize := n.Bounds().Max
	imageSize := img.Bounds().Max
	if imageSize.X != expectSize.X || imageSize.Y != expectSize.Y {
		return fmt.Errorf("Image %s has different size", filename)
	}
	
	for j := 0; j < imageSize.Y; j++ {
		for k := 0; k < imageSize.X; k++ {
			n.Set(k, imageSize.Y - (j + 1), img.At(k,j))	
		}
	}
	return nil
}


type Pixel struct {
	R,G,B float32
}

func saveToJPEG(filename string, width, height int, data []Pixel) error {	
	output := image.NewNRGBA(image.Rectangle{Max: image.Point{X:width, Y:height}})
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			
			//prevent overflow for each channel when packed into 8 bits
			red := data[i*width +j].R*255
			if red > 255 {
				red = 255
			}
			green := data[i*width +j].G*255
			if green > 255 {
				green = 255
			}
			blue := data[i*width +j].B*255
			if blue > 255 {
				blue = 255
			}
	
			output.Set(j, height - (i+1), color.NRGBA{
					R:uint8(red),
					G:uint8(green),
					B:uint8(blue),
					A:uint8(255),
				})
		}
	}
		
	outfile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Failed to open output file: %s", filename)
	}
	defer outfile.Close()

	err = jpeg.Encode(outfile, output, nil)
	if err != nil {
		return fmt.Errorf("Failed to encode output file: %v", err)
	}
	
	return nil
}

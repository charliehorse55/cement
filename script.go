package main

import (
	"os"
	"fmt"
	"image"
	"io/ioutil"
	"encoding/json"
    "path/filepath"
	Ecem "github.com/charliehorse55/libcement"
)


type Lightvector struct {
	Filename string
	Intensity Ecem.RGB32f
}

type Config struct {
	filepath string
	basePath string
	Background string
	ResponseFunction string
	Vectors []Lightvector
}


func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	
	var result Config 
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	dir, err := filepath.Abs(filepath.Dir(path))
    if err != nil {
		return nil, err
    }
	result.basePath = dir
	result.filepath, err = filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Config)LoadPainting() (*Ecem.Painting, error) {
	//load the background
	path := filepath.Join(c.basePath, c.Background)
	
	width, height, err := getImageRes(path)
	if err != nil {
		return nil, err
	}
	
	nrgba := image.NewNRGBA(image.Rectangle{Max:image.Point{X:width, Y:height}})
	err = loadImage(path, nrgba)
	if err != nil {
		return nil, fmt.Errorf("Failed to load background %s:%v", path, err)
	}
	background := createTexture(nrgba)
	
	painting := Ecem.NewPainting(width, height, nil, background)
		
	for _,vector := range c.Vectors {
		path := filepath.Join(c.basePath, vector.Filename)
		err = loadImage(path, nrgba)
		if err != nil {
			return nil, fmt.Errorf("Failed to load image %s: %v", path, err)
		}
		painting.AddLightvector(createTexture(nrgba))
	}
	return painting, nil
}

func (c *Config)Save(path string) error {
	bytes, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}
	
	return ioutil.WriteFile(path, bytes, 0666)
}
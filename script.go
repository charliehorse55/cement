package main

import (
	"os"
	// "bufio"
	"fmt"
	Ecem "github.com/charliehorse55/libcement"
)

// func loadScript(path string) (Ecem.Painting, error) {
// 	
// }

func saveScript(filename string, p Ecem.Painting) error {
	outfile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Failed to open output file: %s", filename)
	}
	defer outfile.Close()
	
	fmt.Fprintf(outfile, "#background\n")
	fmt.Fprintf(outfile, "%s\n\n", p[0].Filename)
	
	fmt.Fprintf(outfile, "#vectors\n")
	q := p[1:]
	for i := range q {
		fmt.Fprintf(outfile, "%-30s\t%5.3f\t%5.3f\t%5.3f\n", q[i].Filename, q[i].Intensity.R, q[i].Intensity.G, q[i].Intensity.B)
	}
	
	return nil
}
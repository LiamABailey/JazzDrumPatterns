package svgutil 

import (
	"encoding/xml"
	"errors"
)

// given two SVGs X, Y(as text), returns a 
// string representing the combined files: 
// The metadata from the first, and all <g>
// from each
func CombineSVG(svg1, svg2 string) (string, error) {
	svg2Groups, g2err := getGroups(svg2)
	if g2err != nil {
		return "", errors.New("Failed to recover groups from second SVG")
	}
	combinedSvg, comboerr := insertGroups(svg1, svg2Groups)
	if comboerr != nil {
		return "", errors.New("Failed to merge collected groups into second SVG")
	}
	return scombinedSvg, nil 
}

// recover all group data from the SVG, retaining depth
func getGroups(svg string) ([]string, error) {
	return make([]string, 0), nil 
}

// insert groups into an SVG at the first layer (cshild of the
// svg tag)
func insertGroups(svg string, groups []string) (string, error) { 
	return "", nil
}
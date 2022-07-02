package svgutil 

import (
	"fmt"
	"encoding/xml"
	"errors"
)

// given two SVGs X, Y(as text), returns a 
// string representing the combined files: 
// The metadata from the first, and all <g>
// from each
func CombineSVG(svg1, svg2 []byte) ([]byte, error) {
	svg2Groups, g2err := getGroups(svg2)
	if g2err != nil {
		return make([]byte, 0), errors.New(fmt.Sprintf("Failed to recover groups from second SVG: %s", g2err.Error()))
	}
	combinedSvg, comboerr := insertGroups(svg1, svg2Groups)
	if comboerr != nil {
		return make([]byte, 0), errors.New("Failed to merge collected groups into second SVG")
	}
	return combinedSvg, nil 
}

func getGroups(svg []byte) ([]G, error) {
	svgData, e := bytesToSVG(svg)
	if e != nil {
		return make([]G, 0), e
	}
	return svgData.Groups, nil 
}


// insert groups into an SVG at the first layer (cshild of the
// svg tag)
func insertGroups(svg []byte, groups []G) ([]byte, error) { 
	svgData, e := bytesToSVG(svg)
	if e != nil {
		return svg, e
	}
	svgData.Groups = append(svgData.Groups, groups...)
	remarshaled, re := xml.Marshal(svgData)
	return remarshaled, re 
}

func bytesToSVG(svg []byte) (SVG, error) {
	var svgData SVG
	e := xml.Unmarshal(svg, &svgData)
	return svgData, e
}
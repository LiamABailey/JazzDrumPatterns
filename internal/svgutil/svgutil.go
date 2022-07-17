package svgutil 

import (
	"fmt"
	"encoding/xml"
	"errors"
)

// given two SVGs X, Y returns a
// []byte containing the <g> and <path> (group, path) elements
// in the second level of Y (children of <svg>)
// as children of <svg> in X.
func CombineSVG(svg1, svg2 []byte) ([]byte, error) {
	svg2Groups,svg2Paths, g2err := getGroupsPaths(svg2)
	if g2err != nil {
		return make([]byte, 0), errors.New(fmt.Sprintf("Failed to recover groups from second SVG: %s", g2err.Error()))
	}
	combinedSvg, comboerr := insertGroupsPaths(svg1, svg2Groups, svg2Paths)
	if comboerr != nil {
		return make([]byte, 0), errors.New("Failed to merge collected groups into second SVG")
	}
	return combinedSvg, nil 
}

func getGroupsPaths(svg []byte) ([]E, []E, error) {
	svgData, e := bytesToSVG(svg)
	if e != nil {
		return make([]E, 0),make([]E, 0), e
	}
	return svgData.Groups, svgData.Paths, nil 
}


// insert groups, paths into an SVG at the first layer 
// (child of the svg tag)
func insertGroupsPaths(svg []byte, groups []E, paths []E) ([]byte, error) { 
	svgData, e := bytesToSVG(svg)
	if e != nil {
		return svg, e
	}
	svgData.Groups = append(svgData.Groups, groups...)
	svgData.Paths = append(svgData.Paths, paths...)
	remarshaled, re := xml.Marshal(svgData)
	return remarshaled, re 
}

func bytesToSVG(svg []byte) (SVG, error) {
	var svgData SVG
	e := xml.Unmarshal(svg, &svgData)
	return svgData, e
}
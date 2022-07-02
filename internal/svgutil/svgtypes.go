package svgutil

import "encoding/xml"

type SVG struct {
	XMLName 	xml.Name
	Groups 		[]G			`xml:"g"`
}

type G struct {
	XMLName xml.Name
	Content []byte		`xml:",innerxml"`
}
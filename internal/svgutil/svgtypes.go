package svgutil

import "encoding/xml"

type SVG struct {
	XMLName xml.Name
	Attrs 	[]xml.Attr 	`xml:",any,attr"`
	Groups 	[]G			`xml:"g"`
	Content []byte		`xml:",innerxml"`
}

type G struct {
	XMLName xml.Name
	Content []byte		`xml:",innerxml"`
}
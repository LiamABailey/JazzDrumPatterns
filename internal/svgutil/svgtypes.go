package svgutil

import "encoding/xml"

type SVG struct {
	XMLName 	xml.Name
	Groups 		[]E			`xml:"g"`
	Paths		[]E			`xml:"path"`
}

type E struct {
	XMLName xml.Name
	Transform	string		`xml:"transform,attr"`
	D			string		`xml:"d,attr"`
	Style		string		`xml:"style,attr"`
	Content 	[]byte		`xml:",innerxml"`
}

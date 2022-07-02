module patterngen

go 1.16

require (
	assets/beatimages v1.0.0
	github.com/gin-gonic/gin v1.8.1
	golang.org/x/net v0.0.0-20211118161319-6a13c67c3ce4 // indirect
	internal/patterns v1.0.0
	internal/svgutil v0.0.0-00010101000000-000000000000
)

replace assets/beatimages => ./../../assets/beatimages

replace internal/patterns => ./../../internal/patterns

replace internal/svgutil => ./../../internal/svgutil

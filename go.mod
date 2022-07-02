module jazzdrumpatterns

go 1.16

require (
	api/patterngen v1.0.0
	internal/svgutil v0.0.0-00010101000000-000000000000 // indirect
)

replace api/patterngen => ./api/patterngen

replace assets/beatimages => ./assets/beatimages

replace internal/patterns => ./internal/patterns

replace internal/svgutil => ./internal/svgutil

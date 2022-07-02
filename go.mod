module jazzdrumpatterns

go 1.16

require (
	api/patterngen v1.0.0
	github.com/srwiley/oksvg v0.0.0-20220128195007-1f435e4c2b44 // indirect
)

replace api/patterngen => ./api/patterngen

replace internal/patterns => ./internal/patterns

replace assets/beatimages => ./assets/beatimages

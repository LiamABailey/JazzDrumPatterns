module jazzdrumpatterns

go 1.16

require api/patterngen v1.0.0

replace api/patterngen => ./api/patterngen

replace internal/patterns => ./internal/patterns

replace assets/beatimages => ./assets/beatimages

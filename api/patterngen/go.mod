module patterngen

go 1.16

require (
	assets/beatimages v1.0.0
	github.com/gin-gonic/gin v1.8.1
	github.com/srwiley/oksvg v0.0.0-20220128195007-1f435e4c2b44
	internal/patterns v1.0.0
)

replace internal/patterns => ./../../internal/patterns

replace assets/beatimages => ./../../assets/beatimages

package patterngen

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"strings"
)

type PatternServer struct {
	router gin.Engine
}

func NewPatternServer() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return 	(strings.HasPrefix(origin, "http://localhost")) || (strings.HasPrefix(origin, "http://127.0.0.1"))
 		},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	r.GET("/GetMeasure", getMeasure)
	r.GET("/GetMeasureFromBeats", generateMeasureFromBeats)
	r.GET("/GetMeasureImage", getMeasureImage)

	return r
}

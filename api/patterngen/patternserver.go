package patterngen
import (
  "github.com/gin-gonic/gin"
)

type PatternServer struct {
  router gin.Engine
}

func NewPatternServer() *gin.Engine {
  r := gin.Default()
  r.GET("/GetMeasures", getMeasures)

  return r
}

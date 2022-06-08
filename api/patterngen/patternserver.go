package patterngen
import (
  "github.com/gin-gonic/gin"
)

type PatternServer struct {
  router gin.Router
}

func NewPatternServer() *PatternServer {
  // add routes here
}

func (p *PatternServer) Run() {
  p.router.run()
}

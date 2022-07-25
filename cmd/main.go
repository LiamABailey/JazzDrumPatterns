package main

import (
  "api/patterngen"
)

func main() {
  patternsrv := patterngen.NewPatternServer()
  patternsrv.Run(":8050")
}

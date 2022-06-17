// +build exclude

package patterngen

import (
  "image"
  "image/png"
  "image/draw"
  "os"

  "github.com/gin-gonic/gin"
)


// given a collection of beats (a measure),
// produce an image
func getMeasureImage(ctx *gin.Context) {
  // TODO: implement
}

// given a known ride/hihat/snare/kick pattern,
// compose and return a single beat
func getBeatImage(ride, hihat, snare, kict int) (image.Image, error){
  // TODO: implement
}

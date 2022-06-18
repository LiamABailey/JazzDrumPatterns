package patterngen

import (
  "embed"
  "fmt"
  "io/fs"

  "github.com/gin-gonic/gin"

  "assets/beatimages"
)
//"image"
//"image/png"
//"image/draw"


// given a collection of beats (a measure),
// produce an image
//func getMeasureImage(ctx *gin.Context) {
  // TODO: implement
  // step 1: parse each beat in the request
  // step 2: get images for each beat
  // step 3: concatenate images
//}

// given a known ride/hihat/snare/kick pattern,
// compose and return a single beat
//func getBeatImage(ride, snare, kick, hihat int) (image.Image, error){
  // TODO: implement
  // step 1: get the image associated with the [ride, snare] or [kick, hihat] group
  // step 2: layer the images to produce a composite and return
//}

// retrieve a single beat component image from a given source
func retrieveImage(rhythm1, rhythm2 int, dir embed.FS) ([]byte, error){
  fmt.Println(rhythm1, rhythm2)

  testNames, _ := fs.Glob(dir, fmt.Sprintf("%s/*.png", getDirName(dir)))
  return fs.ReadFile(dir,  testNames[0])
  // TODO: impelement
  // imname := fmt.Sprintf("%s_%s.png", rhythm1, rhythm2)

  // construct path
  // check if file exists. If so, return w/o error. else, return nil and FileNotFound. 
}

func TestRetrieve(ctx *gin.Context) {
  fmt.Println(ctx)
  data, err := retrieveImage(0,0, beatimages.RideSnareImages)
  fmt.Println(data, err)
}

// get the name of the root of the
// provided directory
func getDirName(dir embed.FS) string {
  entries, _  := dir.ReadDir(".")
  return entries[0].Name()
}
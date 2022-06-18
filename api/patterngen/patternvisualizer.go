package patterngen

import (
  "embed"
  "encoding/base64"
  "fmt"
  "io/fs"
  "net/http"

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
func retrieveImage(rhythm1, rhythm2 int, dir embed.FS) (string, error) {
  imName := fmt.Sprintf("%d_%d.svg", rhythm1, rhythm2)
  imbytes, rerr := fs.ReadFile(dir,  fmt.Sprintf("%s/%s", getDirName(dir), imName))
  if rerr != nil {
    return "", rerr 
  }
  imb64 := base64.StdEncoding.EncodeToString(imbytes)
  return imb64, nil

}

func TestRetrieve(ctx *gin.Context) {
  fmt.Println(ctx)
  data, err := retrieveImage(0,0, beatimages.RideSnareImages)
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
  }
  ctx.JSON(http.StatusOK, gin.H{"svg": data})
}

// get the name of the root of the
// provided directory
func getDirName(dir embed.FS) string {
  entries, _  := dir.ReadDir(".")
  return entries[0].Name()
}
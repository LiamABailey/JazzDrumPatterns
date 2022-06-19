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

const (
  querySep string = ","
)


// given a collection of beats (a measure),
// produce an image
func getMeasureImage(ctx *gin.Context) {
  // get each individual beat: ride,snare,bass,hh
  beatSegments, qerr := splitBeatQuery(ctx.Request.URL.Query())
  if qerr != nil {
    ctx.JSON(http.StatusBadRequest,
        gin.H{"error": fmt.Sprintf("Unable to parse query: %s", qerr.Error())})
  }
  // for each segment
  for _, bSeg := range beatSegments {
    //split into ride/snare/bass/hh 
    beatComponents, serr := splitConvertPatternLists(bSeg, querySep)
    if serr != nil {
      ctx.JSON(http.StatusBadRequest,
        gin.H{"error": fmt.Sprintf("Unable to parse query segment %s: %s",bSeg, serr.Error())})
    }
    if len(beatComponents) != 4 {
      ctx.JSON(http.StatusBadRequest,
        gin.H{"error": fmt.Sprintf("Expected 4 segments: ride/snare/bass/hh, received: %s",bSeg)})
    }
    imRS, imRSerr := retrieveImage(beatComponents[0], beatComponents[1], beatimages.RideSnareImages)
    imBHH, imBHHerr := retrieveImage(beatComponents[2], beatComponents[3], beatimages.KickHiHatImages)
    if (imRSerr != nil) || (imBHHerr != nil){
      ctx.JSON(http.StatusInternalServerError,
        gin.H{"error": fmt.Sprintf("Unable to retrieve images. Errors: Ride/Snare: %s; Bass/HiHat: %s", imRSerr.Error(), imBHHerr.Error())})
    }
    ctx.JSON(http.StatusOK, gin.H{"Ride-Snare": imRS, "Bass-HiHat": imBHH})
  }

}

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
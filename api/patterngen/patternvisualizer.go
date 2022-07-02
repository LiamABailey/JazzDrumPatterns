package patterngen

import (
  "embed"
  "encoding/base64"
  "fmt"
  "io/fs"
  "net/http"
  "strconv"

  "github.com/gin-gonic/gin"

  "assets/beatimages"
  "internal/svgutil"
)

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
        return
  }
  // for each segment
  returnData := gin.H{}
  for beatIx, bSeg := range beatSegments {
    //split into ride/snare/bass/hh 
    beatComponents, serr := splitConvertPatternLists(bSeg, querySep)
    if serr != nil {
      ctx.JSON(http.StatusBadRequest,
        gin.H{"error": fmt.Sprintf("Unable to parse query segment %s: %s",bSeg, serr.Error())})
        return
    }
    if len(beatComponents) != 4 {
      ctx.JSON(http.StatusBadRequest,
        gin.H{"error": fmt.Sprintf("Expected 4 segments: ride/snare/bass/hh, received: %s",bSeg)})
        return
    }
    imRS, imRSerr := retrieveImage(beatComponents[0], beatComponents[1], beatimages.RideSnareImages)
    imBHH, imBHHerr := retrieveImage(beatComponents[2], beatComponents[3], beatimages.KickHiHatImages)
    if (imRSerr != nil) || (imBHHerr != nil){
      ctx.JSON(http.StatusInternalServerError,
        gin.H{"error": fmt.Sprintf("Unable to retrieve images. Errors: Ride/Snare: %s; Bass/HiHat: %s", imRSerr.Error(), imBHHerr.Error())})
      return
    }
    beatIxStr := strconv.Itoa(beatIx)
    mergedSvg, mergeErr := svgutil.CombineSVG(imRS, imBHH)
    if mergeErr != nil {
      ctx.JSON(http.StatusInternalServerError,
        gin.H{"error": fmt.Sprintf("Unable to merge images. Error: %s", mergeErr.Error())})
      return 
    }
    imb64 := base64.StdEncoding.EncodeToString(mergedSvg)
    returnData[beatIxStr] = gin.H{"image":imb64}
  }
  ctx.JSON(http.StatusOK, returnData)

}

// given a known ride/hihat/snare/kick pattern,
// compose and return a single beat
//func getBeatImage(ride, snare, kick, hihat int) (image.Image, error){
  // TODO: implement
  // step 1: get the image associated with the [ride, snare] or [kick, hihat] group
  // step 2: layer the images to produce a composite and return
//}

// retrieve a single beat component image from a given source
func retrieveImage(rhythm1, rhythm2 int, dir embed.FS) ([]byte, error) {
  imName := fmt.Sprintf("%d_%d.svg", rhythm1, rhythm2)
  imbytes, rerr := fs.ReadFile(dir,  fmt.Sprintf("%s/%s", getDirName(dir), imName))
  if rerr != nil {
    return make([]byte, 0), rerr 
  }
  return imbytes, nil
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
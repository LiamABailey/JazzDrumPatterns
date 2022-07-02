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
    beatEle, serr := splitConvertPatternLists(bSeg, querySep)
    if serr != nil {
      ctx.JSON(http.StatusBadRequest,
        gin.H{"error": fmt.Sprintf("Unable to parse query segment %s: %s",bSeg, serr.Error())})
        return
    }
    if len(beatEle) != 4 {
      ctx.JSON(http.StatusBadRequest,
        gin.H{"error": fmt.Sprintf("Expected 4 segments: ride/snare/bass/hh, received: %s",bSeg)})
        return
    }
    beatIxStr := strconv.Itoa(beatIx)
    imb64, imerr := getBeatImageStr(beatEle[0], beatEle[1], beatEle[2], beatEle[3])
    if imerr != nil {
      ctx.JSON(http.StatusBadRequest,
        gin.H{"error": fmt.Sprintf("Encountered while generating beat %s image: %s", beatIxStr, imerr)})
        return
    }
    returnData[beatIxStr] = gin.H{"image":imb64}
  }
  ctx.JSON(http.StatusOK, returnData)

}

// given a known ride/hihat/snare/kick pattern,
// compose and return a single beat image as a b64 string
func getBeatImageStr(ride, snare, kick, hihat int) (string, error){
  imRS, imRSerr := retrieveImage(ride, snare, beatimages.RideSnareImages)
  imBHH, imBHHerr := retrieveImage(kick, hihat, beatimages.KickHiHatImages)
  if (imRSerr != nil) || (imBHHerr != nil){
    return  "", fmt.Errorf("Unable to retrieve images. Errors: Ride/Snare: %s; Bass/HiHat: %s", imRSerr.Error(), imBHHerr.Error())
  }
  mergedSvg, mergeErr := svgutil.CombineSVG(imRS, imBHH)
  if mergeErr != nil {
    return "", fmt.Errorf("Unable to merge images. Error: %s", mergeErr.Error())
  }
  imb64 := base64.StdEncoding.EncodeToString(mergedSvg)
  return imb64, nil
}

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
package patterngen

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "strings"
  "strconv"

  "github.com/gin-gonic/gin"

  "internal/patterns"
)

const (
  defaultPatterns string = "0"
  defaultBeats string = "4"
  beatComponents int = 4
  innerSep string = "|"
  outerSep string = ","
)

// Generate a single measure of beats
// the allowed patterns for each limb +
// number of desired beats
func getMeasure(ctx *gin.Context) {
  bass := ctx.DefaultQuery("bass", defaultPatterns)
  hihat := ctx.DefaultQuery("hihat", defaultPatterns)
  ride := ctx.DefaultQuery("ride", defaultPatterns)
  snare := ctx.DefaultQuery("snare", defaultPatterns)
  beats, _ := strconv.Atoi(ctx.DefaultQuery("beats", defaultBeats))

  bassPatterns, berr := splitConvertPatternLists(bass, outerSep)
  hihatPatterns, hherr := splitConvertPatternLists(hihat, outerSep)
  ridePatterns, rerr := splitConvertPatternLists(ride, outerSep)
  snarePatterns, serr := splitConvertPatternLists(snare, outerSep)

  if berr != nil || hherr != nil || rerr != nil || serr != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": "Received malformed pattern parameter"})
    return
  }

  consBeats := make([]patterns.Beat, beats)
  for i := 0; i < beats; i++ {
    consBeats[i] = *patterns.GenerateBeatPattern(
                        ridePatterns, snarePatterns,
                          bassPatterns, hihatPatterns)
  }
  meas := patterns.NewMeasure(consBeats)
  pattern, _ := json.Marshal(*meas)
  ctx.JSON(http.StatusOK, gin.H{"pattern": string(pattern)})
}

// Given the construction request of one or more beats,
// produce a measure meeting the constraints
func generateMeasureFromBeats(ctx *gin.Context) {
  measureDefinitions, qerr := splitBeatQuery(ctx.Request.URL.Query())
  if qerr != nil {
    ctx.JSON(http.StatusBadRequest,
        gin.H{"error": fmt.Sprintf("Unable to parse query: %s", qerr.Error())})
  }
  // generate a pattern from each beat definition
  parsedMeasDef := make([]patterns.Beat, len(measureDefinitions))
  for i, mDef := range measureDefinitions {
    // get each component of the beat
    sepParts := strings.Split(mDef, outerSep)
    if len(sepParts) != beatComponents {
      log.Println("Expected beat construction request to have ", beatComponents,
                  " components, received ", len(sepParts))
      ctx.JSON(http.StatusBadRequest,
            gin.H{"error": fmt.Sprintf("Each beat must have %d components", beatComponents)})
    }
    // convert the requested compoent IDs to integers
    parsedBeatDefinition := make([][]int, beatComponents)
    for c := 0; c < beatComponents; c++ {
      var err error
      parsedBeatDefinition[c], err = splitConvertPatternLists(sepParts[c], innerSep)
      if err != nil{
        log.Println("Encountered error trying to parse beat components: ", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Couldn't parse %s", sepParts[c])})
      }
    }
    // generate the beat & store
    parsedMeasDef[i] = *patterns.GenerateBeatPattern(
                    parsedBeatDefinition[0],
                    parsedBeatDefinition[1],
                    parsedBeatDefinition[2],
                    parsedBeatDefinition[3])
  }
  meas := patterns.NewMeasure(parsedMeasDef)
  pattern, _ := json.Marshal(*meas)
  ctx.JSON(http.StatusOK, gin.H{"pattern": string(pattern)})
}



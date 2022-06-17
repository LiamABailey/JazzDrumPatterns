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
  beatIdentifier string = "b"
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

  bassPatterns, berr := convertPatternLists(bass, outerSep)
  hihatPatterns, hherr := convertPatternLists(hihat, outerSep)
  ridePatterns, rerr := convertPatternLists(ride, outerSep)
  snarePatterns, serr := convertPatternLists(snare, outerSep)

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

//
func generateMeasureFromBeats(ctx *gin.Context) {
  // looks for b# identifiers, starting with b0
  ok := true
  ix := 0
  var measureDefinitions []string
  for ok {
    beatId := beatIdentifier + strconv.Itoa(ix)
    var beatDef string
    beatDef, ok = ctx.GetQuery(beatId)
    if ok {
      measureDefinitions = append(measureDefinitions, beatDef)
    }
    ix += 1
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
            gin.H{"error": fmt.Sprintf("Each beat must have %s components", beatComponents)})
      return
    }
    // convert the requested compoent IDs to integers
    parsedBeatDefinition := make([][]int, beatComponents)
    for c := 0; c < beatComponents; c++ {
      var err error
      parsedBeatDefinition[c], err = convertPatternLists(sepParts[c], innerSep)
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

func convertPatternLists(pat string, sep string) ([]int, error) {
  sepPat := strings.Split(pat, sep)
  iPat := make([]int, len(sepPat))
  var cerr error
  for i, v := range sepPat {
    iPat[i], cerr = strconv.Atoi(v)
    if cerr != nil {
      return nil, cerr
    }
  }
  return iPat, nil
}

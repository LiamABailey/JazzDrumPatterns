package patterngen

import (
  "encoding/json"
  "fmt"
  "net/http"
  "strings"
  "strconv"

  "github.com/gin-gonic/gin"

  "internal/patterns"
)

const (
  defaultPatterns string = "0"
  defaultBeats string = "4"
  defaultMeasures string = "1"
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
  //measures, _:= strconv.Atoi(ctx.DefaultQuery("measures", defaultMeasures))

  bassPatterns, berr := convertPatternLists(bass)
  hihatPatterns, hherr := convertPatternLists(hihat)
  ridePatterns, rerr := convertPatternLists(ride)
  snarePatterns, serr := convertPatternLists(snare)

  if berr != nil || hherr != nil || rerr != nil || serr != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": "Received malformed pattern parameter"})
    return
  }

  consBeats := make([]patterns.Beat, beats)
  for i := 0; i < beats; i++ {
    fmt.Println(ridePatterns)
    consBeats[i] = *patterns.GenerateBeatPattern(
                        ridePatterns, snarePatterns,
                          bassPatterns,hihatPatterns)
  }
  meas := patterns.NewMeasure(consBeats)
  pattern, _ := json.Marshal(*meas)
  ctx.JSON(http.StatusOK, gin.H{"pattern": string(pattern)})
}


func convertPatternLists(pat string) ([]int, error) {
  sepPat := strings.Split(pat, ",")
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

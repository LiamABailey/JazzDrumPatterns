package patterngen

import (
  "strings"
  "strconv"

  "net/http"
  "github.com/gin-gonic/gin"

  "internal/patterns"
)

const (
  defaultPatterns string = "0"
  defaultBeats string = "4"
  defaultMeasures string = "1"
)
// Generate one or more measures of beats given
// the allowed patterns for each limb +
// number of desired beats
func getMeasures(ctx *gin.Context) {
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
  // TODO: What if multiple measures are requested?
  consBeats := make([]patterns.Beat, beats)
  for i := 0; i < beats; i++ {
    consBeats[i] = *patterns.GenerateBeatPattern(
                        ridePatterns, snarePatterns,
                          bassPatterns,hihatPatterns)
  }

  // TODO: how to structure and return?
  ctx.JSON(http.StatusOK, gin.H{"message": "created pattern"})
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

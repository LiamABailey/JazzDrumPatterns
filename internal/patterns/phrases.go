package patterns

type Beat struct {
  Ride  int `json: ride`
  Snare int  `json: snare`
  Bass  int  `json: bass`
  Hh    int  `json: hihat`
}

func NewBeat(ride, snare, bass, hh int) *Beat {
  return &Beat{ride, snare, bass, hh}
}


type Measure struct {
  Beats []Beat `json: beats`
}

func NewMeasure(beats []Beat) *Measure {
  return &Measure{beats}
}

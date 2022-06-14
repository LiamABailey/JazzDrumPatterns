package patterns

type Beat struct {
  Ride  int `json: ride_pattern`
  Snare int  `json: snare_pattern`
  Bass  int  `json: bass_pattern`
  Hh    int  `json: hh_pattern`
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

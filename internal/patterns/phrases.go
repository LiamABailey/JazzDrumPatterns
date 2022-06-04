package patterns

type Beat struct {
  ride  int `json: ride_pattern`
  snare int  `json: snare_pattern`
  bass  int  `json: bass_pattern`
  hh    int  `json: hh_pattern`
}

func NewBeat(ride, snare, bass, hh int) *Beat {
  return &Beat{ride, snare, bass, hh}
}


type Measure struct {
  beats []Beat
}

func NewMeasure(beats []Beat) *Measure {
  return &Measure{beats}
}

package patterns

import (
  "math/rand"
  "time"
)

// each of the 8 combinations of possible notes played
// within a triplet is represented as an integer, 0-7
// We can pick from any 'eligible' patterns at random
func PickPattern(p []int) int {
  // seed the source of randomneess
  rand_src := rand.NewSource(time.Now().UnixNano())
  r := rand.New(rand_src)
  // return a value at a random index
  return p[r.Intn(len(p))]
}

// generate a random beat from allowed ride/snare/bass/hh patterns
func GenerateBeatPattern(ride_op, snare_op, bass_op, hh_op []int) *Beat{
  ride := PickPattern(ride_op)
  snare := PickPattern(snare_op)
  bass := PickPattern(bass_op)
  hh := PickPattern(hh_op)
  return NewBeat(ride, snare, bass, hh)
}

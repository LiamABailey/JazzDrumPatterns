package main

import (
  "fmt"

  "internal/patterns"
)

func main() {
  fmt.Println(patterns.GenerateBeatPattern([]int{0,1,2},[]int{4,5},
                                           []int{1},[]int{1}))
}

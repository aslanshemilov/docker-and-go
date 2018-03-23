//Square root function
//Given a number x, we want to find the number z for which z^2 is most nearly x.
//Computers typically compute the square root of x using a loop. 
//Starting with some guess z, we can adjust z based on how close z^2 is to x, producing a better guess: z -= (z*z - x)/ (2*z)

package main

import (
    "fmt"
    "math"
)

const DELTA = 0.0000001
const z_0 = 100.0

func Sqrt(x float64) float64 {
    z = z_0
    
  step := func() float64 {
    return z - (z*z -x) / (2*z)
  }
  for zz := step(); math.Abs(zz - z) > DELTA
  {
  z = zz
    zz = step()
  }
  return 
}

func main() {
  fmt.Println(Sqrt(500))
  fmt.Println(math.Sqrt(500))
}

package common

import (
	"math"
  "math/rand"
  "crypto/sha512"
  "encoding/binary"
  "time"
  "strings"
  // "fmt"
)

// STRINGS
func Split(what string) []string { return strings.Split(what, "|") }
func ElemInList(what string, list []string) bool { if len(list) > len(Elements) {return false}; for _, each := range list { if each==what {return true} }; return false }
// func ParseTags(what string) map[string]string {
//   buffer, tags := make(map[string]string), Split(what)
//   for _, each := range tags { 
//     row := strings.Split(each, "=") 
//     if len(row) == 2 { buffer[row[0]] = row[1] } 
//   }
//   return buffer
// }


// FLOATS
func Log2(a float64) float64 { return math.Log2(a+1) }
func Log7(a float64) float64 { return math.Log2(a+1)/math.Log2(7) }
func Log10(a float64) float64 { return math.Log10(a+1) }
func Popow(a float64) float64 { return math.Pow(a,a) }
func Cbrt(a float64) float64 { return math.Cbrt(a+1)-1 }

func IsWithin(rand, base float64) bool { 
  entropy := Log10(base)/25
  if (1+0.0132437+entropy)*base+0.001 > rand && (1+0.0132437+entropy)*base-0.001 < rand { return true }
  // if (1+0.0132437+entropy)*(1+0.0132437+entropy)*base+0.001 > rand && (1+0.0132437+entropy)*(1+0.0132437+entropy)*base-0.001 < rand { return true }
  if base/(1+0.0132437+entropy)+0.001 > rand && base/(1+0.0132437+entropy)-0.001 < rand { return true }
  // if base/(1+0.0132437+entropy)/(1+0.0132437+entropy)+0.001 > rand && base/(1+0.0132437+entropy)/(1+0.0132437+entropy)-0.001 < rand { return true }
  if base+0.001 > rand && base-0.001 < rand { return true }
  return false
}
func Ntrp(a float64) float64 { 
  randy := Epoch()/250 % 4
  entropy := Log10(a)/25 
  if randy == 1 { a = a*(1+0.0132437+entropy) }
  // if randy == 2 || randy == 3 { a = a*(1+0.0132437+entropy)*(1+0.0132437+entropy) }
  if randy == 3 { a = a/(1+0.0132437+entropy) }
  // if randy == 7 || randy == 8 { a = a/(1+0.0132437+entropy)/(1+0.0132437+entropy) }
  return math.Round( a*1000 ) / 1000
}

func Rand() float64 {
  x := (time.Now().UnixNano())
  in_bytes := make([]byte, 8)
  binary.LittleEndian.PutUint64(in_bytes, uint64(x))
  hsum := sha512.Sum512(in_bytes)
  sum  := binary.BigEndian.Uint64(hsum[:])
  return rand.New(rand.NewSource( int64(sum) )).Float64()
}

func Between(xyz, abc [3]int) float64 { return Vector( float64(xyz[0]-abc[0]), float64(xyz[1]-abc[1]), float64(xyz[2]-abc[2]) ) }
func Vector(props ...float64) float64 {
  sum := 0.0
  for _, each := range props { sum += each*each }
  return math.Sqrt(sum)
}


// FLOAT to INT
func Round(a float64) int { return int(math.Round(a)) } // for create stream
func CeilRound(a float64) int { return int(math.Ceil(a)) } // for create stream
func FloorRound(a float64) int { return int(math.Floor(a)) } // for filter streams
func ChancedRound(a float64) int {
  b,l:=math.Ceil(a),math.Floor(a)
  c:=math.Abs(math.Abs(a)-math.Abs(math.Min(b, l)))
  if a<0 {c = 1-c}
  if Rand() < c {return int(b)} else {return int(l)}
  return 0
}

func BornLuck(time int) int { if time%10 == 0 {return 2} else if time%10 == 9 {return 5} else if time%10 < 5 {return 3} else {return 4} ; return 0}
func Epoch() int { return int(time.Now().UnixNano())/1000000 }
func EpochNS() int { return int(time.Now().UnixNano()) }

// TIME
func Wait(ms float64) { time.Sleep( time.Millisecond * time.Duration( ms )) }
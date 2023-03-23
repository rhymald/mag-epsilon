package common

import (
	"math"
  "math/rand"
  "crypto/sha512"
  "encoding/binary"
  "time"
  "strings"
)

func Split(what string) []string { return strings.Split(what, "|") }
// XYZ+RRR from str
// str to XYZ+RRR

func Log(a float64) float64 { return math.Log2(a+1)/math.Log2(math.Pi) }

func Rand() float64 {
  x := (time.Now().UnixNano())
  in_bytes := make([]byte, 8)
  binary.LittleEndian.PutUint64(in_bytes, uint64(x))
  hsum := sha512.Sum512(in_bytes)
  sum  := binary.BigEndian.Uint64(hsum[:])
  return rand.New(rand.NewSource( int64(sum) )).Float64()
}

func Vector(props ...float64) float64 {
  sum := 0.0
  for _, each := range props { sum += each*each }
  return math.Sqrt(sum)
}

func Round(a float64) int { return int(math.Round(a)) } // for create stream
func CeilRound(a float64) int { return int(math.Ceil(a)) } // for create stream
// func FloorRound(a float64) int { return int(math.Floor(a)) } // for filter streams
func ChancedRound(a float64) int {
  b,l:=math.Ceil(a),math.Floor(a)
  c:=math.Abs(math.Abs(a)-math.Abs(math.Min(b, l)))
  if a<0 {c = 1-c}
  if Rand() < c {return int(b)} else {return int(l)}
  return 0
}

func BornLuck(time int) int { if time%10 == 0 {return 2} else if time%10 == 9 {return 5} else if time%10 < 5 {return 3} else {return 4} ; return 0}
func Epoch() int { return int(time.Now().UnixNano()) }
func Wait(ms float64) { time.Sleep( time.Millisecond * time.Duration( ms )) }
package common

import (
	"math"
  "math/rand"
  "crypto/sha512"
  "encoding/binary"
  "time"
)

type Dot map[string]int 
type Stream map[string][3]int

var Elements []string = []string{"◌ ", "🌪 ", "🔥"}//, "🪨", "🧊", "🌑", "🩸", "🎶", "☀️ "} 
var Physical []string = []string{"◌ ", "🌱",  "🪵", "🪨", "🛡 "} // none, flesh/plant (punchD), wood/shell (stingA), stone (chopDa), armored

func (str *Stream) Cre() float64 { buf := *str ; return float64(buf[str.Elem()][0])/1000 }
func (str *Stream) Alt() float64 { buf := *str ; return float64(buf[str.Elem()][1])/1000 }
func (str *Stream) Des() float64 { buf := *str ; return float64(buf[str.Elem()][2])/1000 }
func (str *Stream) Elem() string { for elem, _ := range *str { return elem } ; return "ERR" }
func BRandNewStream(mean int, elem string) *Stream {
	c, a, d := 1.0, 1.0, 1.0
	if mean>0 { for x:=0; x<mean-1; x++ { c, a, d = c+Rand()+Rand(), a+Rand()+Rand(), d+Rand()+Rand() } }
	modifier := float64(mean) / (Vector(c, a, d) / math.Sqrt(3))
	return &Stream{ elem: [3]int{ ChancedRound(c*modifier), ChancedRound(a*modifier), ChancedRound(d*modifier) }}
}

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

// func CeilRound(a float64) int { return int(math.Ceil(a)) }
// func FloorRound(a float64) int { return int(math.Floor(a)) }
func ChancedRound(a float64) int {
  b,l:=math.Ceil(a),math.Floor(a)
  c:=math.Abs(math.Abs(a)-math.Abs(math.Min(b, l)))
  if a<0 {c = 1-c}
  if Rand() < c {return int(b)} else {return int(l)}
  return 0
}

func Epoch() int { return int(time.Now().UnixNano()) }
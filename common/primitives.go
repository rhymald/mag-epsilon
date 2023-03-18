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

var Elements []string = []string{"◌ ", "🌪 ", "🔥", "🪨", "🧊", "🌑", "🩸", "🎶", "☀️ "} 
var Physical []string = []string{"◌ ", "🌱",  "🪵", "🪨", "🛡 "} // none, flesh/plant (punchD), wood/shell (stingA), stone (chopDa), armored

func (dot *Dot) Weight() float64 { buf := *dot ; return float64(buf[dot.Elem()]) }
func (dot *Dot) Elem() string { for elem, _ := range *dot { return elem } ; return "ERR" }
func (str *Stream) EmitDot() *Dot { return &Dot{ str.Elem(): ChancedRound( math.Log2(2+Vector(str.Des(0),str.Alt(0),str.Cre(0))) )} }
func (str *Stream) Vol(add float64) float64 { return str.Cre(add)*str.Alt(add)*str.Des(add)-add*add*add }
func (str *Stream) Len(add float64) float64 { return Vector(str.Cre(0),str.Alt(0),str.Des(0))+add }
func (str *Stream) Cre(add float64) float64 { buf := *str ; return float64(buf[str.Elem()][0])/1024+add }
func (str *Stream) Alt(add float64) float64 { buf := *str ; return float64(buf[str.Elem()][1])/1024+add }
func (str *Stream) Des(add float64) float64 { buf := *str ; return float64(buf[str.Elem()][2])/1024+add }
func (str *Stream) Elem() string { for elem, _ := range *str { return elem } ; return "ERR" }
func BRandNewStream(elem string, leng int) *Stream {
	c, a, d := math.Pi+Rand()+Rand(), math.Pi+Rand()+Rand(), math.Pi+Rand()+Rand()
	for r:=0; r<leng; r++ { c, a, d = c+math.Pi+Rand()+Rand(), a+math.Pi+Rand()+Rand(), d+math.Pi+Rand()+Rand() }
	modifier := float64(leng) / Vector(c, a, d)
	return &Stream{ elem: [3]int{ CeilRound(c*1024*modifier), CeilRound(a*1024*modifier), CeilRound(d*1024*modifier) }}
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
func Epoch() int { return int(time.Now().UnixNano()) }
func Wait(ms float64) { time.Sleep( time.Millisecond * time.Duration( ms )) }
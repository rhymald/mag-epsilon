package fancy

import "fmt"

var (
  // Extra
  B = "\u001b[1m"
  U = "\u001b[4m"
  R = "\u001b[7m"
  E = []string{
    // "\u001b[48;5;245m", //Grey
    "\u001b[0m", //Reset
    "\u001b[38;5;70m", //Air
    "\u001b[38;5;208m", //Fire
    "\u001b[38;5;95m", //Earth
    "\u001b[38;5;39m", //Water
    "\u001b[38;5;54m", //Void
    "\u001b[38;5;124m", //Mallom
    "\u001b[38;5;244m", //Noise
    "\u001b[38;5;229m", //Light
  }
)

func Err() string { return fmt.Sprintf("%s%s%s", B, R, E[6]) }
func Wrn() string { return fmt.Sprintf("%s%s%s", B, U, E[8]) }
func Clr(e int) string { if e > 0 && e < 9 { return fmt.Sprintf("%s%s", B, E[e]) } ; return fmt.Sprintf("%s%s", B, E[0]) }
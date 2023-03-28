package common

import (
	// "math" // for dot only
	// "math/rand"
	// "time"
)

type Element struct {
	Aka string // Name
	Next []string // Following attunement allowed
	Fractals []string // simple skills allowed for element
	// Native: native - penetrates easier, lucky effect
  // Makes stream unstable, but powerful: fuel - easiest penetration
} // Makes strem weak, but chills: chill - harder penetration
// else: no - regular penetration

var Elements, Physical []string = []string{"◌ ", "🌪 ", "🔥", "🪨", "🧊", "🌑", "🩸", "🎶", "☀️ "}, []string{"◌ ", "🌱", "🪵", "🪨", "🛡 "}
var ElemList, PhysList map[string]*Element = InitEpoch(1)

func (e *Element) AllowAttunes(nexts []string) { for _, each := range nexts  { (*e).Next = append((*e).Next, each) } }
func (e *Element) AllowFractal(frac string) { (*e).Fractals = append((*e).Fractals, frac) }

func InitEpoch(epoch int) (map[string]*Element, map[string]*Element) {
	// Prepared:
	var elements map[string]*Element = map[string]*Element{
		// Epoch 1: before Echaen's reserches: humans and feathers learn destruction
		Elements[0]: &Element{Aka: "Common"},
		Elements[1]: &Element{Aka: "Air"},
		Elements[2]: &Element{Aka: "Fire"},
		// Epoch 2: +alt +cre - less focus on element, more schools, +fractals series: x2 complexity
		Elements[3]: &Element{Aka: "Earth"},
		Elements[4]: &Element{Aka: "Water"},
		// Epoch 4: enslaved wild shadows learn to live with humans: x3 complexity
		Elements[5]: &Element{Aka: "Void"},
		Elements[6]: &Element{Aka: "Mallom"},
		// Epoch 8: mind controllers on 4, annihilators on 1
		Elements[7]: &Element{Aka: "Noise"},
		Elements[8]: &Element{Aka: "Resonance"},
	}
	var physicals map[string]*Element = map[string]*Element{
		// Sandy, ooze, flaiming or whirling spirits
		Physical[0]: &Element{Aka: "Spirit"},
		// Living creatures:
		Physical[1]: &Element{Aka: "Flesh/Plant"}, 
		Physical[2]: &Element{Aka: "Wood/Shell"}, 
		// Summoned and elementals:
		Physical[3]: &Element{Aka: "Hard"}, 
		// Plate armor and mechanisms
		Physical[4]: &Element{Aka: "Forged"}, 
	}


	// EPOCH 0: simplest energy jinxing - could not distinguish energy kinds: nature/element
	elements[Elements[0]].AllowFractal("Jinx")  // upgradable to any basic
	physicals[Physical[1]].AllowFractal("Punch")  // upgradable to any basic
	physicals[Physical[2]].AllowFractal("Stab")  // upgradable to any basic
	if epoch == 0 { return elements, physicals }


	// EPOCH 1: before Echaen: destruction mostly, 
	// air by feathers, fire by humans
	elements[Elements[0]].AllowAttunes(Elements[0:3])  // upgradable to any basic
	elements[Elements[1]].AllowAttunes(Elements[0:1])  // used by feathers
	elements[Elements[2]].AllowAttunes(Elements[0:1])  // used by humans
	elements[Elements[2]].AllowFractal("Cycle") // alt
	elements[Elements[2]].AllowFractal("Burn") // des 
	elements[Elements[1]].AllowFractal("Pressure") // des
	elements[Elements[1]].AllowFractal("Distribution") // alt
	if epoch == 1 { return elements, physicals }
	

	// EPOCH 2: schools
	// air by feathers, fire by humans
	elements[Elements[0]].AllowAttunes(Elements[3:5])  // upgradable to any basic
	elements[Elements[3]].AllowAttunes(Elements[0:1])  // used by feathers
	elements[Elements[4]].AllowAttunes(Elements[0:1])  // used by feathers
	elements[Elements[1]].AllowAttunes(Elements[1:2])  // used by feathers
	elements[Elements[2]].AllowAttunes(Elements[2:3])  // used by humans
	elements[Elements[3]].AllowAttunes(Elements[3:4])  // used by feathers
	elements[Elements[4]].AllowAttunes(Elements[4:5])  // used by feathers
	elements[Elements[2]].AllowAttunes(Elements[6:7])  // used by humans
	elements[Elements[4]].AllowAttunes(Elements[6:7])  // used by humans
	elements[Elements[6]].AllowAttunes(Elements[6:7])  // used by humans
	elements[Elements[6]].AllowAttunes(Elements[2:3])  // used by humans
	elements[Elements[6]].AllowAttunes(Elements[4:5])  // used by humans
	elements[Elements[6]].AllowAttunes(Elements[0:1])  // used by humans
	if epoch == 2 { return elements, physicals }


	// return LATEST EPOCH by defult
	return elements, physicals
}
package common

import (
	// "math" // for dot only
	// "math/rand"
	// "time"
)

type Element struct {
	Aka string
	//   ^ Name
	Eats, EatenBy, MadeOf, Next []string 
	// ^    ^        ^ Native: native - penetrates easier, lucky effect
  // ^    ^ Makes stream unstable, but powerful: fuel - easiest penetration
} // ^ Makes strem weak, but chills: chill - harder penetration
// else: no - regular penetration

var Elements []string = []string{"◌ ", "🌪 ", "🔥", "🪨", "🧊", "🌑", "🩸", "🎶", "☀️ "} 
var ElemList map[string]Element = map[string]Element{
	// Epoch 1: before Echaen's reserches: humans and feathers learn destruction
	Elements[0]: Element{Aka: "Common", EatenBy: []string{Elements[6]}, Next: Elements[0:3]},
	Elements[1]: Element{Aka: "Air",    Eats: []string{Elements[3]}, EatenBy: []string{Elements[2], Elements[7]}, Next: Elements[0:1]},
	Elements[2]: Element{Aka: "Fire",   Eats: []string{Elements[1]}, EatenBy: []string{Elements[4], Elements[6]}, Next: Elements[0:1]},
	// Epoch 2: +alt +cre - less focus on element, more schools, +fractals series
	Elements[3]: Element{Aka: "Earth",  Eats: []string{Elements[4]}, EatenBy: []string{Elements[1], Elements[7]}, Next: Elements[0:1]},
	Elements[4]: Element{Aka: "Water",  Eats: []string{Elements[2]}, EatenBy: []string{Elements[3], Elements[6]}, Next: Elements[0:1]},
	// Epoch 3: enslaved wild shadows learn to live with humans
	Elements[5]: Element{Aka: "Void",   Eats: Elements[:], MadeOf: []string{Elements[0]}},
	// Epoch 4: rhyxxix break out - ethernal shadows release
	Elements[6]: Element{Aka: "Mallom", Eats: []string{Elements[2], Elements[4]}, EatenBy: []string{Elements[5], Elements[2]}, MadeOf: []string{Elements[2], Elements[4]}},
	// Epoch 5: mind controllers on 4, annihilators from 1
	Elements[7]: Element{Aka: "Noise",  Eats: []string{Elements[1], Elements[3]}, EatenBy: []string{Elements[1], Elements[8]}, MadeOf: []string{Elements[1], Elements[3]}},
	Elements[8]: Element{Aka: "Resonance", Eats: Elements[1:5], EatenBy: Elements, MadeOf: []string{Elements[2]}},
}

var Physical []string = []string{"◌ ", "🌱", "🪵", "🪨", "🛡 "} // none, flesh/plant (punchD), wood/shell (stingA), stone (chopDa), armored
var PhysList map[string]Element = map[string]Element{
	// Sandy, ooze, flaiming or whirling spirits
	Physical[0]: Element{Aka: "Spirit"},
	// Living creatures:
	Physical[1]: Element{Aka: "Flesh/Plant"}, 
	Physical[2]: Element{Aka: "Wood/Shell"}, 
	// Summoned and elementals:
	Physical[3]: Element{Aka: "Hard",   MadeOf: []string{Physical[1]}}, 
	// Plate armor and mechanisms
	Physical[4]: Element{Aka: "Forged", MadeOf: []string{Physical[1], Physical[2]}}, 
}
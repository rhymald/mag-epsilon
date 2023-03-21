# Resources

## Basic:
```yaml
Stream:
  MadeOf:
    - Element: [Element, Physical]
    - StreamStats: [Creation, Alteration, Destruction]
  Examples: 
    - "Fire stream of 56cre, 27alt and 78des"
```
```yaml
StreamStats:
  MadeOf:
    - ElementToSwap
    - [3]StatsImprove
  Examples:
    - "New element mallom, with stats -17cre, -56alt, +71des"
```

## Atomic: 
```yaml
HP:
  MadeOf: 
    - CurrentAmount
  Examples: 
    - "Regenerating"
    - "Supressed"
    - "Dead"
```
```yaml
Heat:
  MadeOf: 
    - CurrentAmount
    - Stream.ID
  Examples:
    - "Calm and less efective"
    - "Warmed up and efective"
    - "Overheated and Unstable"
```
```yaml
XYZ:
  MadeOf: 
    - [3]Coordinates
    - []Affections
  Examples:
    - "Ashtown Hill: windy weather"
    - "Hatchery furnace: warm and cozy"
    - "Waterfall: wet and noisy"
```
```yaml
Meta:
  MadeOf:
    - ID: BornTime
    - LastTime: Stats last increase time
    - Chill: Time before next login / spawn
    - Busy: [true, false]
    - LoggedIn: [true, false]
    - NPC: [true, false]
    - Description: [ERROR, FlatText]
  Examples: 
    - Dummy NPC with stats X spawned in time before time-escape, then can be respawned after chill-tame
    - Player not NPC with stats born at born-time
    - Player not NPC with stats logged in, name "Mystical Adventurer"
```


## Consumables:
```yaml
Dot:
  MadeOf: 
    - Element: []All
    - Weight
  Examples: 
    - "Fire dot of 4"
    - "Void dot of 8"
    - "Common dot of 1"
```


## Transaction:
```yaml
Event: 
  MadeOf:
    - From: CharacterID
    - At: XYZ
    - ID: ActionID
    - Description: Some
    - Effect: complex
  Examples:
    - Heat (+heat)
    - Overheat (+damage yourself and friends)
    - Fed agressive stronger well (+well stats)
    - Fed native curse (+cuse stats)
    - Provoked carnivorous shadows (+provocation, +spawn)
    - Damage weak targets
    - Damage weak eatable wells
    - Emower native phenomenaes
    - Damage targets
    - Restoration
    - Degeneration
    - Barrier dome or buf
    - Curse or enchantment
    - ... many of them

```
```yaml
Action:
  MadeOf: 
    - From: CharacterID
    - Description: Name
    - To: XYZ
    - By: [ToolItems, Streams, Flocks, Fractals]
    - With: []Dots
    - ID: Time
    - Failed: [true, false]
  Examples: 
    - "You, id 1234-123456789-1-1234567"
    - "Cast summon the flame scheme"
    - "With wooden sword id 9"
    - "In vector, cone front of you at cordinates 23,4,67"
    - "Made of 3 fractals [fire, air, fire]"
    - "Using the flocks [summon, pressure, burn]"
    - "Used streams 12..13, 16..23. 14..19"
    - "Consumed dots 0..199"
    - [success, fail]
```


## Tools:
```yaml
Fractal:
  MadeOf:
    - Func: any cast or action
    - Flock: []streams, physical, XYZ
  Examples: 
    - Jinx: forJinx, forFire other
    - Pyrogramm: forPyro, forFire other
    - Windblow: forWind, forAir other
    - Step: direction
    - Punch: with physical
```


## Mutateble:
```yaml
Flock:
  MadeOf:
    - Streams: []Id
    - Next: Id
    - NextBefore: For ringrobin use
  Examples: 
    - "For windblow, streams 1, 4, 7 and 9 - last 7 for next 5 min, then 1"
```


---
# TBD
```yaml
Scheme:
  MadeOf:
    - [4]Fractals series
    - [3]Connections series
    - []Actions series
```
```yaml
Stamina:
  MadeOf: 
    - CurrentAmount
  Examples: 
    - "Exhausted"
    - "Full"
    - "Not enough"
```
```yaml
ConsumableItem:
  MadeOf:
    - Func
    - Effectiveness
  Examples: 
    - "Energy potion"
    - "Crafting resource"
    - Coins
```
```yaml
Item:
  MadeOf:
    - []Streams
    - Nature: [ c-well, d-curse, a-mod ]
    - ...TBD
  Examples:
    - Robe
    - "Cursed ring"
    - "Bag of fire salt"
```
```yaml
ToolItem:
  MadeOf:
    - []Fractals
    - []Actions
    - InUse: [ true, false ]
    - ...TBD
  Examples:
    - "Concentrator"
    - "Energy capacitor"
    - "Gatherring tool"
```
```yaml
Well:
```
```yaml
Curse:
```
```yaml
Effect:
```
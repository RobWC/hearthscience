package hearthscience

import (
	hs "github.com/robwc/go-hearthstone"
)

type Minion struct {
	Name               string   `json:"name"`
	Id                 string   `json:"id"`
	Attack             int      `json:"attack"`
	Health             int      `json:"health"`
	Mechanics          []string `json:"mechanics"`
	Cost               int      `json:"cost"`
	Collectable        bool     `json:"collectible"`
	AttackToCost       float32
	HealthToCost       float32
	AttackHealthToCost float32
	AttackHealthRatio  float32
	AttackHealthTotal  float32
	MechanicsValue     float32
}

func calculateMinionValue(card hs.Card) *Minion {
	cv := &Minion{Collectable: card.Collectable, Id: card.Id, Name: card.Name, Attack: card.Attack, Health: card.Health, Mechanics: card.Mechanics, Cost: card.Cost}
	//AttackHealthRatio
	cv.AttackHealthRatio = float32(card.Attack) / float32(card.Health)
	//AttackHealthTotal
	cv.AttackHealthTotal = float32(card.Attack) + float32(card.Health)
	//MechanicsValue
	var mechValue float32
	mechValue = 0
	if len(card.Mechanics) > 0 {
		for item := range card.Mechanics {
			switch card.Mechanics[item] {
			case "Battlecry":
				mechValue = mechValue + BattlecryValue
			case "Taunt":
				mechValue = mechValue + TauntValue
			case "Deathrattle":
				mechValue = mechValue + DeathrattleValue
			case "Spellpower":
				mechValue = mechValue + SpellpowerValue
			case "Charge":
				mechValue = mechValue + ChargeValue
			case "Stealth":
				mechValue = mechValue + StealthValue
			case "Windfury":
				mechValue = mechValue + WindfuryValue
			case "Combo":
				mechValue = mechValue + ComboValue
			case "Aura":
				mechValue = mechValue + AuraValue
			case "Divine Shield":
				mechValue = mechValue + DivineShieldValue
			}
		}
	}
	cv.MechanicsValue = mechValue
	//AttackHealthToCost
	cv.AttackHealthToCost = cv.AttackHealthTotal / float32(card.Cost)
	//HealthToCost
	cv.HealthToCost = float32(card.Health) / float32(card.Cost)
	//AttackToCost
	cv.AttackToCost = float32(card.Attack) / float32(card.Cost)
	return cv
}

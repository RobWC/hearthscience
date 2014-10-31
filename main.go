package hearthscience

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	hs "github.com/robwc/go-hearthstone"

	"appengine"
	"appengine/datastore"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler).Methods("GET")
	r.HandleFunc("/update", updateHandler)
	r.HandleFunc("/cards", cardList)
	r.HandleFunc("/cardlist", allCards)
	r.HandleFunc("/card/{id:[0-9a-zA-Z]+_[0-9a-zA-Z]+}", cardInfo)
	http.Handle("/", r)
}

func allCards(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("hs.Card").Filter("Collectible =", true).Filter("Type =", "Minion").Filter("Type =", "Minion").Order("-Attack").Order("Name")
	var cards []hs.Card
	if _, err := q.GetAll(c, &cards); err != nil {
		c.Errorf(err.Error())
		return
	}
	q2 := datastore.NewQuery("hs.Card").Order("Name").Filter("Type =", "Spell")
	if _, err := q2.GetAll(c, &cards); err != nil {
		c.Errorf(err.Error())
		return
	}
	fmt.Fprintf(w, "<html><body>")
	for card := range cards {
		fmt.Fprintf(w, "Name: %s Cost: %d Link: <a href=\"/card/%s\">LINK</a><br>", cards[card].Name, cards[card].Cost, cards[card].Id)
	}
	fmt.Fprintf(w, "</body></html>")
}

var (
	BattlecryValue float32 = 4
	TauntValue float32 = 3
	DeathrattleValue float32 = 4
	SpellpowerValue float32 = 4
	ChargeValue float32 = 4
	StealthValue float32 = 4
	WindfuryValue float32 = 4
	ComboValue float32 = 2
	DivineShieldValue float32 = 2
	FreezeValue float32 = 3
	SecretValue float32 = 2
	OneTurnEffectVale float32 = 1
	AuraValue float32 = 2
)

type MinionValue struct {
	AttackToCost float32
	HealthToCost float32
	AttackHealthToCost float32
	AttackHealthRatio float32
	AttackHealthTotal float32
	MechanicsValue float32
}

func calculateMinionValue(card hs.Card) *MinionValue {
	cv := &MinionValue{}
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
				case "Battlecry": mechValue = mechValue + BattlecryValue
				case "Taunt": mechValue = mechValue + TauntValue
				case "Deathrattle": mechValue = mechValue + DeathrattleValue
				case "Spellpower": mechValue = mechValue + SpellpowerValue
				case "Charge": mechValue = mechValue + ChargeValue
				case "Stealth": mechValue = mechValue + StealthValue
				case "Windfury": mechValue = mechValue + WindfuryValue
				case "Combo": mechValue = mechValue + ComboValue
				case "Aura": mechValue = mechValue + AuraValue
				case "Divine Shield": mechValue = mechValue + DivineShieldValue
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

func cardInfo(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	var fetchedCards []hs.Card
	vars := mux.Vars(r)
	id := vars["id"]

	cardAsset, _ := Asset("templates/card.tmpl")
	headerAsset, _ := Asset("templates/header.tmpl")
	footerAsset, _ := Asset("templates/footer.tmpl")
	testmathAsset, _ := Asset("templates/testmath.tmpl")

	finalTemplate := template.New("finalCard")

	finalTemplate.Parse(string(testmathAsset))
	finalTemplate.Parse(string(headerAsset))
	finalTemplate.Parse(string(footerAsset))
	finalTemplate.Parse(string(cardAsset))

	q := datastore.NewQuery("hs.Card").Filter("Id =", id)
	if _, err := q.GetAll(c, &fetchedCards); err != nil {
		c.Errorf(err.Error())
	}
	if len(fetchedCards) > 0 {
		finalTemplate.ExecuteTemplate(w, "header", fetchedCards[0])
		finalTemplate.ExecuteTemplate(w, "card", fetchedCards[0])
		finalTemplate.ExecuteTemplate(w, "testmath", calculateMinionValue(fetchedCards[0]))
		finalTemplate.ExecuteTemplate(w, "footer", nil)
	} else {
		fmt.Fprintf(w, "NONE")
		finalTemplate.Execute(w, nil)
	}
	//
}

func cardList(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("hs.Card").Filter("Type =", "Minion").Filter("Type =", "Minion").Order("Name")
	var cards []hs.Card
	if _, err := q.GetAll(c, &cards); err != nil {
		c.Errorf(err.Error())
		return
	}
	q2 := datastore.NewQuery("hs.Card").Order("Name").Filter("Type =", "Spell")
	if _, err := q2.GetAll(c, &cards); err != nil {
		c.Errorf(err.Error())
		return
	}
	fmt.Fprintf(w, "<html><body>")
	for card := range cards {
		fmt.Fprintf(w, "Name: %s Cost: %d<br>", cards[card].Name, cards[card].Cost)
	}
	fmt.Fprintf(w, "</body></html>")
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	input, err := ioutil.ReadFile("/home/rcameron/gopath/src/github.com/robwc/hearthscience.com/input_files/AllSets.json")
	if err != nil {
		c.Errorf(err.Error())
	}
	//cardList := make(map[string][]hs.Card)
	cardSet := make(map[string][]hs.Card)
	cardReader := bytes.NewReader(input)
	dec := json.NewDecoder(cardReader)
	for {
		if err := dec.Decode(&cardSet); err == io.EOF {
			break
		} else if err != nil {
			c.Errorf(err.Error())
		}
		c.Infof(strconv.Itoa(len(cardSet)))

	}
	for set := range cardSet {
		c.Infof(strconv.Itoa(len(cardSet[set])))
		for item := range cardSet[set] {
			var card hs.Card
			card = cardSet[set][item]
			key := datastore.NewIncompleteKey(c, "hs.Card", nil)
			_, err := datastore.Put(c, key, &card)
			if err != nil {
				c.Errorf(err.Error())
			}
		}
	}
	c.Infof("Length: " + strconv.Itoa((len(input))))
	fmt.Fprintf(w, "Import")
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	//key := datastore.NewIncompleteKey(c, "hs.Card", nil)
	//_, err := datastore.Put(c, key, goo)
	//if err != nil {
	//	fmt.Println(err)
	//}
	fmt.Fprintf(w, "Hello Stone!")
}

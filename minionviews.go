package hearthscience

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	hs "github.com/robwc/go-hearthstone"
	"gopkg.in/mgo.v2/bson"
)

func allMinions(w http.ResponseWriter, r *http.Request) {
	count := 0
	cardsConn := mongoSession.DB("").C("cards")
	var cards []hs.Card
	var err = cardsConn.Find(bson.M{"collectible": true, "type": "Minion"}).Sort("name").All(&cards)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	fmt.Fprintf(w, "<html><body>")
	for card := range cards {
		fmt.Fprintf(w, "Count: %d Name: %s Cost: %d Link: <a href=\"/minion/%s\">LINK</a><br>", count, cards[card].Name, cards[card].Cost, cards[card].Id)
		count = count + 1
	}
	fmt.Fprintf(w, "</body></html>")
}

func allMinionsJSONv1(w http.ResponseWriter, r *http.Request) {
cardsConn := mongoSession.DB("").C("cards")
var cards []Minion
var err = cardsConn.Find(bson.M{"collectible": true, "type": "Minion"}).Sort("name").All(&cards)
if err != nil {
	log.Printf(err.Error())
	return
}
newJSON, _ := json.Marshal(cards)
fmt.Fprintf(w, "{\"total\":"+strconv.Itoa(len(cards))+" ,\"minions\":"+string(newJSON)+"}")
}


func minionInfo(w http.ResponseWriter, r *http.Request) {

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

	cardsConn := mongoSession.DB("").C("cards")
	card := hs.Card{}
	var err = cardsConn.Find(bson.M{"id": id}).One(&card)
	if err != nil {
		log.Printf(err.Error())
		fmt.Fprintf(w, "NONE")
		finalTemplate.Execute(w, nil)
	}
	finalTemplate.ExecuteTemplate(w, "header", card)
	finalTemplate.ExecuteTemplate(w, "card", card)
	finalTemplate.ExecuteTemplate(w, "testmath", calculateMinionValue(card))
	finalTemplate.ExecuteTemplate(w, "footer", nil)
}

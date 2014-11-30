package hearthscience

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	hs "github.com/robwc/go-hearthstone"
	"gopkg.in/mgo.v2/bson"
)

func allMechanics(w http.ResponseWriter, r *http.Request) {
	cardsConn := mongoSession.DB("").C("cards")
	var mechanics []string
	var err = cardsConn.Find(bson.M{"collectible": true, "type": "Minion"}).Distinct("mechanics", &mechanics)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	sort.Strings(mechanics)
	fmt.Fprintf(w, "<html><body>")
	for mechanic := range mechanics {
		if mechanics[mechanic] != "" {
			fmt.Fprintf(w, "Mechanic: %s Link: <a href=\"/mechanic/%s\">LINK</a><br>", mechanics[mechanic], mechanics[mechanic])
		}
	}
	fmt.Fprintf(w, "</body></html>")
}

func mechanicsCards(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mechanic := vars["mechanic"]

	count := 0
	cardsConn := mongoSession.DB("").C("cards")
	var cards []hs.Card
	var err = cardsConn.Find(bson.M{"mechanics": mechanic}).Sort("name").All(&cards)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	fmt.Fprintf(w, "<html><body>")
	for card := range cards {
		fmt.Fprintf(w, "Count: %d Name: %s Cost: %d Link: <a href=\"/card/%s\">LINK</a><br>", count, cards[card].Name, cards[card].Cost, cards[card].Id)
		count = count + 1
	}
	fmt.Fprintf(w, "</body></html>")
}

func allMechanicsJSONv1(w http.ResponseWriter, r *http.Request) {
	cardsConn := mongoSession.DB("").C("cards")
	var mechanics []string
	var err = cardsConn.Find(bson.M{"collectible": true, "type": "Minion"}).Distinct("mechanics", &mechanics)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	sort.Strings(mechanics)
	for mechanic := range mechanics {
		fmt.Println(mechanics[mechanic])
		if mechanics[mechanic] == "" {
			mechanics = mechanics[:mechanic+copy(mechanics[mechanic:], mechanics[mechanic+1:])]
			break
		}
	}
	newJSON, _ := json.Marshal(mechanics)
	fmt.Fprintf(w, "{\"total\":"+strconv.Itoa(len(mechanics))+" ,\"mechanics\":"+string(newJSON)+"}")
}

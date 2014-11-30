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

func allRaces(w http.ResponseWriter, r *http.Request) {
	cardsConn := mongoSession.DB("").C("cards")
	var races []string
	var err = cardsConn.Find(bson.M{"collectible": true, "type": "Minion"}).Distinct("race", &races)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	sort.Strings(races)
	fmt.Fprintf(w, "<html><body>")
	for race := range races {
		if races[race] != "" {
			fmt.Fprintf(w, "Race: %s Link: <a href=\"/race/%s\">LINK</a><br>", races[race], races[race])
		}
	}
	fmt.Fprintf(w, "</body></html>")
}

func raceCards(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	race := vars["race"]

	count := 0
	cardsConn := mongoSession.DB("").C("cards")
	var cards []hs.Card
	var err = cardsConn.Find(bson.M{"race": race}).Sort("name").All(&cards)
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

func allRacesJSONv1(w http.ResponseWriter, r *http.Request) {
	cardsConn := mongoSession.DB("").C("cards")
	var races []string
	var err = cardsConn.Find(bson.M{"collectible": true, "type": "Minion"}).Distinct("race", &races)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	sort.Strings(races)
	for race := range races {
		fmt.Println(races[race])
		if races[race] == "" {
			races = races[:race+copy(races[race:], races[race+1:])]
			break
		}
	}
	newJSON, _ := json.Marshal(races)
	fmt.Fprintf(w, "{\"total\":"+strconv.Itoa(len(races))+" ,\"races\":"+string(newJSON)+"}")
}

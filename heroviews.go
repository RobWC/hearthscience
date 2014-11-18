package hearthscience

import (
	"fmt"
	"log"
	"net/http"

	hs "github.com/robwc/go-hearthstone"
	"gopkg.in/mgo.v2/bson"
)

func allHeros(w http.ResponseWriter, r *http.Request) {
	var card hs.Card
	fmt.Println(card)
	count := 0
	cardsConn := mongoSession.DB("").C("cards")
	var cards []hs.Card
	var err = cardsConn.Find(bson.M{"collectible": true, "type": "Hero"}).Sort("name").All(&cards)
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

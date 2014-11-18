package hearthscience

import (
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

var mongoSession *mgo.Session

func Main() {
	//Create mongodb connection
	var err error
	mongoSession, err = mgo.Dial("localhost/hearthscience")
	if err != nil {
		log.Printf(err.Error())
	}
	defer mongoSession.Close()
	mongoSession.SetMode(mgo.Monotonic, true)
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler).Methods("GET")
	r.HandleFunc("/update", updateHandler).Methods("GET")
	r.HandleFunc("/heros", allHeros).Methods("GET")
	r.HandleFunc("/heropowers", allHeroPowers).Methods("GET")
	r.HandleFunc("/weapons", allWeapons).Methods("GET")
	r.HandleFunc("/spells", allSpells).Methods("GET")
	r.HandleFunc("/minions", allMinions).Methods("GET")
	r.HandleFunc("/minion/{id:[0-9a-zA-Z]+_[0-9a-zA-Z]+}", minionInfo).Methods("GET")
	r.HandleFunc("/cards", allCards).Methods("GET")
	r.HandleFunc("/card/{id:[0-9a-zA-Z]+_[0-9a-zA-Z]+}", cardInfo).Methods("GET")
	r.HandleFunc("/rarity/{rarity:[A-Za-z]+}", cardInfo).Methods("GET")
	r.HandleFunc("/cost/{cost:[0-9]+}", cardInfo).Methods("GET")
	r.HandleFunc("/attack/{attack:[0-9]+}", cardInfo).Methods("GET")
	r.HandleFunc("/health/{health:[0-9]+}", cardInfo).Methods("GET")
	r.HandleFunc("/mechanic/{mechanic:[0-9]+}", cardInfo).Methods("GET")
	r.HandleFunc("/race/{race:[A-Za-z]+}", cardInfo).Methods("GET")
	//JSON APIs
	r.HandleFunc("/v1/cards", allCardsJSONv1).Methods("GET").Queries("type", "json")
	r.HandleFunc("/v1/minions", allMinionsJSONv1).Methods("GET").Queries("type", "json")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func cardList(w http.ResponseWriter, r *http.Request) {
	/*
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
	*/
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href=\"/cards\">Cards</a><br>")
	fmt.Fprintf(w, "<a href=\"/minions\">Minions</a><br>")
	fmt.Fprintf(w, "<a href=\"/spells\">Spells</a><br>")
	fmt.Fprintf(w, "<a href=\"/weapons\">Weapons</a><br>")
	fmt.Fprintf(w, "<a href=\"/heros\">Heros</a><br>")
	fmt.Fprintf(w, "<a href=\"/heropowers\">Hero Powers</a><br>")
}

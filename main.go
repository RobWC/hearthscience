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
	r.HandleFunc("/js/d3.js", jsD3Handler).Methods("GET")
	r.HandleFunc("/js/jquery.js", jsJQueryHandler).Methods("GET")
	r.HandleFunc("/update", updateHandler).Methods("GET")
	r.HandleFunc("/heros", allHeros).Methods("GET")
	r.HandleFunc("/heropowers", allHeroPowers).Methods("GET")
	r.HandleFunc("/weapons", allWeapons).Methods("GET")
	r.HandleFunc("/spells", allSpells).Methods("GET")
	r.HandleFunc("/minions", allMinions).Methods("GET")
	r.HandleFunc("/minion/{id:[0-9a-zA-Z]+_[0-9a-zA-Z]+}", minionInfo).Methods("GET")
	r.HandleFunc("/cards", allCards).Methods("GET")
	r.HandleFunc("/card/{id:[0-9a-zA-Z]+_[0-9a-zA-Z]+}", cardInfo).Methods("GET")
	r.HandleFunc("/rarities", allRarities).Methods("GET")
	r.HandleFunc("/rarity/{rarity:[A-Za-z]+}", rarityCards).Methods("GET")
	r.HandleFunc("/cost/{cost:[0-9]+}", costCards).Methods("GET")
	r.HandleFunc("/attack/{attack:[0-9]+}", attackCards).Methods("GET")
	r.HandleFunc("/health/{health:[0-9]+}", healthCards).Methods("GET")
	r.HandleFunc("/mechanics", allMechanics).Methods("GET")
	r.HandleFunc("/mechanic/{mechanic:[0-9a-zA-Z ]+}", mechanicsCards).Methods("GET")
	r.HandleFunc("/races", allRaces).Methods("GET")
	r.HandleFunc("/race/{race:[A-Za-z]+}", raceCards).Methods("GET")
	//JSON APIs
	r.HandleFunc("/v1/cards", allCardsJSONv1).Methods("GET").Queries("type", "json")
	r.HandleFunc("/v1/minions", allMinionsJSONv1).Methods("GET").Queries("type", "json")
	r.HandleFunc("/v1/races", allRacesJSONv1).Methods("GET").Queries("type", "json")
	r.HandleFunc("/v1/mechanics", allMechanicsJSONv1).Methods("GET").Queries("type", "json")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href=\"/cards\">Cards</a><br>")
	fmt.Fprintf(w, "<a href=\"/minions\">Minions</a><br>")
	fmt.Fprintf(w, "<a href=\"/spells\">Spells</a><br>")
	fmt.Fprintf(w, "<a href=\"/weapons\">Weapons</a><br>")
	fmt.Fprintf(w, "<a href=\"/heros\">Heros</a><br>")
	fmt.Fprintf(w, "<a href=\"/races\">Races</a><br>")
	fmt.Fprintf(w, "<a href=\"/rarities\">Rarities</a><br>")
	fmt.Fprintf(w, "<a href=\"/mechanics\">Mechanics</a><br>")
	fmt.Fprintf(w, "<a href=\"/heropowers\">Hero Powers</a><br>")
}

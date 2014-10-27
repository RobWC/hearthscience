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
	r.HandleFunc("/cardlist",allCards)
	r.HandleFunc("/card/{id:[0-9a-zA-Z]+_[0-9a-zA-Z]+}", cardInfo)
	http.Handle("/", r)
}

func allCards(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("hs.Card").Filter("Type =", "Minion").Filter("Type =", "Minion").Order("-Attack").Order("Name")
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
		fmt.Fprintf(w, "Name: %s Cost: %d Link: <a href=\"/card/%s\">LINK</a><br>", cards[card].Name, cards[card].Cost,cards[card].Id)
	}
	fmt.Fprintf(w, "</body></html>")
}

func cardInfo(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	var fetchedCards []hs.Card
	vars := mux.Vars(r)
	id := vars["id"]

	finalTemplate, err := template.ParseFiles("templates/header.tmpl", "templates/card.tmpl", "templates/footer.tmpl")
	if err != nil {
		c.Errorf(err.Error())
	}

	q := datastore.NewQuery("hs.Card").Filter("Id =", id)
	if _, err := q.GetAll(c, &fetchedCards); err != nil {
		c.Errorf(err.Error())
	}
	if len(fetchedCards) > 0 {
		finalTemplate.ExecuteTemplate(w,"header",fetchedCards[0])
		finalTemplate.ExecuteTemplate(w,"card",fetchedCards[0])
		finalTemplate.ExecuteTemplate(w,"footer",nil)
	} else {
		fmt.Fprintf(w, "NONE")
	}
	finalTemplate.Execute(w,nil)
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

package hearthscience

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	hs "github.com/robwc/go-hearthstone"
)

func updateHandler(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadFile("/home/rcameron/gopath/src/github.com/robwc/hearthscience.com/input_files/AllSets.json")
	if err != nil {
		log.Printf(err.Error())
	}

	cardSet := make(map[string][]hs.Card)
	cardReader := bytes.NewReader(input)
	dec := json.NewDecoder(cardReader)
	for {
		if err := dec.Decode(&cardSet); err == io.EOF {
			break
		} else if err != nil {
			log.Printf(err.Error())
		}
		log.Printf(strconv.Itoa(len(cardSet)))

	}
	for set := range cardSet {
		log.Printf(strconv.Itoa(len(cardSet[set])))
		cardsConn := mongoSession.DB("").C("cards")
		minionConn := mongoSession.DB("").C("minions")
		for item := range cardSet[set] {
			var card hs.Card
			card = cardSet[set][item]
			var err = cardsConn.Insert(card)
			if err != nil {
				log.Printf(err.Error())
			}
			if card.Type == "Minion" {
				minionCard := calculateMinionValue(card)
				var err = minionConn.Insert(minionCard)
				if err != nil {
					log.Printf(err.Error())
				}
			}
		}
	}
	log.Printf("Length: " + strconv.Itoa((len(input))))
	fmt.Fprintf(w, "Import")
}

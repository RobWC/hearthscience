package hearthscience

import (
    "fmt"
    "log"
    "net/http"
    "sort"

    "github.com/gorilla/mux"
    hs "github.com/robwc/go-hearthstone"
    "gopkg.in/mgo.v2/bson"
)

func allRarities(w http.ResponseWriter, r *http.Request) {
    cardsConn := mongoSession.DB("").C("cards")
    var rarities []string
    var err = cardsConn.Find(bson.M{"collectible": true}).Distinct("rarity", &rarities)
    if err != nil {
        log.Printf(err.Error())
        return
    }
    sort.Strings(rarities)
    fmt.Fprintf(w, "<html><body>")
    for rarity := range rarities {
        if rarities[rarity] != "" {
            fmt.Fprintf(w, "Rarity: %s Link: <a href=\"/rarity/%s\">LINK</a><br>", rarities[rarity], rarities[rarity])
        }
    }
    fmt.Fprintf(w, "</body></html>")
}

func rarityCards(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    rarity := vars["rarity"]

    count := 0
    cardsConn := mongoSession.DB("").C("cards")
    var cards []hs.Card
    var err = cardsConn.Find(bson.M{"rarity": rarity}).Sort("name").All(&cards)
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

/*
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
*/

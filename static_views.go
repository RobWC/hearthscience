package hearthscience

import (
	"fmt"
	"net/http"
)

func jsD3Handler(w http.ResponseWriter, r *http.Request) {
	jsFile, _ := Asset("templates/d3.js")
	w.Header().Set("Content-Type", "application/javascript")
	fmt.Fprintf(w, string(jsFile[:len(jsFile)]))
}

func jsJQueryHandler(w http.ResponseWriter, r *http.Request) {
	jsFile, _ := Asset("templates/jquery-1.11.1.min.js")
	w.Header().Set("Content-Type", "application/javascript")
	fmt.Fprintf(w, string(jsFile[:len(jsFile)]))
}

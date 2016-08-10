package handlers

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dude, this is the PokeIntel backend.\n") // Joke (c) Niantic Labs https://pgorelease.nianticlabs.com/plfe/
}

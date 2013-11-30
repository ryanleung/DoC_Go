package controllers

import (
	"fmt"
	"net/http"
	// "encoding/json"
	"deck_service/base"
)

type JsonService struct {
	GameMgr *base.GameManager
}

func (js *JsonService) Serve(port int) {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "hI!")
		})

	http.HandleFunc(
		"/new_room",
		func(w http.ResponseWriter, r *http.Request) {
			handleAddNewRoom(w, r, js.GameMgr)
		})
	http.HandleFunc(
		"/renew_deck",
		func(w http.ResponseWriter, r *http.Request) {
			handleRenewDeck(w, r, js.GameMgr)
		})

	http.HandleFunc(
		"/draw_card",
		func(w http.ResponseWriter, r *http.Request) {
			handleDrawCard(w, r, js.GameMgr)
		})

	http.HandleFunc(
		"/add_player",
		func (w http.ResponseWriter, r *http.Request) {
			handleAddPlayer(w, r, js.GameMgr)
		})

	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func handleAddNewRoom(w http.ResponseWriter, r *http.Request, gm *base.GameManager) {
	name := r.FormValue("roomName")
	gm.AddNewRoom(name)
	fmt.Fprintf(w, "Rooms %v", gm.Rooms)
}

func handleRenewDeck(w http.ResponseWriter, r *http.Request, gm *base.GameManager) {
	name := r.FormValue("roomName")
	gm.Rooms[name].RenewDeck()
	fmt.Fprintf(w, "The deck %v", gm.Rooms[name].Deck)
}
func handleDrawCard(w http.ResponseWriter, r *http.Request, gm *base.GameManager) {
	name := r.FormValue("roomName")
	fmt.Fprintf(w, "Card drawn %v", gm.Rooms[name].DrawCard())
}

func handleAddPlayer(w http.ResponseWriter, r *http.Request, gm *base.GameManager) {
	playerName := r.FormValue("playerName")
	roomName := r.FormValue("roomName")
	gm.Rooms[roomName].AddPlayer(playerName)
	fmt.Fprintf(w, "Players %v", gm.Rooms[roomName].Players)
}



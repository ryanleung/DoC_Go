package base

import (
	"fmt"
	"math/rand"
	"time"
)

//////////////////////// INITIALIZERS/TYPES ////////////////////////

// A card is represented by the first letter of the ranks
// and the first letter of the suit. ex => Ace Spades => As
type Card string

type Deck []Card

type Player struct {
	name string
	cards []Card
}

type Board []Card

type Room struct {
	Name string

	// name of player => Player
	Players map[string]*Player
	Board Board
	Deck Deck
}

type GameManager struct {

	// name of room => Room
	Rooms map[string]*Room
}

func NewGameManager() *GameManager {
	gm := &GameManager{
		Rooms: make(map[string]*Room),
	}

	// random seed used for drawing cards
	rand.Seed( time.Now().UTC().UnixNano())

	return gm
}

func (gm *GameManager) AddNewRoom(name string) {
	newRoom := &Room{
		Name: name,
		Players:	make(map[string]*Player),
		Board:		make([]Card, 0),
		Deck:		make([]Card, 52),
	}

	newRoom.RenewDeck()
	gm.Rooms[name] = newRoom
}

//////////////////////// ROOM METHODS ////////////////////////

func (room *Room) RenewDeck() {
	room.Deck = []Card{}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suits := []string{"d", "c", "h", "s"}
	for i := 0; i < len(ranks); i++ {
		for j := 0; j < len(suits); j++ {
			room.Deck = append(room.Deck, Card(fmt.Sprintf("%v%v", ranks[i], suits[j])))
		}
	}
}

func (room *Room) DrawCard() Card {
	if len(room.Deck) <= 0 {
		return Card("")
	} else {
		drawingIndex := rand.Intn(len(room.Deck))
		cardToReturn := room.Deck[drawingIndex]
		
		// delete card
		room.Deck[drawingIndex] = room.Deck[len(room.Deck)-1]
		room.Deck = room.Deck[:len(room.Deck)-1]
		return cardToReturn
	}
}

func (room *Room) AddPlayer(name string) {
	newPlayer := &Player{
		name: name,
		cards: make([]Card, 0),
	}
	room.Players[name] = newPlayer
}
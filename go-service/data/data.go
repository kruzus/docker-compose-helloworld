package data

//Lets assume we are getting simple Player data from our fake DB.
//It returns simple JSON data with player object in it.

//Player object.
type Player struct {
	Name   string `json:"name"`
	Level  uint8  `json:"level"`
	Class  string `json:"class"`
	Active bool   `json:"status"`
}

var Bob = Player{
	Name:   "Bob",
	Level:  40,
	Class:  "Warrior",
	Active: true,
}

var Steve = Player{
	Name:   "Steve",
	Level:  80,
	Class:  "Paladin",
	Active: false,
}

var Elise = Player{
	Name:   "Elise",
	Level:  55,
	Class:  "Rogue",
	Active: true,
}
var Maria = Player{
	Name:   "Maria",
	Level:  80,
	Class:  "Warlock",
	Active: true,
}

type Players []*Player

func (p *Players) NewList(players ...*Player) Players {
	var data Players
	data = append(data, players...)
	return data
}

var AllPlayers Players

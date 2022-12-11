package metro

import "fmt"

type Inventory struct {
	passType map[string]PassengerType
	cards    map[string]*Card
	stations map[string]*Station
}

func NewInventory() *Inventory {
	inv := &Inventory{
		passType: make(map[string]PassengerType),
		cards:    make(map[string]*Card),
		stations: make(map[string]*Station),
	}
	inv.AddDefaultStations()
	inv.AddPassengerTypes()
	return inv
}

func (i *Inventory) AddPassengerTypes() {
	adult := Adult{}
	i.passType[adult.Type()] = adult
	senior := SeniorCitizen{}
	i.passType[senior.Type()] = senior
	kid := Kid{}
	i.passType[kid.Type()] = kid
}

func (i *Inventory) AddDefaultStations() {
	airport := &Station{
		Name:              "AIRPORT",
		ServiceFeePercent: 2,
		ReturnDiscount:    50,
	}
	i.stations[airport.Name] = airport

	central := &Station{
		Name:              "CENTRAL",
		ServiceFeePercent: 2,
		ReturnDiscount:    50,
	}
	i.stations[central.Name] = central
}

func (i *Inventory) AddCard(card *Card) {
	i.cards[card.ID] = card
}

func (i *Inventory) GetCard(id string) *Card {
	return i.cards[id]
}

func (i *Inventory) GetStation(id string) *Station {
	return i.stations[id]
}

func (i *Inventory) GetToStation(from string) *Station {
	if from == "AIRPORT" {
		return i.stations["CENTRAL"]
	}
	if from == "CENTRAL" {
		return i.stations["AIRPORT"]
	}
	panic(fmt.Sprintf("unknown station: %s", from))
}

func (i *Inventory) GetPassType(typeID string) PassengerType {
	return i.passType[typeID]
}

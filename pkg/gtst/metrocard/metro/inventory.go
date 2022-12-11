package metro

import "fmt"

type Inventory struct {
	passType map[string]PassengerType
	cards    map[string]*TravelCard
	stations map[string]*Station
}

func NewInventory() *Inventory {
	inv := &Inventory{
		passType: make(map[string]PassengerType),
		cards:    make(map[string]*TravelCard),
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
	discounts := make([]RideDiscount, 0)
	discounts = append(discounts, NewReturnJourneyDiscount(50.0))

	// airport := &Station{
	// 	Name:              "AIRPORT",
	// 	ServiceFeePercent: 2,
	// 	ReturnDiscount:    50,
	// }
	// i.stations[airport.Name] = airport
	i.stations["AIRPORT"] = NewStation("AIRPORT", discounts)

	// central := &Station{
	// 	Name:              "CENTRAL",
	// 	ServiceFeePercent: 2,
	// 	ReturnDiscount:    50,
	// }
	i.stations["CENTRAL"] = NewStation("CENTRAL", discounts)
}

func (i *Inventory) AddCard(card *TravelCard) {
	i.cards[card.ID] = card
}

func (i *Inventory) GetCard(id string) *TravelCard {
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

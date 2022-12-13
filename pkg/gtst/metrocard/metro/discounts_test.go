package metro

import (
	"fmt"
	"testing"
)

func TestReturnJourneyDiscount(t *testing.T) {
	inv := NewInventory()
	card1 := NewTravelCard("MC1", 500)
	inv.AddCard(card1)
	retDisc := NewReturnJourneyDiscount(50)
	passType := inv.GetPassType("ADULT")
	tr1 := &TravelDetail{
		Card:        card1,
		FromStation: "AIRPORT",
		PassType:    passType.Type(),
		Charge:      passType.JourneyRate(),
	}
	card1.InitiateNewJourney("AIRPORT", "CENTRAL")
	retDisc.Apply(tr1)
	card1.CompleteJourney()

	fmt.Printf("travel details after round1: %v\n", tr1)

	tr2 := &TravelDetail{
		Card:        card1,
		FromStation: "CENTRAL",
		PassType:    passType.Type(),
		Charge:      passType.JourneyRate(),
	}
	card1.InitiateNewJourney("CENTRAL", "AIRPORT")
	retDisc.Apply(tr2)
	card1.CompleteJourney()

	fmt.Printf("travel details after round2: %v\n", tr2)

	card1.InitiateNewJourney("AIRPORT", "CENTRAL")
	tr3 := &TravelDetail{
		Card:        card1,
		FromStation: "AIRPORT",
		PassType:    passType.Type(),
		Charge:      passType.JourneyRate(),
	}
	retDisc.Apply(tr3)

	fmt.Printf("travel details after round3: %v\n", tr3)
}

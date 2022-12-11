package metro

import "fmt"

type Journey struct {
	FromStn  string
	ToStn    string
	IsReturn bool
}

func NewJourney(from, to string) *Journey {
	return &Journey{
		FromStn: from,
		ToStn:   to,
	}
}

type PassengerType interface {
	Type() string
	JourneyRate() float64
}

type Adult struct{}

func (a Adult) Type() string {
	return "ADULT"
}

func (a Adult) JourneyRate() float64 {
	return 200.0
}

type SeniorCitizen struct{}

func (s SeniorCitizen) Type() string {
	return "SENIOR_CITIZEN"
}

func (s SeniorCitizen) JourneyRate() float64 {
	return 100.0
}

type Kid struct{}

func (s Kid) Type() string {
	return "KID"
}

func (s Kid) JourneyRate() float64 {
	return 50.0
}

type TravelCard struct {
	ID             string
	balance        float64
	Journeys       []*Journey
	OngoingJourney *Journey
}

func NewTravelCard(id string, balance float64) *TravelCard {
	return &TravelCard{
		ID:      id,
		balance: balance,
	}
}

func (c *TravelCard) GetLastJourney() *Journey {
	if len(c.Journeys) > 0 {
		return c.Journeys[len(c.Journeys)-1]
	}
	return nil
}

func (c *TravelCard) GetCurrentJourney() *Journey {
	return c.OngoingJourney
}

func (c *TravelCard) InitiateNewJourney(from, to string) {
	journey := NewJourney(from, to)
	lastJourney := c.GetLastJourney()
	if lastJourney != nil {
		if lastJourney.ToStn == from && !lastJourney.IsReturn {
			journey.IsReturn = true
		}
	}
	c.OngoingJourney = journey
}

func (c *TravelCard) CompleteJourney() {
	c.AddJourney(c.OngoingJourney)
	c.OngoingJourney = nil
}

func (c *TravelCard) AddJourney(journey *Journey) {
	c.Journeys = append(c.Journeys, journey)
}

func (c *TravelCard) GetBalance() float64 {
	return c.balance
}

func (c *TravelCard) Recharge(recharge float64) {
	c.balance += recharge
}

func (c *TravelCard) Debit(amount float64) error {
	if amount > c.balance {
		return fmt.Errorf("insufficient balance, has: %f, need: %f", c.balance, amount)
	}
	c.balance -= amount
	return nil
}

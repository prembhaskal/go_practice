package metro

type Journey struct {
	fromStn    string
	toStn      string
	discounted bool
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

type Card struct {
	ID       string
	Balance  float64
	Journeys []*Journey
}

func (c *Card) GetLastJourney() *Journey {
	if len(c.Journeys) > 0 {
		return c.Journeys[len(c.Journeys)-1]
	}
	return nil
}

func (c *Card) AddJourney(journey *Journey) {
	c.Journeys = append(c.Journeys, journey)
}

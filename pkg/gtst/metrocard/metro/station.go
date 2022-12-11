package metro

type RideDiscount interface {
	Apply(travelDetail *TravelDetail)
}

type TravelDetail struct {
	FromStation string
	Card        *TravelCard
	PassType    string
	Charge      float64
	Discount    float64
	OtherFees   float64
}

type Station struct {
	Name              string
	TravelDetails     []*TravelDetail
	ServiceFeePercent float64
	Discounts         []RideDiscount
}

func NewStation(name string, discounts []RideDiscount) *Station {
	return &Station{
		Name:              name,
		Discounts:         discounts,
		ServiceFeePercent: 2.0,
	}
}

func (s *Station) CheckIn(card *TravelCard, passType PassengerType, toStn string) {
	// update card journey
	// newJourney := &Journey{
	// 	FromStn:    s.Name,
	// 	ToStn:      toStn,
	// 	Discounted: travelDetail.Discount > 0,
	// }
	card.InitiateNewJourney(s.Name, toStn)

	travelDetail := &TravelDetail{
		Card:        card,
		FromStation: s.Name,
		PassType:    passType.Type(),
	}

	// get original cost
	travelDetail.Charge = passType.JourneyRate()

	s.applyDiscounts(travelDetail)

	s.RechargeCard(travelDetail)

	card.Debit(travelDetail.Charge)

	// card.AddJourney(newJourney)
	card.CompleteJourney()
	
	// update travel details.
	s.TravelDetails = append(s.TravelDetails, travelDetail)
}

// TODO - should we move it to card??
func (s *Station) RechargeCard(travelDetail *TravelDetail) {
	card := travelDetail.Card
	if card.GetBalance() >= travelDetail.Charge {
		return
	}
	recharge := travelDetail.Charge - card.GetBalance()
	card.Recharge(recharge)

	travelDetail.OtherFees += s.getServiceCharges(recharge)

	// fmt.Printf("recharge done for %f, fees is %f\n", recharge, travelDetail.OtherFees)
}

func (s *Station) applyDiscounts(rideDetail *TravelDetail) {
	for _, discounter := range s.Discounts {
		discounter.Apply(rideDetail)
	}
}

func (s *Station) getServiceCharges(amount float64) float64 {
	return amount * s.ServiceFeePercent / 100.0
}

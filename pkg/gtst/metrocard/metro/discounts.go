package metro

type ReturnJourneyDiscount struct {
	discountPercent float64
}

func NewReturnJourneyDiscount(discount float64) *ReturnJourneyDiscount {
	return &ReturnJourneyDiscount{
		discountPercent: discount,
	}
}

func (r *ReturnJourneyDiscount) Apply(travelDetail *TravelDetail) {
	originalCost := travelDetail.Charge
	currentJourney := travelDetail.Card.GetCurrentJourney()
	if currentJourney != nil && currentJourney.IsReturn {
		discount := originalCost * r.discountPercent / 100
		newcost := originalCost - discount
		travelDetail.Charge = newcost
		travelDetail.Discount += discount
	}
}

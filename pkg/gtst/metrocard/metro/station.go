package metro

import (
	"fmt"
	"io"
	"sort"
)

type ChargeDetails struct {
	cardID   string
	passType string
	fromStn  string
	charge   float64
	discount float64
}

type Station struct {
	Name              string
	CollectionDetails []ChargeDetails
	ServiceFeePercent float64
	ReturnDiscount    float64
}

func (s *Station) CheckIn(card *Card, passType PassengerType, toStn string) {
	charges := ChargeDetails{
		cardID:   card.ID,
		passType: passType.Type(),
		fromStn:  s.Name,
	}

	// get rate for the travel
	rate, discount := s.GetChargeAndDiscount(card, passType)
	charges.charge += rate
	charges.discount += discount
	// check if enough balance in card
	if card.Balance < rate {
		//   else recharge card for remaining and charge service charge
		recharge := rate - card.Balance
		card.Balance += recharge
		charges.charge += s.getServiceCharges(recharge)
	}

	// update card journey and balance
	card.Balance -= rate
	newJourney := &Journey{
		fromStn:    s.Name,
		toStn:      toStn,
		discounted: discount > 0,
	}
	card.AddJourney(newJourney)

	// update CollectionDetails.
	s.CollectionDetails = append(s.CollectionDetails, charges)
}

func (s *Station) GetChargeAndDiscount(card *Card, passType PassengerType) (float64, float64) {
	baseCharge := passType.JourneyRate()
	// calculate return discount allowed.
	discount := s.ReturnJourneyDiscount(card, passType)
	return baseCharge - discount, discount
}

func (s *Station) ReturnJourneyDiscount(card *Card, passType PassengerType) float64 {
	lastJourney := card.GetLastJourney()
	if lastJourney == nil || lastJourney.discounted || lastJourney.toStn != s.Name {
		return 0.0
	}

	return passType.JourneyRate() * s.ReturnDiscount / 100.0
}

func (s *Station) getServiceCharges(amount float64) float64 {
	return amount * s.ServiceFeePercent / 100.0
}

func (s *Station) PrintSummary(w io.Writer) error {
	var err error
	totalColl, totalDisc := s.getTotalCollectionAndDiscount()
	_, err = fmt.Fprintf(w, "TOTAL_COLLECTION %s %d %d\n", s.Name, int(totalColl), int(totalDisc))
	if err != nil {
		return fmt.Errorf("error writing summary: %w", err)
	}
	_, err = fmt.Fprintf(w, "PASSENGER_TYPE_SUMMARY\n")
	if err != nil {
		return fmt.Errorf("error writing summary: %w", err)
	}
	typeSummary := s.GetPassengerTypeSummary()
	for _, summ := range typeSummary {
		_, err = fmt.Fprintf(w, "%s %d\n", summ.passType, summ.count)
		if err != nil {
			return fmt.Errorf("error writing summary: %w", err)
		}
	}
	return nil
}

func (s *Station) getTotalCollectionAndDiscount() (float64, float64) {
	var totalColl, totalDisc float64
	for _, collection := range s.CollectionDetails {
		totalColl = totalColl + collection.charge
		totalDisc = totalDisc + collection.discount
	}
	return totalColl, totalDisc
}

type typeSumm struct {
	passType  string
	totalColl float64
	count     int
}

func (s *Station) GetPassengerTypeSummary() []*typeSumm {
	summaryMap := make(map[string]*typeSumm)
	for _, collection := range s.CollectionDetails {
		typeSummary := summaryMap[collection.passType]
		if typeSummary == nil {
			typeSummary = &typeSumm{
				passType: collection.passType,
			}
			summaryMap[collection.passType] = typeSummary
		}
		typeSummary.passType = collection.passType
		typeSummary.count++
		typeSummary.totalColl += collection.charge
	}

	summary := make([]*typeSumm, 0)
	for _, summ := range summaryMap {
		summary = append(summary, summ)
	}

	// sort by largest collection for each type
	sort.Slice(summary, func(i, j int) bool {
		return summary[i].totalColl > summary[j].totalColl
	})

	return summary
}

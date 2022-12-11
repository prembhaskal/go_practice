package metro

import (
	"fmt"
	"io"
	"sort"
)

func PrintSummary(s *Station, w io.Writer) error {
	var err error
	totalColl, totalDisc := getTotalCollectionAndDiscount(s)
	_, err = fmt.Fprintf(w, "TOTAL_COLLECTION %s %d %d\n", s.Name, int(totalColl), int(totalDisc))
	if err != nil {
		return fmt.Errorf("error writing summary: %w", err)
	}
	_, err = fmt.Fprintf(w, "PASSENGER_TYPE_SUMMARY\n")
	if err != nil {
		return fmt.Errorf("error writing summary: %w", err)
	}
	typeSummary := GetPassengerTypeSummary(s)
	for _, summ := range typeSummary {
		_, err = fmt.Fprintf(w, "%s %d\n", summ.passType, summ.count)
		if err != nil {
			return fmt.Errorf("error writing summary: %w", err)
		}
	}
	return nil
}

func getTotalCollectionAndDiscount(s *Station) (float64, float64) {
	var totalColl, totalDisc float64
	for _, travelDetail := range s.TravelDetails {
		totalColl = totalColl + travelDetail.Charge + travelDetail.OtherFees
		totalDisc = totalDisc + travelDetail.Discount
	}
	return totalColl, totalDisc
}

type typeSumm struct {
	passType  string
	totalColl float64
	count     int
}

func GetPassengerTypeSummary(s *Station) []*typeSumm {
	summaryMap := make(map[string]*typeSumm)
	for _, travelDetail := range s.TravelDetails {
		typeSummary := summaryMap[travelDetail.PassType]
		if typeSummary == nil {
			typeSummary = &typeSumm{
				passType: travelDetail.PassType,
			}
			summaryMap[travelDetail.PassType] = typeSummary
		}
		typeSummary.passType = travelDetail.PassType
		typeSummary.count++
		typeSummary.totalColl = typeSummary.totalColl + travelDetail.Charge + travelDetail.OtherFees
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

package metro

type AdultPassenger struct{}

func (a AdultPassenger) Charges() float64 {
	return 200
}

type SeniorPassenger struct{}

func (s SeniorPassenger) Charges() float64 {
	return 100
}

type KidPassenger struct{}

func (k KidPassenger) Charges() float64 {
	return 50
}

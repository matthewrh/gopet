package status

type Status int

const (
	Dead     Status = iota
	Starving Status = iota
	Hungry   Status = iota
	Full     Status = iota
	Sad      Status = iota
	Happy    Status = iota
	Tired    Status = iota
	Rested   Status = iota
	Dirty    Status = iota
	Clean    Status = iota
)

func (s Status) String() string {
	return [...]string{"Dead", "Starving", "Hungry", "Full", "Sad", "Happy", "Tired", "Rested", "Dirty", "Clean"}[s]
}

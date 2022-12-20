package petState

type PetState int

const (
	Sad      PetState = iota
	Eating   PetState = iota
	Hungry   PetState = iota
	Normal   PetState = iota
	Play     PetState = iota
	Sleeping PetState = iota
	Tired    PetState = iota
	Dead     PetState = iota
)

func (ps PetState) String() string {
	return [...]string{"Sad", "Eating", "Hungry", "Normal", "Play", "Sleeping", "Tired", "Dead"}[ps]
}

package state

type GameState int

const (
	Active GameState = iota
	Paused GameState = iota
)

func (gs GameState) String() string {
	return [...]string{"Active", "Paused"}[gs]
}

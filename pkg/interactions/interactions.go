package interactions

import "container/ring"

const (
	Feed  = "feed"
	Play  = "play"
	Wash  = "wash"
	Sleep = "sleep"
	Quit  = "quit"
)

var (
	interactions = [...]string{Feed, Play, Wash, Sleep, Quit}
)

func NewInteractions() *ring.Ring {
	r := ring.New(len(interactions))
	for _, v := range interactions {
		r.Value = v
		r = r.Next()
	}
	return r
}

package images

import (
	"container/ring"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gopet/pkg/petState"
)

var Tired1, _, _ = ebitenutil.NewImageFromFile("images/sad_1.png")
var Tired2, _, _ = ebitenutil.NewImageFromFile("images/sad_2.png")
var Tired3, _, _ = ebitenutil.NewImageFromFile("images/sad_3.png")
var Eating1, _, _ = ebitenutil.NewImageFromFile("images/eating_1.png")
var Eating2, _, _ = ebitenutil.NewImageFromFile("images/eating_2.png")
var Eating3, _, _ = ebitenutil.NewImageFromFile("images/eating_3.png")
var Hungry1, _, _ = ebitenutil.NewImageFromFile("images/hungry_1.png")
var Normal1, _, _ = ebitenutil.NewImageFromFile("images/normal_1.png")
var Normal2, _, _ = ebitenutil.NewImageFromFile("images/normal_2.png")
var Play1, _, _ = ebitenutil.NewImageFromFile("images/play_1.png")
var Play2, _, _ = ebitenutil.NewImageFromFile("images/play_2.png")
var Dead1, _, _ = ebitenutil.NewImageFromFile("images/dead_1.png")
var Sleeping1, _, _ = ebitenutil.NewImageFromFile("images/sleeping_1.png")
var Sleeping2, _, _ = ebitenutil.NewImageFromFile("images/sleeping_2.png")
var Sad1, _, _ = ebitenutil.NewImageFromFile("images/tired_1.png")
var Sad2, _, _ = ebitenutil.NewImageFromFile("images/tired_2.png")

var (
	sad      = [...]*ebiten.Image{Sad1, Sad2}
	eating   = [...]*ebiten.Image{Eating1, Eating2, Eating3}
	hungry   = [...]*ebiten.Image{Hungry1}
	normal   = [...]*ebiten.Image{Normal1, Normal2}
	play     = [...]*ebiten.Image{Play1, Play2}
	dead     = [...]*ebiten.Image{Dead1}
	sleeping = [...]*ebiten.Image{Sleeping1, Sleeping2}
	tired    = [...]*ebiten.Image{Tired1, Tired2, Tired3}
)

func GetImages(stat petState.PetState) *ring.Ring {
	var r *ring.Ring
	switch stat {
	case petState.Sad:
		r = ring.New(len(sad))
		for _, v := range sad {
			r.Value = v
			r = r.Next()
		}
	case petState.Eating:
		r = ring.New(len(eating))
		for _, v := range eating {
			r.Value = v
			r = r.Next()
		}
	case petState.Hungry:
		r = ring.New(len(hungry))
		for _, v := range hungry {
			r.Value = v
			r = r.Next()
		}
	case petState.Normal:
		r = ring.New(len(normal))
		for _, v := range normal {
			r.Value = v
			r = r.Next()
		}
	case petState.Play:
		r = ring.New(len(play))
		for _, v := range play {
			r.Value = v
			r = r.Next()
		}
	case petState.Dead:
		r = ring.New(len(dead))
		for _, v := range dead {
			r.Value = v
			r = r.Next()
		}
	case petState.Sleeping:
		r = ring.New(len(sleeping))
		for _, v := range sleeping {
			r.Value = v
			r = r.Next()
		}
	case petState.Tired:
		r = ring.New(len(tired))
		for _, v := range tired {
			r.Value = v
			r = r.Next()
		}
	}
	return r
}

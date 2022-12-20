package game

import (
	"container/ring"
	"errors"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/disintegration/imaging"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gopet/pkg/images"
	"gopet/pkg/interactions"
	"gopet/pkg/pet"
	"gopet/pkg/petState"
	"gopet/pkg/state"
	"gopet/pkg/status"
	"image"
	"image/color"
	"os"
	"strconv"
	"strings"
	"time"
)

var Terminated = errors.New("terminated")

type Game struct {
	state                       state.GameState
	selectedInteraction         *ring.Ring
	currentInteraction          string
	currentInteractionTimestamp int
	petStatePrev                petState.PetState
	petState                    petState.PetState
	currentStatusImage          *ring.Ring
	currentStatusImageTimestamp int
	pet                         *pet.Pet
}

func (g *Game) Create(p *pet.Pet) {
	g.state = state.Active
	g.selectedInteraction = interactions.NewInteractions()
	g.currentInteraction = ""
	g.currentInteractionTimestamp = 0
	g.petStatePrev = petState.Normal
	g.petState = petState.Normal
	g.currentStatusImage = images.GetImages(g.petState)
	g.currentStatusImageTimestamp = int(time.Now().Unix())
	g.pet = p
}

func (g *Game) UpdatePetState(p petState.PetState) {
	if g.petState != p {
		g.petStatePrev = g.petState
	}
	g.petState = p
	g.currentStatusImage = images.GetImages(g.petState)
	g.currentStatusImageTimestamp = int(time.Now().Unix())
}

func (g *Game) LogState() {
	for g.pet.IsAlive() {
		if g.state == state.Active {
			fmt.Print("\033[H\033[2J")
			var petUpdateStatus string = g.pet.Update()
			g.pet.Status()
			fmt.Println("")
			if strings.Contains(petUpdateStatus, status.Dead.String()) {
				fmt.Println(g.pet.Name, "has died :(")
				g.UpdatePetState(petState.Dead)
				break
			} else if strings.Contains(petUpdateStatus, status.Hungry.String()) || strings.Contains(petUpdateStatus, status.Starving.String()) {
				if g.petState != petState.Sleeping && g.petState != petState.Hungry && g.currentInteractionTimestamp == 0 {
					g.UpdatePetState(petState.Hungry)
				}
			} else if strings.Contains(petUpdateStatus, status.Tired.String()) {
				if g.petState != petState.Sleeping && g.petState != petState.Tired && g.currentInteractionTimestamp == 0 {
					g.UpdatePetState(petState.Tired)
				}
			} else if strings.Contains(petUpdateStatus, status.Sad.String()) {
				if g.petState != petState.Sleeping && g.petState != petState.Sad && g.currentInteractionTimestamp == 0 {
					g.UpdatePetState(petState.Sad)
				}
			} else {
				if g.petState != petState.Sleeping && g.currentInteractionTimestamp == 0 {
					g.UpdatePetState(petState.Normal)
				}
			}

			fmt.Println(g.pet.Name, "is", petUpdateStatus)
			time.Sleep(time.Second * 3)
		} else {
			fmt.Print("\033[H\033[2J")
			g.pet.Status()
			fmt.Println("")
			fmt.Println("Game is paused. Press space to continue.")

			time.Sleep(time.Second / 4)
		}
	}
}

func (g *Game) Update() error {
	if g.pet.IsAlive() {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			if g.state == state.Active {
				g.state = state.Paused
			} else {
				g.state = state.Active
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			g.selectedInteraction = g.selectedInteraction.Prev()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			g.selectedInteraction = g.selectedInteraction.Next()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			if g.selectedInteraction.Value == interactions.Feed {
				if g.petState == petState.Sleeping {
					g.UpdatePetState(petState.Normal)
				}
				g.pet.Feed()
				g.UpdatePetState(petState.Eating)
				g.currentInteraction = interactions.Feed
				g.currentInteractionTimestamp = int(time.Now().Unix())
			} else if g.selectedInteraction.Value == interactions.Play {
				if g.petState == petState.Sleeping {
					g.UpdatePetState(petState.Normal)
				}
				g.pet.Play()
				g.UpdatePetState(petState.Play)
				g.currentInteraction = interactions.Play
				g.currentInteractionTimestamp = int(time.Now().Unix())
			} else if g.selectedInteraction.Value == interactions.Wash {
				if g.petState == petState.Sleeping {
					g.UpdatePetState(petState.Normal)
				}
				var cleanState string = g.pet.Clean()
				g.UpdatePetState(petState.Normal)
				g.currentInteraction = interactions.Wash + cleanState
				g.currentInteractionTimestamp = int(time.Now().Unix())
			} else if g.selectedInteraction.Value == interactions.Sleep {
				g.pet.Sleep()
				g.UpdatePetState(petState.Sleeping)
				g.currentInteraction = interactions.Sleep
				g.currentInteractionTimestamp = int(time.Now().Unix())
			} else if g.selectedInteraction.Value == interactions.Quit {
				if err := g.pet.Save("pet.json"); err != nil {
					fmt.Println("Error saving pet data:", err)
				} else {
					fmt.Println("Goodbye!")
				}
				return Terminated
			}
		}
	} else {
		time.Sleep(time.Second * 5)
		if err := os.Remove("pet.json"); err != nil {
			fmt.Println("Error deleting pet data file:", err)
		}
		return Terminated
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Hunger: "+strconv.Itoa(g.pet.Hunger)+" ("+g.pet.HungerStatus()+")", 0, 20-15)
	ebitenutil.DebugPrintAt(screen, "Happiness: "+strconv.Itoa(g.pet.Happiness)+" ("+g.pet.HappinessStatus()+")", 0, 35-15)
	ebitenutil.DebugPrintAt(screen, "Fatigue: "+strconv.Itoa(g.pet.Fatigue)+" ("+g.pet.FatigueStatus()+")", 0, 50-15)
	ebitenutil.DebugPrintAt(screen, "Life: "+strconv.Itoa(g.pet.Life), 0, 65-15)
	if g.pet.IsDirty {
		ebitenutil.DebugPrintAt(screen, "Dirty", 0, 80-15)
	} else {
		ebitenutil.DebugPrintAt(screen, "Clean", 0, 80-15)
	}
	if g.pet.IsSleeping {
		ebitenutil.DebugPrintAt(screen, "Sleeping", 0, 95-15)
	} else {
		ebitenutil.DebugPrintAt(screen, "Awake", 0, 95-15)
	}

	img := g.currentStatusImage.Value.(*ebiten.Image)
	loadImg := image.Image(img)
	inverted := imaging.Invert(loadImg)
	newImg := ebiten.NewImageFromImage(inverted)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(22, 12)
	opts.GeoM.Scale(6, 6)
	screen.DrawImage(newImg, opts)
	if int(time.Now().Unix())-g.currentStatusImageTimestamp >= 1 {
		g.currentStatusImage = g.currentStatusImage.Next()
		g.currentStatusImageTimestamp = int(time.Now().Unix())
	}

	g.DrawInteractionsBar(screen)

	if int(time.Now().Unix())-g.currentInteractionTimestamp < 3 {
		if g.currentInteraction == interactions.Feed {
			ebitenutil.DebugPrintAt(screen, g.pet.Name+" eats a yummy snack!", 60, 210)
		} else if g.currentInteraction == interactions.Play {
			ebitenutil.DebugPrintAt(screen, g.pet.Name+" plays and has fun!", 61, 210)
		} else if g.currentInteraction == interactions.Wash {
			ebitenutil.DebugPrintAt(screen, g.pet.Name+" was given a bath!", 62, 210)
		} else if g.currentInteraction == interactions.Wash+"+" {
			ebitenutil.DebugPrintAt(screen, g.pet.Name+" was given a good bath!", 46, 210)
		} else if g.currentInteraction == interactions.Wash+"-" {
			ebitenutil.DebugPrintAt(screen, g.pet.Name+" was given a bad bath...", 43, 210)
		} else if g.currentInteraction == interactions.Sleep {
			ebitenutil.DebugPrintAt(screen, g.pet.Name+" was put to bed...", 65, 210)
		}
	} else if int(time.Now().Unix())-g.currentInteractionTimestamp < 4 {
		g.currentInteraction = ""
		g.currentInteractionTimestamp = 0
		if g.petState != petState.Sleeping && g.petState != petState.Dead {
			g.UpdatePetState(g.petStatePrev)
		}
	} else {
		g.currentInteraction = ""
		g.currentInteractionTimestamp = 0
	}
}

func (g *Game) DrawInteractionsBar(screen *ebiten.Image) {
	if g.selectedInteraction.Value == interactions.Feed {
		ebitenutil.DrawRect(screen, 20, 200-20, 40, 20, lipgloss.Color("#04B575"))
		ebitenutil.DrawRect(screen, 21, 201-20, 38, 18, color.Black)
		ebitenutil.DebugPrintAt(screen, "Feed", 27, 202-20)
	} else {
		ebitenutil.DrawRect(screen, 20, 200-20, 40, 20, color.White)
		ebitenutil.DrawRect(screen, 21, 201-20, 38, 18, color.Black)
		ebitenutil.DebugPrintAt(screen, "Feed", 27, 202-20)
	}

	if g.selectedInteraction.Value == interactions.Play {
		ebitenutil.DrawRect(screen, 20+45, 200-20, 40, 20, lipgloss.Color("#04B575"))
		ebitenutil.DrawRect(screen, 21+45, 201-20, 38, 18, color.Black)
		ebitenutil.DebugPrintAt(screen, "Play", 27+45, 202-20)
	} else {
		ebitenutil.DrawRect(screen, 20+45, 200-20, 40, 20, color.White)
		ebitenutil.DrawRect(screen, 21+45, 201-20, 38, 18, color.Black)
		ebitenutil.DebugPrintAt(screen, "Play", 27+45, 202-20)
	}

	if g.selectedInteraction.Value == interactions.Wash {
		ebitenutil.DrawRect(screen, 20+90, 200-20, 40, 20, lipgloss.Color("#04B575"))
		ebitenutil.DrawRect(screen, 21+90, 201-20, 38, 18, color.Black)
		ebitenutil.DebugPrintAt(screen, "Wash", 27+90, 202-20)
	} else {
		ebitenutil.DrawRect(screen, 20+90, 200-20, 40, 20, color.White)
		ebitenutil.DrawRect(screen, 21+90, 201-20, 38, 18, color.Black)
		ebitenutil.DebugPrintAt(screen, "Wash", 27+90, 202-20)
	}

	if g.selectedInteraction.Value == interactions.Sleep {
		ebitenutil.DrawRect(screen, 20+135, 200-20, 40, 20, lipgloss.Color("#04B575"))
		ebitenutil.DrawRect(screen, 21+135, 201-20, 38, 18, color.Black)
		ebitenutil.DebugPrintAt(screen, "Sleep", 24+135, 202-20)
	} else {
		ebitenutil.DrawRect(screen, 20+135, 200-20, 40, 20, color.White)
		ebitenutil.DrawRect(screen, 21+135, 201-20, 38, 18, color.Black)
		ebitenutil.DebugPrintAt(screen, "Sleep", 24+135, 202-20)
	}

	if g.selectedInteraction.Value == interactions.Quit {
		ebitenutil.DrawRect(screen, 20+180, 200-20, 40, 20, lipgloss.Color("#04B575"))
		ebitenutil.DrawRect(screen, 21+180, 201-20, 38, 18, color.Black)
		ebitenutil.DebugPrintAt(screen, "Quit", 27+180, 202-20)
	} else {
		ebitenutil.DrawRect(screen, 20+180, 200-20, 40, 20, color.White)
		ebitenutil.DrawRect(screen, 21+180, 201-20, 38, 18, color.Black)
		ebitenutil.DebugPrintAt(screen, "Quit", 27+180, 202-20)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 260, 240
}

package main

import (
	"fmt"
	"os"
	"gopet/pkg/utils"
	"gopet/pkg/game"
	"gopet/pkg/pet"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.Game{}
	var myPet *pet.Pet

	_, err := os.Stat("pet.json")
	if err == nil {
		myPet, err = utils.LoadPet("pet.json")
		if err != nil {
			fmt.Println("Error loading pet data:", err)
			return
		}
		fmt.Println("Welcome back!", myPet.Name, "missed you!")
	} else if os.IsNotExist(err) {
		fmt.Println("Welcome to GoPet! What would you like to name your new pet?")
		var name string
		fmt.Scanln(&name)
		myPet = pet.NewPet(name)
		fmt.Println("Welcome to the world, little", myPet.Name)
	} else {
		fmt.Println("Error checking for pet data file:", err)
		return
	}

	g.Create(myPet)

	go g.LogState()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gopet: " + myPet.Name)
	if err := ebiten.RunGame(&g); err != nil {
		if err == game.Terminated {
			return
		}
		log.Fatal(err)
	}
}
package game

import (
	"fmt"
	"gopet/pkg/pet"
	"gopet/pkg/status"
	"gopet/pkg/utils"
	"os"
	"strings"
	"time"
)

func Game() {
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

	for {
		fmt.Print("\033[H\033[2J")
		myPet.Status()

		var sleepOption string
		if myPet.IsSleeping {
			sleepOption = "wake"
			fmt.Println("What would you like to do? (" + sleepOption + ", nothing, quit)")
		} else {
			sleepOption = "sleep"
			fmt.Println("What would you like to do? (feed, play, wash, " + sleepOption + ", nothing, quit)")
		}

		var command string
		fmt.Scanln(&command)

		fmt.Println("")

		switch command {
		case "feed":
			if sleepOption == "wake" {
				fmt.Println("Invalid command")
				time.Sleep(time.Second)
				continue
			} else {
				myPet.Feed()
			}
		case "play":
			if sleepOption == "wake" {
				fmt.Println("Invalid command")
				time.Sleep(time.Second)
				continue
			} else {
				myPet.Play()
			}
		case "wash":
			if sleepOption == "wake" {
				fmt.Println("Invalid command")
				time.Sleep(time.Second)
				continue
			} else {
				myPet.Clean()
			}
		case "sleep":
			if sleepOption == "wake" {
				fmt.Println("Invalid command")
				time.Sleep(time.Second)
				continue
			} else {
				myPet.Sleep()
			}
		case "wake":
			if sleepOption == "sleep" {
				fmt.Println("Invalid command")
				time.Sleep(time.Second)
				continue
			} else {
				myPet.WakeUp()
			}
		case "nothing":
			fmt.Println("You chose to do nothing.")
		case "quit":
			if err := myPet.Save("pet.json"); err != nil {
				fmt.Println("Error saving pet data:", err)
			} else {
				fmt.Println("Goodbye!")
			}
			return
		default:
			fmt.Println("Invalid command")
			time.Sleep(time.Second)
			continue
		}

		fmt.Println("")

		var petUpdateStatus string = myPet.Update()
		if strings.Contains(petUpdateStatus, status.Dead.String()) {
			fmt.Println(myPet.Name, "has died :(")
			break
		}

		fmt.Println(myPet.Name, "is", petUpdateStatus)
		time.Sleep(time.Second * 3)
	}

	if err := os.Remove("pet.json"); err != nil {
		fmt.Println("Error deleting pet data file:", err)
	}
}

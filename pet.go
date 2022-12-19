package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Pet struct {
	Name       string
	Age        int
	Hunger     int
	Happiness  int
	IsSleeping bool
}

func NewPet(name string) *Pet {
	return &Pet{
		Name:       name,
		Age:        0,
		Hunger:     50,
		Happiness:  50,
		IsSleeping: false,
	}
}

func (p *Pet) Sleep() {
	p.IsSleeping = true
	fmt.Println(p.Name, "is sleeping...")
}

func (p *Pet) WakeUp() {
	p.IsSleeping = false
	fmt.Println(p.Name, "wakes up!")
}

func (p *Pet) Feed() {
	p.Hunger -= 10
	fmt.Println(p.Name, "eats a yummy snack!")
}

func (p *Pet) Play() {
	p.Happiness += 10
	fmt.Println(p.Name, "plays and has fun!")
}

func (p *Pet) IncreaseAge() {
	p.Age++
	fmt.Println(p.Name, "is now", p.Age, "years old!")
}

func (p *Pet) Update() {
	if p.IsSleeping {
		p.WakeUp()
	}
	if p.Hunger < 100 {
		p.Hunger += 5
	}
	if p.Happiness > 0 {
		p.Happiness--
	}
	p.IncreaseAge()

	if p.Hunger >= 100 || p.Happiness <= 0 {
		fmt.Println(p.Name, "has died :(")
	}
}

func (p *Pet) Save(filename string) error {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func LoadPet(filename string) (*Pet, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var pet Pet
	if err := json.Unmarshal(data, &pet); err != nil {
		return nil, err
	}
	return &pet, nil
}

func main() {
	var pet *Pet

	_, err := os.Stat("pet.json")
	if err == nil {
		pet, err = LoadPet("pet.json")
		if err != nil {
			fmt.Println("Error loading pet data:", err)
			return
		}
		fmt.Println("Welcome back,", pet.Name)
	} else if os.IsNotExist(err) {
		fmt.Println("What would you like to name your new pet?")
		var name string
		fmt.Scanln(&name)
		pet = NewPet(name)
		fmt.Println("Welcome to the world, little", pet.Name)
	} else {
		fmt.Println("Error checking for pet data file:", err)
		return
	}

	for {
		fmt.Println("---")
		fmt.Println(pet.Name, "is", pet.Age, "years old,", pet.Hunger, "hunger,", pet.Happiness, "happiness")

		fmt.Println("What would you like to do? (feed, play, sleep, quit)")

		var command string
		fmt.Scanln(&command)

		switch command {
		case "feed":
			pet.Feed()
		case "play":
			pet.Play()
		case "sleep":
			pet.Sleep()
		case "quit":
			if err := pet.Save("pet.json"); err != nil {
				fmt.Println("Error saving pet data:", err)
			} else {
				fmt.Println("Goodbye!")
			}
			return
		default:
			fmt.Println("Invalid command")
		}

		pet.Update()

		if pet.Hunger >= 100 || pet.Happiness <= 0 {
			fmt.Println("Game over :(")
			break
		}

		time.Sleep(time.Second)
	}

	if err := os.Remove("pet.json"); err != nil {
		fmt.Println("Error deleting pet data file:", err)
	}
}

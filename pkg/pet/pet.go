package pet

import (
	"encoding/json"
	"fmt"
	"gopet/pkg/status"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type Pet struct {
	Name       string
	Life       int
	Hunger     int
	Happiness  int
	Fatigue    int
	IsDirty    bool
	IsSleeping bool
	Birthday   int
}

const (
	// Threshold_Life_Min          : Min level of life
	//                             : Life = 0 is dead
	Threshold_Life_Min = 0
	// Threshold_Life_Max          : Max level of life
	//                             : Life = 100 is max
	Threshold_Life_Max = 100

	// Threshold_Hunger_Min        : Min level of hunger
	//                             : Hunger = 0 is min
	Threshold_Hunger_Min = 0
	// Threshold_Hunger_Hungry     : Min level of hunger before hungry
	//                             : Hunger > 50 is hungry
	Threshold_Hunger_Hungry = 50
	// Threshold_Hunger_Starving   : Min level of hunger before starving
	//                             : Hunger > 75 is starving
	Threshold_Hunger_Starving = 75
	// Threshold_Hunger_Max        : Max level of hunger
	//                             : Hunger = 100 is max
	Threshold_Hunger_Max = 100

	// Threshold_Happiness_Min     : Min level of hapiness
	//                             : Happiness = 0 is dead
	Threshold_Happiness_Min = 0
	// Threshold_Happiness_Sad     : Min level of hapiness to be happy
	//                             : Happiness < 25 is sad
	Threshold_Happiness_Sad = 25
	// Threshold_Happiness_Max     : Max level of happiness
	//                             : Happiness = 100 is max
	Threshold_Happiness_Max = 50

	// Threshold_Fatigue_Min       : Min level of fatigue
	//                             : Fatigue = 0 is min
	Threshold_Fatigue_Min = 0
	// Threshold_Fatigue_Tired     : Max level of fatigue before sleeping
	//                             : Fatigue > 90 is tired
	Threshold_Fatigue_Tired = 90
	// Threshold_Fatigue_Max       : Max level of fatigue
	//                             : Fatigue = 100 is max
	Threshold_Fatigue_Max = 100
)

func NewPet(name string) *Pet {
	return &Pet{
		Name:       name,
		Life:       100,
		Hunger:     25,
		Happiness:  50,
		Fatigue:    50,
		IsDirty:    false,
		IsSleeping: false,
		Birthday:   int(time.Now().Unix()),
	}
}

func (p *Pet) Save(filename string) error {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func (p *Pet) Status() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Stats:")
	fmt.Println(" [+]", p.Hunger, "Hunger", "("+p.HungerStatus()+")")
	fmt.Println(" [+]", p.Happiness, "Happiness", "("+p.HappinessStatus()+")")
	fmt.Println(" [+]", p.Fatigue, "Fatigue", "("+p.FatigueStatus()+")")
	fmt.Println(" [+]", p.Life, "Life")
	if p.IsDirty {
		fmt.Println(" [+]", "Dirty")
	} else {
		fmt.Println(" [+]", "Clean")
	}
	if p.IsSleeping {
		fmt.Println(" [+]", "Asleep")
	} else {
		fmt.Println(" [+]", "Awake")
	}
	fmt.Println("")
}

func (p *Pet) Update() string {
	var statuses []string

	if p.IsAlive() {
		p.IncreaseHunger(1)
		p.DecreaseHappiness(1)
		p.IncreaseFatigue(1)

		if p.IsSleeping {
			p.DecreaseFatigue(10)
			p.IncreaseLife(1)
		}

		if p.HungerStatus() == status.Starving.String() {
			p.DecreaseLife(2)
			p.IncreaseHunger(3)
			statuses = append(statuses, status.Starving.String())
		} else if p.HungerStatus() == status.Hungry.String() {
			p.DecreaseLife(1)
			p.IncreaseHunger(2)
			statuses = append(statuses, status.Hungry.String())
		} else if p.HungerStatus() == status.Full.String() {
			statuses = append(statuses, status.Full.String())
		}

		if p.HappinessStatus() == status.Sad.String() {
			p.DecreaseLife(1)
			statuses = append(statuses, status.Sad.String())
		} else if p.HappinessStatus() == status.Happy.String() {
			statuses = append(statuses, status.Happy.String())
		}

		if p.FatigueStatus() == status.Tired.String() {
			p.DecreaseLife(1)
			statuses = append(statuses, status.Tired.String())
		} else if p.FatigueStatus() == status.Rested.String() {
			statuses = append(statuses, status.Rested.String())
		}

		if !p.IsDirty {
			p.RandomizeDirty()
		} else {
			p.DecreaseHappiness(5)
		}

		if p.IsDirty {
			statuses = append(statuses, status.Dirty.String())
		} else {
			statuses = append(statuses, status.Clean.String())
		}

		if !p.IsAlive() {
			statuses = append(statuses, status.Dead.String())
		}
	} else {
		statuses = append(statuses, status.Dead.String())
	}

	return strings.Join(statuses, ", ")
}

func (p *Pet) IsAlive() bool {
	if p.Life <= 0 {
		return false
	}
	return true
}

func (p *Pet) Feed() {
	if p.IsSleeping {
		p.WakeUp()
	}
	p.DecreaseHunger(5)
	fmt.Println(p.Name, "eats a yummy snack!")
}

func (p *Pet) Play() {
	if p.IsSleeping {
		p.WakeUp()
	}
	p.IncreaseHappiness(5)
	fmt.Println(p.Name, "plays and has fun!")
}

func (p *Pet) Sleep() {
	if p.IsSleeping {
		fmt.Println(p.Name, "is already sleeping!")
		return
	}
	p.IsSleeping = true
	fmt.Println(p.Name, "is sleeping...")
}

func (p *Pet) WakeUp() {
	p.IsSleeping = false
	fmt.Println(p.Name, "wakes up!")
}

func (p *Pet) Clean() string {
	if p.IsSleeping {
		p.WakeUp()
	}
	p.IsDirty = false
	if rand.Intn(100) < 10 {
		fmt.Println(p.Name, "was given a bath, but they didn't like it...")
		p.DecreaseHappiness(15)
		return "-"
	} else if rand.Intn(100) > 90 {
		fmt.Println(p.Name, "was given a bath, and they loved it!")
		p.IncreaseHappiness(5)
		return "+"
	} else {
		fmt.Println(p.Name, "was given a bath!")
		return ""
	}
}

func (p *Pet) HungerStatus() string {
	if p.Hunger >= Threshold_Hunger_Starving {
		return status.Starving.String()
	} else if p.Hunger >= Threshold_Hunger_Hungry {
		return status.Hungry.String()
	}
	return status.Full.String()
}

func (p *Pet) HappinessStatus() string {
	if p.Happiness <= Threshold_Happiness_Sad {
		return status.Sad.String()
	}
	return status.Happy.String()
}

func (p *Pet) FatigueStatus() string {
	if p.Fatigue >= Threshold_Fatigue_Tired {
		return status.Tired.String()
	}
	return status.Rested.String()
}

func (p *Pet) RandomizeDirty() {
	if rand.Intn(100) < 5 {
		p.IsDirty = true
	} else {
		p.IsDirty = false
	}
}

func (p *Pet) IncreaseLife(amount int) {
	if (p.Life + amount) > Threshold_Life_Max {
		p.Life = Threshold_Life_Max
	} else {
		p.Life += amount
	}
}

func (p *Pet) DecreaseLife(amount int) {
	if (p.Life - amount) < Threshold_Life_Min {
		p.Life = Threshold_Life_Min
	} else {
		p.Life -= amount
	}
}

func (p *Pet) IncreaseHunger(amount int) {
	if (p.Hunger + amount) > Threshold_Hunger_Max {
		p.Hunger = Threshold_Hunger_Max
	} else {
		p.Hunger += amount
	}
}

func (p *Pet) DecreaseHunger(amount int) {
	if (p.Hunger - amount) < Threshold_Hunger_Min {
		p.Hunger = Threshold_Hunger_Min
	} else {
		p.Hunger -= amount
	}
}

func (p *Pet) IncreaseHappiness(amount int) {
	if (p.Happiness + amount) > Threshold_Happiness_Max {
		p.Happiness = Threshold_Happiness_Max
	} else {
		p.Happiness += amount
	}
}

func (p *Pet) DecreaseHappiness(amount int) {
	if (p.Happiness - amount) < Threshold_Happiness_Min {
		p.Happiness = Threshold_Happiness_Min
	} else {
		p.Happiness -= amount
	}
}

func (p *Pet) IncreaseFatigue(amount int) {
	if (p.Fatigue + amount) > Threshold_Fatigue_Max {
		p.Fatigue = Threshold_Fatigue_Max
	} else {
		p.Fatigue += amount
	}
}

func (p *Pet) DecreaseFatigue(amount int) {
	if (p.Fatigue - amount) < Threshold_Fatigue_Min {
		p.Fatigue = Threshold_Fatigue_Min
	} else {
		p.Fatigue -= amount
	}
}

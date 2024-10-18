package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Creature struct {
	food string
	locomotion string
	noise string
}

func (a Creature) Eat() {
	fmt.Println(a.food)
}

func (a Creature) Move() {
	fmt.Println(a.locomotion)
}

func (a Creature) Speak() {
	fmt.Println(a.noise)
}

var animals = make(map[string]Animal)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		input = strings.TrimSpace(input)
		handleInput(input)
	}
}

func handleInput(input string) {
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		fmt.Println("Error: Incorrect input")
		return
	}

	switch parts[0] {
		case "newanimal":
			AddAnimal(parts[1], parts[2])
		case "query":
			GetAnimalInfo(parts[1], parts[2])
		default:
			fmt.Println("Error: Incorrect input")
	}
}

func GetAnimalInfo(name string, info string) {
	animal, exists := animals[name]

	if !exists {
		fmt.Println("Error: Animal not found")
		return
	}

	switch info {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Error: Incorrect info request")
	}
}

func AddAnimal(name string, animalType string) {
	var animal Animal

	switch animalType {
		case "cow":
			animal = &Creature{"grass", "walk", "moo"}
		case "bird":
			animal = &Creature{"worms", "fly", "peep"}
		case "snake":
			animal = &Creature{"mice", "slither", "hsss"}
		default:
			fmt.Println("Error: Incorrect animal")
			return
	}

	animals[name] = animal
	fmt.Println("Created it!")
}

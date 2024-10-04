package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

var cow = Animal{"grass", "walk", "moo"}
var bird = Animal{"worms", "fli", "peep"}
var snake = Animal{"mice", "slither", "moo"}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		handleInput(input)
	}
}

func handleInput(input string) {
	animal, method, err := GetAnimalAndMethod(input)
	if err != nil {
		fmt.Println(err)
		return;
	}

	switch method {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Error: Incorrect method")
	}
}

func GetAnimalAndMethod(input string) (Animal, string, error) {
	parts := strings.Split(input, " ")

	animalsMap := map[string]Animal{
		"cow": cow,
		"bird": bird,
		"snake": snake,
	}

	if len(parts) != 2 {
		return Animal{}, "", errors.New("Error: Incorrect input")
	}

	animal, exists := animalsMap[parts[0]]

	if exists == false {
		return Animal{}, "", errors.New("Error: Animal not found")
	}

	return animal, strings.TrimSpace(parts[1]), nil
}

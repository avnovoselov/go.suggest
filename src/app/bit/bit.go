package bit

import (
	"fmt"
	"strings"
)

const (
	HAND     byte = 1 << 0
	TAIL     byte = 1 << 1
	EYE      byte = 1 << 2
	TENTACLE byte = 1 << 3
)

// Check do animal has body parts or not?
func animalHasParts(animal byte, parts byte) bool {
	return animal&parts == parts
}

// Return blind animals map
func getBlindAnimals(animals map[string]byte) map[string]byte {
	filteredAnimals := map[string]byte{}
	for name, code := range animals {
		if !animalHasParts(code, EYE) {
			filteredAnimals[name] = code
		}
	}
	return filteredAnimals
}

// Return map of animals has tail and hand
func getAnimalsWithHandAndTail(animals map[string]byte) map[string]byte {
	filteredAnimals := map[string]byte{}
	for name, code := range animals {
		if animalHasParts(code, TAIL|HAND) {
			filteredAnimals[name] = code
		}
	}
	return filteredAnimals
}

func Bit() {
	animals := map[string]byte{
		"monkey":  HAND | TAIL | EYE,
		"human":   HAND | EYE,
		"octopus": EYE | TENTACLE,
		"snake":   EYE | TAIL,
		"mole":    TAIL | HAND,
	}

	for name, code := range animals {
		fmt.Printf("\"%[1]s\" binary code is: %[2]b\n", strings.Title(name), code)
	}

	for name := range getBlindAnimals(animals) {
		fmt.Printf("\"%[1]s\" is blind\n", strings.Title(name))
	}
	for name := range getAnimalsWithHandAndTail(animals) {
		fmt.Printf("\"%[1]s\" has hand and tail\n", strings.Title(name))
	}
}

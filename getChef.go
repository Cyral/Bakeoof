package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
)

// GetChef returns the chef name
func GetChef() string {
	return fmt.Sprintf("Chef %s", createLastName())
}

func createLastName() string {
	numParts := rand.Intn(4) + 1

	nameParts := nameParts()

	var buf bytes.Buffer
	for i := 0; i < numParts; i++ {
		idx := rand.Intn(len(nameParts))
		buf.WriteString(nameParts[idx])
	}

	return strings.Title(buf.String())
}

func nameParts() []string {
	return []string{
		"que",
		"way",
		"eel",
		"ruu",
		"tish",
		"yar",
		"uub",
		"iek",
		"oy",
		"pe",
		"ah",
		"sha",
		"dor",
		"feek",
		"goo",
		"moo",
		"rabi",
		"doku",
		"baca",
		"szum",
		"lan",
		"ski",
		"boy",
		"ardee",
		"gor",
		"don",
		"ram",
		"say",
		"guy",
		"fi",
		"eri",
	}
}

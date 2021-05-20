package main

import (
	"testing"

	// start 857
	"github.com/kr/pretty"
	// star 259, 但是我目前还是选用第二个，因为打印数组用的中括号
	prettyNow "github.com/kylelemons/godebug/pretty"
)

type Student struct {
	Name string
	Age  int
}
type ShipManifest struct {
	Name     string
	Crew     map[string]string
	Androids int
	Stolen   bool
	Arr      []Student
}

func TestPretty(t *testing.T) {
	manifest := &ShipManifest{
		Name: "Spaceship Heart of Gold",
		Crew: map[string]string{
			"Zaphod Beeblebrox": "Galactic President",
			"Trillian":          "Human",
			"Ford Prefect":      "A Hoopy Frood",
			"Arthur Dent":       "Along for the Ride",
		},
		Androids: 1,
		Stolen:   true,
		Arr:      []Student{{Name: "zs", Age: 100}, {Name: "zs", Age: 100}},
	}
	pretty.Log(manifest)
	//
	prettyNow.Print(manifest)
}

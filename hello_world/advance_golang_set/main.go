package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

func main() {
	fmt.Println("hello world")
	requiredClasses := mapset.NewSet()
	requiredClasses.Add("Cooking")
	requiredClasses.Add("English")
	requiredClasses.Add("Math")
	requiredClasses.Add("Biology")

	scienceSlice := []interface{}{"Biology", "Chemistry"}
	scienceClasses := mapset.NewSetFromSlice(scienceSlice)

	electiveClasses := mapset.NewSet()
	electiveClasses.Add("Welding")
	electiveClasses.Add("Music")
	electiveClasses.Add("Automotive")

	bonusClasses := mapset.NewSet()
	bonusClasses.Add("Go Programming")
	bonusClasses.Add("Python Programming")

	//Show me all the available classes I can take
	allClasses := requiredClasses.Union(scienceClasses).Union(electiveClasses).Union(bonusClasses)
	fmt.Println(allClasses)
	// Set{Cooking, English, Math, Chemistry, Welding, Biology, Music, Automotive, Go Programming, Python Programming}

}

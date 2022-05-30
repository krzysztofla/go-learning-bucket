package main

import (
	"fmt"
	"go-learning-bucket/models"
)

// Interfaces section ---------------------------------------------
type bot interface {
	doSomething(say string)
}

type talkingBot struct {
}

func (talkingBot) doSomething(whatToSay string) {
	fmt.Println(whatToSay)
}

func printTalkingBot(bb bot, asd string) {
	bb.doSomething(asd)
}

func intefacecsPractice() {
	tb := talkingBot{}

	printTalkingBot(tb, "hello world")
}

// slicesAndArrays section ---------------------------------------------
func slicesAndArrays() {
	p := models.NewPerson(" Krzysztof", " Łach", " Kraków", " Sezamkowa  18P")

	personPointer := &p // This means give me memory adress (reference) to object person so I can change something

	(*personPointer).UpdateName("krzysztof") // * pointer - means give me direct access to this object in the memory

	// (*personPointer) this means that pointer is turn into the value witch it this case is person instance

	// this will also work
	p.UpdateName("Krzysztof 22")

	// fmt.Printf("%+v", p)

	// Maps
	//   			key	  value
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	var colors2 map[string]string

	colors2["white"] = "#fffff"

	colors3 := make(map[string]string)
	colors3["white"] = "#fffff"

	delete(colors, "red")
	fmt.Println(colors)
}

func main() {
	//slicesAndArrays()
	intefacecsPractice()
}

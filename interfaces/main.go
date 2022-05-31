package main

import "fmt"

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

func main() {
	intefacecsPractice()
}

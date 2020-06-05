package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/yanzay/tbot/v2"
)
func init() { 
	rand.Seed(time.Now().Unix()) // random pick
}
var option = []string{"spock", "lizard", "scissors", "paper", "rock"} // list from where the bot chooses at random

func makeButtons() *tbot.InlineKeyboardMarkup {
	btnSpock := tbot.InlineKeyboardButton{ //spock button
		Text:         "Spock",// text printed on button
		CallbackData: "spock", //value associated with text
	}
	
	btnScissors := tbot.InlineKeyboardButton{ //scissor button
		Text:         "Scissors",// text printed on button
		CallbackData: "scissors", //value associated with text
	}
	btnLizard := tbot.InlineKeyboardButton{ //lizard button
		Text:         "Lizard",// text printed on button
		CallbackData: "lizard", //value associated with text
	}
	btnPaper := tbot.InlineKeyboardButton{ //paper button
		Text:         "Paper",// text printed on button
		CallbackData: "paper", //value associated with text
	}
	btnRock := tbot.InlineKeyboardButton{ //rock button
		Text:         "Rock", // text printed on button
		CallbackData: "rock", //value associated with text
	}
	
	return &tbot.InlineKeyboardMarkup{ //return in the inline keyboard
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			[]tbot.InlineKeyboardButton{btnSpock, btnPaper,btnLizard, btnScissors,btnRock,},
		},
	}
}
func (a *application) draw(player string) (output string) {
	var outcome string
	machine := option[rand.Intn(len(option))] // pick any random option for machine play
	switch player { //compares player and bot move
	case mapz[machine]: // compare from mapz
		outcome = "lose"
		a.losses++
	case machine: //if equal then draw
		outcome = "draw"
		a.draws++ //increment draw counter
	default: // otherwise bot wins
		outcome = "win"
		a.wins++
	}
	output = fmt.Sprintf("The bot picked %s!! You %s.",  machine, outcome) //result output
	return
}
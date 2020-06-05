package main

import (
	"fmt"
	"github.com/yanzay/tbot/v2" //importing library - we are using modules so v2
)
func (a *application) sthd(m *tbot.Message) { //handles the start. This message will pop up during the start
	output := "Welcome to stone paper scissor lizard spock bot. Enjoy a game with this bot here.\nThe instructions are:\n1. /play is to play.\n2./score is to see your score against the bot.\n3. If you lost a lot, you can reset using  /reset."
	a.client.SendMessage(m.Chat.ID, output) //sends the starting message
}
//added functionailty
func (a *application) schd(m *tbot.Message) { // keeps track of the score
	output := fmt.Sprintf("Score:\nNumber of Losses: %v\nNumber of Draws: %v\nNumber of Wins: %v", a.losses, a.draws, a.wins)
	a.client.SendMessage(m.Chat.ID, output) //output to chat
}
func (a *application) plhd(m *tbot.Message) { //handles play
	buttons := makeButtons() //create button for play 
	a.client.SendMessage(m.Chat.ID, "Choose your weapon", tbot.OptInlineKeyboardMarkup(buttons)) // buttons come up near keyboard
}
func (a *application) cbh(cq *tbot.CallbackQuery) {
	player := cq.Data //input from player
	output := a.draw(player)
	a.client.DeleteMessage(cq.Message.Chat.ID, cq.Message.MessageID)
	a.client.SendMessage(cq.Message.Chat.ID, output)
}
//added functionality
func (a *application) rst(m *tbot.Message) {
	a.draws, a.losses, a.wins = 0, 0, 0 //resets all to 0 
	output := "The game has been reset"
	a.client.SendMessage(m.Chat.ID, output) //outputs to chat
}


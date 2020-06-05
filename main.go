package main

import (
	"log"

	"github.com/yanzay/tbot/v2"
)

type score struct {
	wins, draws, losses uint
}

type application struct {
	client *tbot.Client
	score
}

var (
	app     application
	bot     *tbot.Server
	options = map[string]string{
		// choice : beats
		"paper":    "rock",
		"rock":     "scissors",
		"scissors": "paper",
	}
)

const token = "897697024:AAFGfJ11_uop4gPo1YmnCzguERCfliNH4Tw"


func main() {
	bot = tbot.New(token)
	app.client = bot.Client()
	bot.HandleMessage("/start", app.startHandler)
	bot.HandleMessage("/play", app.playHandler)
	bot.HandleMessage("/score", app.scoreHandler)
	bot.HandleMessage("/reset", app.resetHandler)
	bot.HandleCallback(app.callbackHandler)
	log.Fatal(bot.Start())
}
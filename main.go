package main

import (
	"os"
	"log"
	"github.com/yanzay/tbot/v2"
	"github.com/joho/godotenv"
)



type application struct {
	client *tbot.Client
	
}

var (
	app     application
	bot     *tbot.Server
	token   string
	
)

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	token = os.Getenv("TELEGRAM_TOKEN")
}

func main() {
	bot = tbot.New(token)
	app.client = bot.Client()
	bot.HandleMessage("/start", app.startHandler)
	log.Fatal(bot.Start())
}

func (a *application) startHandler(m *tbot.Message) {
	msg:= " This is a bot"
	a.client.SendMessage(m.chat.ID, msg)
}
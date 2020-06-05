package main
import (
	"log"
	"github.com/yanzay/tbot/v2"
)

type score struct { //types of outcomes
	losses, draws, wins uint
}

type application struct { 
	client *tbot.Client //chatbot
	score 
}

var (
	app     application
	bot     *tbot.Server
	mapz = map[string]string{ //map of what beats what
		"scissors": "lizard", //scissors 
		"spock": " rock", //spoc vaporizes rock
		"lizard": "spock", //lizard poisons spock
		"lizard": " paper", //lizard eats paper
		"paper":    "rock", //paper covers rock
		"spock": "scissors", //spock smashes scissors
		"paper": "spock", // paper disproves spock	
		"rock":     "scissors", //rocks crushes scissors
		"rock":  "lizard", //rock crushes lizard
		"scissors": "paper", //scissors cut paper
			}
)

const token = "897697024:AAFGfJ11_uop4gPo1YmnCzguERCfliNH4Tw" //telegram token
func main() {
	 //implementing a chatbot with telegram api token
	bot = tbot.New(token)
	 //new chatbot created
	app.client = bot.Client() 
	bot.HandleMessage("/reset", app.rst)// handles reset
	bot.HandleMessage("/start", app.sthd) //handles starting chatbot
	bot.HandleMessage("/score", app.schd) //handles score
	bot.HandleMessage("/play", app.plhd)  // handles starting game
	bot.HandleCallback(app.cbh) //button press
	log.Fatal(bot.Start()) //start bot
}
package main

import (
	"log"
	"os"
	"strconv"
	"telegobot/keyboard"
	"telegobot/urlstruct"

	"github.com/joho/godotenv"
)

func Start(incomingMessages urlstruct.IncomingMessage, ba urlstruct.BotApi) {

	for _, m := range incomingMessages.Result {

		messageText := m.Message.Text
		if messageText == "/start" {

			var newKeyboard keyboard.Keyboard
			newKeyboard.ByDefault()
			newKeyboard.AddButtonRequestContact(`Отправить номер`)

			var mess urlstruct.Message
			mess.ChatID = m.Message.From.Id
			mess.Text = `Добрый день, уважаемые коллеги! Для получения доступа к функциям чат-бота, потвердите личность, нажав на кнопку "Отправить номер телефона"`
			mess.AddKeyboard(newKeyboard)
			ba.SendMessage(mess)
		} else if m.Type == "Contact" {
			urlstruct.GetЕmployeesData(m.Message.Contact.Phone_number, strconv.Itoa(m.Message.From.Id))
		}

	}
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	var botApi urlstruct.BotApi
	botApi.TeleToken = os.Getenv("teleToken")
	botApi.ByDefault()
	botApi.SetStartFunction(Start)
	botApi.RunLongPolling()
}

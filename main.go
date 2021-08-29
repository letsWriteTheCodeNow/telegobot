package main

import (
	"fmt"
	"log"
	"os"
	"telegobot/keyboard"
	"telegobot/urlstruct"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	var botApi urlstruct.BotApi
	botApi.ByDefault()
	botApi.TeleToken = os.Getenv("teleToken")

	for true {

		// urlGetUpdates := "https://api.telegram.org/bot" + teleToken + "/getUpdates?timeout=15"
		// if lastMessage != 0 {
		// 	urlGetUpdates = urlGetUpdates + "&offset=" + strconv.Itoa(lastMessage+1)
		// }
		// resp, err := http.Get(urlGetUpdates)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// defer resp.Body.Close()
		// log.Println(string(body))

		// var incomingMessages keyboard.IncomingMessage
		incomingMessages := botApi.GetUpdates()
		// json.Unmarshal([]byte(body), &incomingMessages)

		for _, message := range incomingMessages.Result {

			botApi.LastMessage = message.Update_id + 1

			fmt.Println(message)
			messageText := message.Message.Text
			if messageText == "/start" {

				// 		var keyboardButton keyboard.KeyboardButton
				// 		var keyboardButtonArray []keyboard.KeyboardButton

				// 		keyboardButton.Text = "Отправить номер"
				// 		keyboardButton.Request_contact = true

				// 		keyboardButtonArray = append(keyboardButtonArray, keyboardButton)

				var newKeyboard keyboard.Keyboard
				newKeyboard.ByDefault()
				messageText = `Добрый день, уважаемые коллеги! Для получения доступа к функциям чат-бота, потвердите личность, нажав на кнопку "Отправить номер телефона"`
				newKeyboard.AddButtonRequestContact(messageText)
				// 		keyboardStruct.KeyboardButtonArray = append(keyboardStruct.KeyboardButtonArray, keyboardButtonArray)
				// 		keyboardStruct.Resize_keyboard = true
				// 		keyboardStruct.One_time_keyboard = true
				// 		json_data, err := json.Marshal(keyboardStruct)
				// 		if err != nil {
				// 			log.Fatalln(err)
				// 		}

				// 		println(json_data)
				// 		replyMarkupText = "&reply_markup=" + string(json_data)
				// 		println(replyMarkupText)
				// 	} else if message.Message.Contact.Phone_number != "" {

				// 	}
				// 	// s := fmt.Sprintf("%s is a %s Portal.\n", name, dept)
				// 	urlSendMessage := "https://api.telegram.org/bot" + teleToken + "/sendMessage?chat_id=" + strconv.Itoa(message.Message.From.Id) + "&text=" + messageText + replyMarkupText
				// 	_, err := http.Get(urlSendMessage)
				// 	if err != nil {
				// 		log.Fatalln(err)
			}
		}
		// // a := result["ok"]a.(data)
		// fmt.Println(incomingMessages)
	}

}

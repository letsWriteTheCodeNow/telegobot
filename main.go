package main

import (
	"log"
	"os"
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
			// var newKeyboard keyboard.Keyboard
			// newKeyboard.ByDefault()
			// messageText = `Добрый день, уважаемые коллеги! Для получения доступа к функциям чат-бота, потвердите личность, нажав на кнопку "Отправить номер телефона"`
			// newKeyboard.AddButtonRequestContact(messageText)
			// urlSendMessage := "https://api.telegram.org/bot" + teleToken + "/sendMessage?chat_id=" + strconv.Itoa(message.Message.From.Id) + "&text=" + messageText + replyMarkupText
			// // 			_, err := http.Get(urlSendMessage)
			// replyMarkupText = "&reply_markup=" + string(json_data)
			// if err != nil {
			// 	log.Fatalln(err)
			// }
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

// for true {

// 	incomingMessages := botApi.GetUpdates()
// 	// botApi.FuncStart = botApi.FuncStart(Start(incincomingMessages))
// 	// json.Unmarshal([]byte(body), &incomingMessages)

// 	for _, message := range incomingMessages.Result {

// 		botApi.LastMessage = message.Update_id + 1

// 		fmt.Println(message)
// 		messageText := message.Message.Text
// 		if messageText == "/start" {

// 			// 		var keyboardButton keyboard.KeyboardButton
// 			// 		var keyboardButtonArray []keyboard.KeyboardButton

// 			// 		keyboardButton.Text = "Отправить номер"
// 			// 		keyboardButton.Request_contact = true

// 			// 		keyboardButtonArray = append(keyboardButtonArray, keyboardButton)

// 			var newKeyboard keyboard.Keyboard
// 			newKeyboard.ByDefault()
// 			messageText = `Добрый день, уважаемые коллеги! Для получения доступа к функциям чат-бота, потвердите личность, нажав на кнопку "Отправить номер телефона"`
// 			newKeyboard.AddButtonRequestContact(messageText)
// 			// 		keyboardStruct.KeyboardButtonArray = append(keyboardStruct.KeyboardButtonArray, keyboardButtonArray)
// 			// 		keyboardStruct.Resize_keyboard = true
// 			// 		keyboardStruct.One_time_keyboard = true
// 			// 		json_data, err := json.Marshal(keyboardStruct)
// 			// 		if err != nil {
// 			// 			log.Fatalln(err)
// 			// 		}

// 			// 		println(json_data)
// 			// 		replyMarkupText = "&reply_markup=" + string(json_data)
// 			// 		println(replyMarkupText)
// 			// 	} else if message.Message.Contact.Phone_number != "" {

// 			// 	}
// 			// 	// s := fmt.Sprintf("%s is a %s Portal.\n", name, dept)
// 			// 	urlSendMessage := "https://api.telegram.org/bot" + teleToken + "/sendMessage?chat_id=" + strconv.Itoa(message.Message.From.Id) + "&text=" + messageText + replyMarkupText
// 			// 	_, err := http.Get(urlSendMessage)
// 			// 	if err != nil {
// 			// 		log.Fatalln(err)
// 		}
// 	}
// 	// // a := result["ok"]a.(data)
// fmt.Println(incomingMessages)
// }

// }

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

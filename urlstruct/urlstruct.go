package urlstruct

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"telegobot/keyboard"
	"time"
)

type IncomingMessage struct {
	Ok     bool `json:"ok"`
	Result []struct {
		Update_id int `json:"update_id"`
		Message   struct {
			Message_id int `json:"message_id"`
			From       struct {
				Id int `json:"id"`
			} `json:"from"`
			Text    string `json:"text"`
			Contact struct {
				Phone_number string `json:"phone_number"`
			} `json:"contact"`
			Entities []struct {
				Type string `json:"type"`
			} `json:"entities"`
		} `json:"message"`
		HandlerFunction struct {
			Name string `json:"data"`
		} `json:"callback_query"`
		Type string
	} `json:"result"`
}

type Message struct {
	ChatID      int
	Text        string
	ReplyMarkup string
}

func (m *Message) AddKeyboard(k keyboard.Keyboard) {

	json_data, err := json.Marshal(k)
	if err != nil {
		log.Fatalln(err)
	}
	m.ReplyMarkup = string(json_data)

}

type BotApi struct {
	Host               string
	GetApdateAddress   string
	Timeout            int
	LastMessage        int
	Offset             int
	ReplyMarkup        string
	TeleToken          string
	SendMessageAddress string
	Client             http.Client
	FuncStart          reflect.Value
}

func (u *BotApi) ByDefault() {
	u.Host = "https://api.telegram.org/bot"
	u.GetApdateAddress = "/getUpdates?timeout="
	u.SendMessageAddress = "/sendMessage?chat_id="
	u.Timeout = 120
	u.LastMessage = 0
	u.Offset = 0
	u.Client.Timeout = 120 * time.Second
}

func (ba *BotApi) RunLongPolling() {

	for true {

		incomingMessages := ba.GetUpdates()

		for imess, message := range incomingMessages.Result {
			for _, entity := range message.Message.Entities {
				if entity.Type == "bot_command" {
					incomingMessages.Result[imess].Type = "Command"
				}
			}
			if message.Message.Contact.Phone_number != "" {
				incomingMessages.Result[imess].Type = "Contact"
			}
			ba.LastMessage = message.Update_id + 1
		}

		inValue := make([]reflect.Value, 2)
		inValue[0] = reflect.ValueOf(incomingMessages)
		inValue[1] = reflect.ValueOf(*ba)
		ba.FuncStart.Call(inValue)

	}
}

func (ba *BotApi) SetStartFunction(startFunction func(incomingMessages IncomingMessage, b BotApi)) {

	ba.FuncStart = reflect.ValueOf(startFunction)

}

// func (ba *BotApi) SendKeyboardButton(startFunction func(incomingMessages IncomingMessage, b BotApi)) {
// 	ba.
// }

func (u *BotApi) GetUpdates() IncomingMessage {

	urlGetUpdates := u.Host + u.TeleToken + u.GetApdateAddress + strconv.Itoa(u.Timeout)
	if u.LastMessage != 0 {
		urlGetUpdates = urlGetUpdates + "&offset=" + strconv.Itoa(u.LastMessage)
	}

	resp, err := u.Client.Get(urlGetUpdates)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var incomingMessages IncomingMessage
	json.Unmarshal([]byte(body), &incomingMessages)

	fmt.Println(string(body))
	return incomingMessages

}

func (ba *BotApi) SendMessage(m Message) {

	urlSendMessage := ba.Host + ba.TeleToken + ba.SendMessageAddress + strconv.Itoa(m.ChatID)
	if m.Text != "" {
		urlSendMessage = urlSendMessage + "&text=" + m.Text
	}
	if m.ReplyMarkup != "" {
		urlSendMessage = urlSendMessage + "&reply_markup=" + m.ReplyMarkup
	}
	resp, err := ba.Client.Get(urlSendMessage)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	// urlSendMessage := "https://api.telegram.org/bot" + teleToken + "/sendMessage?chat_id=" + strconv.Itoa(message.Message.From.Id) + "&text=" + messageText + replyMarkupText
	// // if ba.ReplyMarkup != "" {
	// // 	urlSendMessage = urlSendMessage + "&reply_markup=" +
	// // }

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer resp.Body.Close()

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
	// var incomingMessages IncomingMessage
	// json.Unmarshal([]byte(body), &incomingMessages)

	// return incomingMessages
}

// messageText := message.Message.Text
// if messageText == "/start" {

// 	var newKeyboard keyboard.Keyboard
// 	newKeyboard.ByDefault()
// 	messageText = `Добрый день, уважаемые коллеги! Для получения доступа к функциям чат-бота, потвердите личность, нажав на кнопку "Отправить номер телефона"`
// 	newKeyboard.AddButtonRequestContact(messageText)
// }

// func GetURLStructByDefault() urlSrtuct {

// 	var urlSrtuct urlSrtuct
// 	urlSrtuct.host = host
// 	urlSrtuct.getApdateAddress = getUpdate
// 	urlSrtuct.timeout = 120
// 	urlSrtuct.lastMessage = 0
// 	urlSrtuct.offset = 0

// 	return urlSrtuct
// }

// func RunByDefault(teleToken string) {

// 	urlstruct := GetURLStructByDefault()

// 	for true {

// 		urlGetUpdates := urlstruct.host + teleToken + urlstruct.getApdateAddress + "?timeout=" + strconv.Itoa(urlstruct.timeout)
// 		if urlstruct.lastMessage != 0 {
// 			urlGetUpdates = urlGetUpdates + "&offset=" + strconv.Itoa(urlstruct.lastMessage+1)
// 		}
// 		resp, err := http.Get(urlGetUpdates)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		body, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		defer resp.Body.Close()
// 		log.Println(string(body))

// 		var incomingMessages keyboard.IncomingMessage

// 		json.Unmarshal([]byte(body), &incomingMessages)

// 		for _, message := range incomingMessages.Result {

// 			lastMessage = message.Update_id
// 			fmt.Println(message)
// 			messageText := message.Message.Text
// 			if messageText == "/start" {

// 				messageText = `Добрый день, уважаемые коллеги! Для получения доступа к функциям чат-бота, потвердите личность, нажав на кнопку "Отправить номер телефона"`

// 				var keyboardButton keyboard.KeyboardButton
// 				var keyboardButtonArray []keyboard.KeyboardButton

// 				keyboardButton.Text = "Отправить номер"
// 				keyboardButton.Request_contact = true

// 				keyboardButtonArray = append(keyboardButtonArray, keyboardButton)

// 				var keyboardStruct keyboard.Keyboard
// 				keyboardStruct.KeyboardButtonArray = append(keyboardStruct.KeyboardButtonArray, keyboardButtonArray)
// 				keyboardStruct.Resize_keyboard = true
// 				keyboardStruct.One_time_keyboard = true
// 				json_data, err := json.Marshal(keyboardStruct)
// 				if err != nil {
// 					log.Fatalln(err)
// 				}

// 				println(json_data)
// 				replyMarkupText = "&reply_markup=" + string(json_data)
// 				println(replyMarkupText)
// 			} else if message.Message.Contact.Phone_number != "" {

// 			}
// 			// s := fmt.Sprintf("%s is a %s Portal.\n", name, dept)
// 			urlSendMessage := "https://api.telegram.org/bot" + teleToken + "/sendMessage?chat_id=" + strconv.Itoa(message.Message.From.Id) + "&text=" + messageText + replyMarkupText
// 			_, err := http.Get(urlSendMessage)
// 			if err != nil {
// 				log.Fatalln(err)
// 			}
// 		}
// 		// a := result["ok"]a.(data)
// 		fmt.Println(incomingMessages)
// 	}
// }

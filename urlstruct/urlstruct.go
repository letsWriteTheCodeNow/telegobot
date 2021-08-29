package urlstruct

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"telegobot/keyboard"
// )

// const (
// 	host      = "https://api.telegram.org/bot"
// 	getUpdate = "/getUpdates" //?timeout=15
// )
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
		} `json:"message"`
		HandlerFunction struct {
			Name string `json:"data"`
		} `json:"callback_query"`
	} `json:"result"`
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
	Client             *http.Client
}

func (u *BotApi) ByDefault() {
	u.Host = "https://api.telegram.org/bot"
	u.GetApdateAddress = "/getUpdates"
	u.SendMessageAddress = "/sendMessage"
	u.Timeout = 120
	u.LastMessage = 0
	u.Offset = 0
	u.Client.Timeout = 120
}

func (u *BotApi) GetUpdates() IncomingMessage {

	urlGetUpdates := u.Host + u.TeleToken + u.GetApdateAddress + "?timeout=" + strconv.Itoa(u.Timeout)
	if u.LastMessage != 0 {
		urlGetUpdates = urlGetUpdates + "&offset=" + strconv.Itoa(u.LastMessage+1)
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

	return incomingMessages

}

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

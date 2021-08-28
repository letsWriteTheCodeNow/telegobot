package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type butt struct {
	Text          string `json:"text"`
	Callback_data string `json:"callback_data"`
}

type keyboardButton struct {
	Text            string `json:"text"`
	Request_contact bool   `json:"request_contact"`
}

type inli struct {
	ButtonType        [][]keyboardButton `json:"keyboard"`
	Resize_keyboard   bool               `json:"resize_keyboard"`
	One_time_keyboard bool               `json:"one_time_keyboard"`
	// ButtonData []struct {
	// 	Text            string `json:"text"`
	// 	handlerFunction string `json:"callback_data"`
	// }
}

type inlineKeyboard struct {
	ButtonType [][]butt `json:"inline_keyboard"`
	// ButtonData []struct {
	// 	Text            string `json:"text"`
	// 	handlerFunction string `json:"callback_data"`
	// }
}

// func addNewInlineKey(url string) {

// }

type incomingMessage struct {
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

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	teleToken := os.Getenv("teleToken")
	// textOffset := "&offset="
	lastMessage := 0
	replyMarkupText := ""
	for true {

		urlGetUpdates := "https://api.telegram.org/bot" + teleToken + "/getUpdates?timeout=15"
		if lastMessage != 0 {
			urlGetUpdates = urlGetUpdates + "&offset=" + strconv.Itoa(lastMessage+1)
		}
		resp, err := http.Get(urlGetUpdates)
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		log.Println(string(body))

		var incomingMessages incomingMessage

		json.Unmarshal([]byte(body), &incomingMessages)

		for _, message := range incomingMessages.Result {

			// switch message.HandlerFunction.Name: {
			// case: "/start":

			// }
			switch message.HandlerFunction.Name {
			case "getAPhoneNumber":

				// miValue := reflect.ValueOf(СallFunc{})
				// reflect.ValueOf(&СallFunc{}).MethodByName("Add").Call([]reflect.Value{})
				fmt.Println(1)

				// var b СallFunc
				// method()
				// var i interface
				// ptr = reflect.New(reflect.TypeOf(i))
				// method := ptr.MethodByName(message.HandlerFunction.Name)
				// continue
			}

			lastMessage = message.Update_id
			fmt.Println(message)
			messageText := message.Message.Text
			if messageText == "/start" {

				messageText = `Добрый день, уважаемые коллеги! Для получения доступа к функциям чат-бота, потвердите личность, нажав на кнопку "Отправить номер телефона"`

				var inK inli
				var keyboardButtonGetPhone keyboardButton
				var keyboardButtonGetPhoneArray []keyboardButton

				keyboardButtonGetPhone.Text = "Отправить номер"
				keyboardButtonGetPhone.Request_contact = true
				keyboardButtonGetPhoneArray = append(keyboardButtonGetPhoneArray, keyboardButtonGetPhone)
				inK.ButtonType = append(inK.ButtonType, keyboardButtonGetPhoneArray)
				inK.Resize_keyboard = true
				inK.One_time_keyboard = true
				json_data, err := json.Marshal(inK)
				if err != nil {
					log.Fatalln(err)
				}

				println(json_data)
				replyMarkupText = "&reply_markup=" + string(json_data)
				println(replyMarkupText)
			} else if message.Message.Contact.Phone_number != "" {

			}
			// s := fmt.Sprintf("%s is a %s Portal.\n", name, dept)
			urlSendMessage := "https://api.telegram.org/bot" + teleToken + "/sendMessage?chat_id=" + strconv.Itoa(message.Message.From.Id) + "&text=" + messageText + replyMarkupText
			_, err := http.Get(urlSendMessage)
			if err != nil {
				log.Fatalln(err)
			}
		}
		// a := result["ok"]a.(data)
		fmt.Println(incomingMessages)
	}

}

// &reply_markup= ""inline_keyboard"": [
// 	|		[{
// 	|                ""text"": ""Обычная кнопка"",
// 	|                ""callback_data"": ""ОтветНаСообщение1""
// 	|            }
// 	|        ],
// if strings.Contains(messageText, "переводы") {
// 	messageText = "Да, слышали что - то о переводах"
// }
// switch messageText {

// case "Привет":
// 	messageText = "Ну привет!"
// case "Как дела?":
// 	messageText = "Хорошо, а у тебя?"
// case "тоже":
// 	messageText = "ну и отлично?"
// case strings.Contains(messageText, "переводы"):
// 	messageText = "ну и отлично?"
// }
// АдресЗапроса =
// 	"bot"
// +МойToken
// +"/sendMessage"
// +"?chat_id="
// +ЧатID
// +"&text="
// +ТекстСообщения
// log.Println(result)
// log.Println(result["data"])
// https: //api.telegram.org/bot123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11/getMe

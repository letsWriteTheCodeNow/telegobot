package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"telegobot/keyboard"

	"github.com/joho/godotenv"
)

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

		var incomingMessages keyboard.IncomingMessage

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

				var keyboardButton keyboard.KeyboardButton
				var keyboardButtonArray []keyboard.KeyboardButton

				keyboardButton.Text = "Отправить номер"
				keyboardButton.Request_contact = true

				keyboardButtonArray = append(keyboardButtonArray, keyboardButton)

				var keyboardStruct keyboard.Keyboard
				keyboardStruct.KeyboardButtonArray = append(keyboardStruct.KeyboardButtonArray, keyboardButtonArray)
				keyboardStruct.Resize_keyboard = true
				keyboardStruct.One_time_keyboard = true
				json_data, err := json.Marshal(keyboardStruct)
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

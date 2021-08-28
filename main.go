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

type incomingMessage struct {
	Ok     bool `json:"ok"`
	Result []struct {
		Update_id int `json:"update_id"`
		Message   struct {
			Message_id int `json:"message_id"`
			From       struct {
				Id int `json:"id"`
			} `json:"from"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"result"`
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	teleToken := os.Getenv("teleToken")
	// textOffset := "&offset="
	lastMessage := 0
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
			lastMessage = message.Update_id
			fmt.Println(message)
			messageText := message.Message.Text

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
			urlSendMessage := "https://api.telegram.org/bot" + teleToken + "/sendMessage?chat_id=" + strconv.Itoa(message.Message.From.Id) + "&text=" + messageText
			_, err := http.Get(urlSendMessage)
			if err != nil {
				log.Fatalln(err)
			}
		}
		// a := result["ok"]a.(data)
		fmt.Println(incomingMessages)
	}
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

}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type incomingMessage struct {
	Ok     bool `json:"ok"`
	Result []struct {
		Update_id int `json:"update_id"`
		Message   struct {
			Message_id int `json:"message_id"`
			From       struct {
				Id     int  `json:"id"`
				Is_bot bool `json:"is_bot"`
			} `json:"from"`
		} `json:"message"`
	} `json:"result"`
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	teleToken := os.Getenv("teleToken")
	// textOffset := "&offset="
	// lastMessage := ""
	for true {
		url := "https://api.telegram.org/bot" + teleToken + "/getUpdates?timeout=300"
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		log.Println(string(body))

		var result map[string]interface{}
		var result1 incomingMessage
		// if err := json.NewDecoder([]byte(body).Decode(&result)//; err != nil {
		// 	log.Fatal("ooopsss! an error occurred, please try again")
		// }

		// json.NewDecoder(resp.Body).Decode(&result)result["ok"].(data)
		json.Unmarshal([]byte(body), &result)
		json.Unmarshal([]byte(body), &result1)
		log.Println(result)
		for i, v := range result {
			fmt.Printf("2**%d = %d\n", i, v)
		}
		// a := result["ok"]a.(data)
		// fmt.Println(a)
	}

	// log.Println(result)
	// log.Println(result["data"])
	// https: //api.telegram.org/bot123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11/getMe

}

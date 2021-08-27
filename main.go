package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	teleToken := os.Getenv("teleToken")
	url := "https://api.telegram.org/bot" + teleToken + "/getUpdates" + "?timeout=300"
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

	// var result map[string]interface{}

	// json.NewDecoder(resp.Body).Decode(&result)

	// log.Println(result)
	// log.Println(result["data"])
	// https: //api.telegram.org/bot123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11/getMe

}

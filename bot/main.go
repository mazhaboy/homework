package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	botToken := "1425310108:AAEYKRVIgoUx59Ke9-FoJ7SruZ1p_czLz18"
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
	offset := 0
	for {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Fatal(err)
		}
		for _, update := range updates {
			err = respond(botUrl, update)
			offset = update.UpdateID + 1
		}
		fmt.Println(updates)
	}
}

func getUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var restResponse RestResponse

	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return restResponse.Result, nil

}

func respond(botUrl string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatID = update.Message.Chat.ChatID

	if update.Message.Text == "Привет" {
		botMessage.Text = "Привет\n" + update.Message.Text
	}
	if update.Message.Text == "Как тебя зовут?\n" {
		botMessage.Text = "Меня зовут EchoBot\n" + update.Message.Text
	}


	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	return nil
}

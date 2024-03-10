package main

import (
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()
	client := &http.Client{}
	bot, err := messaging_api.NewMessagingApiAPI(
		os.Getenv("LINE_CHANNEL_TOKEN"),
		messaging_api.WithHTTPClient(client),
	)
	if err != nil {
		panic(err)
	}

	req := messaging_api.PushMessageRequest{
		To: "Ub4d5acd9de7b1ed43b06190e596eb2b2",
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Text: "こんにちは。",
			},
		},
	}
	bot.PushMessage(&req, "")
}

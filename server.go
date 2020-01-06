package botton

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/translate"
	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func ReplayEnglish(w http.ResponseWriter, req *http.Request) {
	bot, err := linebot.New(
		"CHANNEL_SECRET",
		"CHANNEL_TOKEN",
	)

	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(req)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				response, err := translating(message.Text)
				if err != nil {
					log.Fatal(err)
				}
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(response)).Do(); err != nil {
					log.Print(err)
				}
				w.WriteHeader(200)
			}
		}
	}
}

func translating(text string) (string, error) {
	ctx := context.Background()

	//環境変数に設定したAPIキーを取得
	apiKey := os.Getenv("APIKEY")
	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, language.English, nil)
	if err != nil {
		return "", err
	}

	return resp[0].Text, nil
}

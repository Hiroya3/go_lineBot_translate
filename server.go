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
		//The LINE Messaging API defines 7 types of event
		//EventTypeMessage, EventTypeFollow, EventTypeUnfollow, EventTypeJoin, EventTypeLeave, EventTypePostback, EventTypeBeacon
		if event.Type == linebot.EventTypeMessage {
			//参考：https://developers.line.biz/ja/reference/messaging-api/#message-event
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

	//respの内容は[]Translation{Text string,Source language.Tag,Model string}
	//https://godoc.org/cloud.google.com/go/translate#Translation
	resp, err := client.Translate(ctx, []string{text}, language.English, nil)
	if err != nil {
		return "", err
	}

	//convertStrが複数種類入っている可能性があるので
	//if文を使用
	if strings.Contains(resp[0].Text, convertStr[0]) {
		resp[0].Text = strings.Replace(resp[0].Text, convertStr[0], "<", -1)
	}

	if strings.Contains(resp[0].Text, convertStr[1]) {
		resp[0].Text = strings.Replace(resp[0].Text, convertStr[1], ">", -1)
	}

	if strings.Contains(resp[0].Text, convertStr[2]) {
		resp[0].Text = strings.Replace(resp[0].Text, convertStr[2], "&", -1)
	}

	if strings.Contains(resp[0].Text, convertStr[3]) {
		resp[0].Text = strings.Replace(resp[0].Text, convertStr[3], "\"", -1)
	}

	if strings.Contains(resp[0].Text, convertStr[4]) {
		resp[0].Text = strings.Replace(resp[0].Text, convertStr[4], "'", -1)
	}

	if strings.Contains(resp[0].Text, convertStr[5]) {
		resp[0].Text = strings.Replace(resp[0].Text, convertStr[5], " ", -1)
	}
	
	if strings.Contains(resp[0].Text, "』") {
		resp[0].Text = strings.Replace(resp[0].Text, "』", "\"", -1)
	}

	return resp[0].Text, nil
}

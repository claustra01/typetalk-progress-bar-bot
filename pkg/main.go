package main

import (
	"os"

	"github.com/claustra01/typetalk-progress-bar-bot/pkg/bot"
)

func main() {
	topicId := os.Getenv("TYPETALK_TOPIC_ID")
	if topicId == "" {
		panic("TYPETALK_TOPIC_ID is not set")
	}

	token := os.Getenv("TYPETALK_TOKEN")
	if token == "" {
		panic("TYPETALK_TOKEN is not set")
	}

	filename := bot.GenerateImage()
	fileKey := bot.UploadImage(topicId, token, filename)
	bot.PostMessage(topicId, token, fileKey)

	// s := gocron.NewScheduler(time.UTC)
	// s.Every(1).Day().At("00:00").Do(post)
	// s.StartBlocking()
}

package main

import (
	"fmt"
	"log"

	"github.com/Totus-Floreo/tg-pwd-manager/config"
	"github.com/Totus-Floreo/tg-pwd-manager/pkg/application"
	"go.uber.org/zap"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func main() {

	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer zapLogger.Sync()
	logger := zapLogger.Sugar()

	bot, err := telego.NewBot(config.APItoken /* or os.Getenv("APItoken")*/, telego.WithDefaultDebugLogger())
	if err != nil {
		logger.Fatalln(err)
	}

	commands := application.NewCommandParams()

	if err := bot.SetMyCommands(commands); err != nil {
		logger.Fatalln(err)
	}

	err = bot.SetWebhook(&telego.SetWebhookParams{
		URL: "https://d01b-46-138-39-206.ngrok-free.app/bot" + bot.Token(),
	})
	if err != nil {
		logger.Fatalln(err)
	}

	info, err := bot.GetWebhookInfo()
	fmt.Printf("Webhook Info: %+v\n", info)
	if err != nil {
		logger.Fatalln(err)
	}

	updates, err := bot.UpdatesViaWebhook("/bot" + bot.Token())
	if err != nil {
		logger.Fatalln(err)
	}

	botHandler, err := th.NewBotHandler(bot, updates)
	if err != nil {
		logger.Fatalln(err)
	}

	botHandler.Handle()

	go func() {
		err = bot.StartWebhook("localhost:8443")
		if err != nil {
			logger.Fatalln(err)
		}
	}()

	defer func() {
		err = bot.StopWebhook()
		if err != nil {
			logger.Fatalln(err)
		}
	}()

}

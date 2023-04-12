package main

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/config"
	"github.com/KuratovIgor/gram-work-bot/pkg/repository/graphqldb"
	"github.com/KuratovIgor/gram-work-bot/pkg/server"
	"github.com/KuratovIgor/gram-work-bot/pkg/telegram"
	headhunter "github.com/KuratovIgor/head_hunter_sdk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/machinebox/graphql"
	"log"
)

func main() {
	cfg, err := config.Init()

	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	graphqlClient := graphql.NewClient("https://uk.api.8base.com/clbmch9qo002v08lfal9zccgy")
	graphQlRepository := graphqldb.NewGraphqlRepository(graphqlClient)

	headhunterClient, err := headhunter.NewClient(cfg.ClientID, cfg.ClientSecret, cfg.RedirectURI)
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot, cfg, headhunterClient, cfg.Messages, graphQlRepository, cfg.LkUrl)

	authorizationServer := server.NewAuthorizationServer(graphQlRepository, "https://t.me/gram_work_bot", headhunterClient)

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := authorizationServer.Start(); err != nil {
		log.Fatal(err)
	}
}

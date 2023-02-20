package main

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/config"
	"github.com/KuratovIgor/gram-work-bot/pkg/repository"
	"github.com/KuratovIgor/gram-work-bot/pkg/repository/boltdb"
	"github.com/KuratovIgor/gram-work-bot/pkg/repository/graphqldb"
	"github.com/KuratovIgor/gram-work-bot/pkg/server"
	"github.com/KuratovIgor/gram-work-bot/pkg/telegram"
	headhunter "github.com/KuratovIgor/head_hunter_sdk"
	"github.com/boltdb/bolt"
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

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	tokenRepository := boltdb.NewTokenRepository(db)

	graphqlClient := graphql.NewClient("https://uk.api.8base.com/clbmch9qo002v08lfal9zccgy")
	graphQlRepository := graphqldb.NewGraphqlRepository(graphqlClient)

	headhunterClient, err := headhunter.NewClient(cfg.ClientID, cfg.ClientSecret, cfg.RedirectURI)
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot, headhunterClient, cfg.Messages, tokenRepository, graphQlRepository)

	authorizationServer := server.NewAuthorizationServer(tokenRepository, "https://t.me/gram_work_bot", headhunterClient)

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := authorizationServer.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return err
		}

		_, err2 := tx.CreateBucketIfNotExists([]byte(repository.RefreshToken))
		if err != nil {
			return err2
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}

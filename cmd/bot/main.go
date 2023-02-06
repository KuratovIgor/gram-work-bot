package main

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/config"
	"github.com/KuratovIgor/gram-work-bot/pkg/repository"
	"github.com/KuratovIgor/gram-work-bot/pkg/repository/boltdb"
	"github.com/KuratovIgor/gram-work-bot/pkg/server"
	"github.com/KuratovIgor/gram-work-bot/pkg/telegram"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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

	telegramBot := telegram.NewBot(bot, cfg.Messages, tokenRepository)

	authorizationServer := server.NewAuthorizationServer(tokenRepository, "https://t.me/gram_work_bot")

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

		_, err2 := tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return err2
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}

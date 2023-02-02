package telegram

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	messages config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, messages config.Messages) *Bot {
	return &Bot{bot: bot, messages: messages}
}

func (b *Bot) Start() error {
	updates, err := b.initUpdatesChannel()

	if err != nil {
		return err
	}

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		} else {
			b.handleCommandFromKeybord(update.Message)
		}
	}
}

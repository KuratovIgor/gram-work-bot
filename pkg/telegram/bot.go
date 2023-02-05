package telegram

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/api"
	"github.com/KuratovIgor/gram-work-bot/pkg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	messages config.Messages
	mode     string
}

func NewBot(bot *tgbotapi.BotAPI, messages config.Messages) *Bot {
	return &Bot{bot: bot, messages: messages, mode: ""}
}

func (b *Bot) Start() error {
	updates, err := b.initUpdatesChannel()

	if err != nil {
		return err
	}

	api.GetAllAreas()
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
			b.handleBaseKeyboard(update.Message)
			b.handleVacanciesKeyboard(update.Message)
			b.handleFiltersKeyboard(update.Message)
			b.handleScheduleKeyboard(update.Message)
			b.handleExperienceKeyboard(update.Message)

			if !Contains(baseCommands, update.Message.Text) &&
				!Contains(vacanciesCommands, update.Message.Text) &&
				!Contains(filterCommands, update.Message.Text) &&
				b.mode != "" {
				b.handleMessage(update.Message)
			}
		}
	}
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

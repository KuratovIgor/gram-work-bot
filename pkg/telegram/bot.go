package telegram

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/api"
	"github.com/KuratovIgor/gram-work-bot/pkg/config"
	"github.com/KuratovIgor/gram-work-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot             *tgbotapi.BotAPI
	messages        config.Messages
	mode            string
	tokenRepository repository.TokenRepository
}

func NewBot(bot *tgbotapi.BotAPI, messages config.Messages, tr repository.TokenRepository) *Bot {
	return &Bot{bot: bot, messages: messages, mode: "", tokenRepository: tr}
}

func (b *Bot) Start(config *config.Config) error {
	updates, err := b.initUpdatesChannel()

	if err != nil {
		return err
	}

	api.GetAllAreas()
	b.handleUpdates(updates, config)

	return nil
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel, cfg *config.Config) {
	for update := range updates {
		_, err := b.getAccessToken(update.Message.Chat.ID)

		if err != nil {
			b.initAuthorizationProcess(cfg, update.Message)
			continue
		}

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message, cfg)
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

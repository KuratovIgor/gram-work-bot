package telegram

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/config"
	"github.com/KuratovIgor/gram-work-bot/pkg/repository"
	headhunter "github.com/KuratovIgor/head_hunter_sdk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot             *tgbotapi.BotAPI
	client          *headhunter.Client
	messages        config.Messages
	mode            string
	chosenVacancyId string
	applyMessage    string
	tokenRepository repository.TokenRepository
}

func NewBot(bot *tgbotapi.BotAPI, client *headhunter.Client, messages config.Messages, tr repository.TokenRepository) *Bot {
	return &Bot{
		bot:             bot,
		client:          client,
		messages:        messages,
		mode:            "",
		chosenVacancyId: "",
		applyMessage:    "",
		tokenRepository: tr,
	}
}

var AllAreas []headhunter.AreaType

func (b *Bot) Start() error {
	updates, err := b.initUpdatesChannel()

	if err != nil {
		return err
	}

	AllAreas, _ = b.client.GetAllAreas()

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
		if update.CallbackQuery != nil {
			b.handleInlineCommand(update)
			continue
		}

		if update.Message == nil {
			continue
		}

		_, err := b.getAccessToken(update.Message.Chat.ID)
		if err != nil {
			b.initAuthorizationProcess(update.Message)
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
			b.handleCancelMessageKeyboard(update.Message)

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

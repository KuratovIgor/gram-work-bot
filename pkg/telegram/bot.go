package telegram

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/config"
	"github.com/KuratovIgor/gram-work-bot/pkg/repository"
	headhunter "github.com/KuratovIgor/head_hunter_sdk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type (
	Bot struct {
		bot               *tgbotapi.BotAPI
		client            *headhunter.Client
		messages          config.Messages
		mode              string
		chosenVacancyId   string
		applyMessage      string
		graphqlRepository repository.GraphqlRepository
	}
)

func NewBot(bot *tgbotapi.BotAPI, client *headhunter.Client, messages config.Messages, graphqlRepository repository.GraphqlRepository) *Bot {
	return &Bot{
		bot:               bot,
		client:            client,
		messages:          messages,
		mode:              "",
		chosenVacancyId:   "",
		applyMessage:      "",
		graphqlRepository: graphqlRepository,
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

		//_, err := b.getAccessToken(update.Message.Chat.ID)
		_, err := b.getAccessToken(update.Message.Chat.ID)
		if err != nil {
			b.initAuthorizationProcess(update.Message)
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		} else {
			b.handleKeyboards(update.Message)
		}
	}
}

func (b *Bot) handleKeyboards(message *tgbotapi.Message) {
	b.handleBaseKeyboard(message)
	b.handleVacanciesKeyboard(message)
	b.handleFiltersKeyboard(message)
	b.handleScheduleKeyboard(message)
	b.handleExperienceKeyboard(message)
	b.handleCancelMessageKeyboard(message)

	if !Contains(baseCommands, message.Text) &&
		!Contains(vacanciesCommands, message.Text) &&
		!Contains(filterCommands, message.Text) &&
		b.mode != "" {
		b.handleMessage(message)
	}
}

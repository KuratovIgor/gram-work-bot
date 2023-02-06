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

var isLogin = false

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if isLogin == false {
			cfg, _ := config.Init()
			b.Login(cfg, update.Message)
		} else {
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
}

func (b *Bot) Login(config *config.Config, message *tgbotapi.Message) {
	fullAuthURI := b.generateAuthorizationLink(config)

	var button = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Авторизироваться", fullAuthURI),
		),
	)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Привет!\nДля начала поиска тебе нужно авторизироваться.\nПожалуйста, перейди по ссылке, нажав на кнопку ниже.")
	msg.ReplyMarkup = button

	b.bot.Send(msg)
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

package telegram

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/config"
	"github.com/KuratovIgor/gram-work-bot/pkg/repository"
	headhunter "github.com/KuratovIgor/head_hunter_sdk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

type (
	Bot struct {
		bot               *tgbotapi.BotAPI
		config            *config.Config
		client            *headhunter.Client
		messages          config.Messages
		mode              string
		chosenVacancyId   string
		applyMessage      string
		graphqlRepository repository.GraphqlRepository
		lkUrl             string
		users             map[int64]*headhunter.Client
	}
)

func NewBot(bot *tgbotapi.BotAPI, cfg *config.Config, client *headhunter.Client, messages config.Messages, graphqlRepository repository.GraphqlRepository, lkUrl string) *Bot {
	return &Bot{
		bot:               bot,
		config:            cfg,
		client:            client,
		messages:          messages,
		mode:              "",
		chosenVacancyId:   "",
		applyMessage:      "",
		graphqlRepository: graphqlRepository,
		lkUrl:             lkUrl,
		users:             map[int64]*headhunter.Client{},
	}
}

var AllAreas []headhunter.AreaType

func (b *Bot) Start() error {
	updates := b.initUpdatesChannel()

	AllAreas, _ = b.client.GetAllAreas()

	chatIds, error := b.getSessions()
	if error != nil {
		return error
	}

	for _, chatId := range chatIds {
		headhunterClient, clientErr := headhunter.NewClient(b.config.ClientID, b.config.ClientSecret, b.config.RedirectURI)
		if clientErr != nil {
			return clientErr
		}

		b.users[chatId] = headhunterClient
	}

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	//quit := make(chan bool)

	conf := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     "start",
			Description: "Начать работу",
		},
		tgbotapi.BotCommand{
			Command:     "help",
			Description: "Помощь",
		},
		tgbotapi.BotCommand{
			Command:     "lk",
			Description: "Войти в личный кабинет",
		},
		tgbotapi.BotCommand{
			Command:     "logout",
			Description: "Завершить поиск",
		},
	)

	_, error := b.bot.Request(conf)
	if error != nil {
		log.Panic(error)
	}

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
			headhunterClient, err := headhunter.NewClient(b.config.ClientID, b.config.ClientSecret, b.config.RedirectURI)
			if err != nil {
				log.Fatal(err)
			}

			b.users[update.Message.Chat.ID] = headhunterClient

			authErr := b.initAuthorizationProcess(update.Message)
			if authErr != nil {
				delete(b.users, update.Message.Chat.ID)
				log.Fatal(err)
			}

			go func() {
				upt := update
				for {
					err := b.runResponsesUpdates(upt)
					if err != nil {
						return
					}
				}
			}()

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
	err := b.handleBaseKeyboard(message)
	if err != nil {
		b.handleRelogin(message)
	}

	err = b.handleVacanciesKeyboard(message)
	if err != nil {
		b.handleRelogin(message)
	}

	err = b.handleFiltersKeyboard(message)
	if err != nil {
		b.handleRelogin(message)
	}

	err = b.handleScheduleKeyboard(message)
	if err != nil {
		b.handleRelogin(message)
	}

	err = b.handleExperienceKeyboard(message)
	if err != nil {
		b.handleRelogin(message)
	}

	err = b.handleCancelMessageKeyboard(message)
	if err != nil {
		b.handleRelogin(message)
	}

	if !Contains(baseCommands, message.Text) &&
		!Contains(vacanciesCommands, message.Text) &&
		!Contains(filterCommands, message.Text) &&
		b.mode != "" {
		b.handleMessage(message)
	}
}

func (b *Bot) runResponsesUpdates(update tgbotapi.Update) error {
	responses, err := b.getResponses(update.Message.Chat.ID)
	if err != nil && err.Error() != "unauthorized user" {
		return err
	}

	for _, response := range responses {
		me, err := b.getInfoAboutMe(update.Message.Chat.ID)
		if err != nil {
			return err
		}

		err = b.graphqlRepository.UpdateResponseStatus(response.Vacancy.Id, me.UserID, response.State)
		if err != nil {
			return err
		}
	}

	time.Sleep(time.Hour)

	return nil
}

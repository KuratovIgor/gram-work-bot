package telegram

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/config"
	"github.com/KuratovIgor/gram-work-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

const authURI = "https://hh.ru/oauth/authorize"

func (b *Bot) initAuthorizationProcess(config *config.Config, message *tgbotapi.Message) error {
	fullAuthURI := b.generateAuthorizationLink(config, message.Chat.ID)

	var button = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Авторизироваться", fullAuthURI),
		),
	)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Привет!\nДля начала поиска тебе нужно авторизироваться.\nПожалуйста, перейди по ссылке, нажав на кнопку ниже.\n\nПосле авторизации введи команду /start для начала работы.")
	msg.ReplyMarkup = button

	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) getAccessToken(chatID int64) (string, error) {
	return b.tokenRepository.Get(chatID, repository.AccessTokens)
}

func (b *Bot) generateAuthorizationLink(config *config.Config, chatID int64) string {
	redirectURL := getRedirectURL(config, chatID)

	return authURI + "?response_type=code&client_id=" + config.ClientID + "&redirect_uri=" + redirectURL
}

func getRedirectURL(config *config.Config, chatID int64) string {
	return config.RedirectURI + "/?chat_id=" + strconv.Itoa(int(chatID))
}

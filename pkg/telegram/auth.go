package telegram

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const authURI = "https://hh.ru/oauth/authorize"

func (b *Bot) initAuthorizationProcess(message *tgbotapi.Message) error {
	authorizeLink, _ := b.client.GetAuthorizationURL(message.Chat.ID)

	var button = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Авторизироваться", authorizeLink),
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

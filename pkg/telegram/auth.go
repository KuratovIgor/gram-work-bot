package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (b *Bot) initAuthorizationProcess(message *tgbotapi.Message) error {
	authorizeLink, _ := b.client.GetAuthorizationURL(message.Chat.ID)

	return b.displayAuthorizeMessage(authorizeLink, message)
}

func (b *Bot) handleLogout(message *tgbotapi.Message) error {
	resErr := b.logout(message.Chat.ID)
	if resErr != nil {
		return resErr
	}

	removeErr := b.graphqlRepository.RemoveSession(message.Chat.ID)
	if removeErr != nil {
		log.Panic(removeErr)
		return removeErr
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Работа завершена!\nВозвращайся скорее!")

	_, sendErr := b.bot.Send(msg)
	if sendErr != nil {
		return sendErr
	}

	return nil
}

func (b *Bot) getAccessToken(chatID int64) (string, error) {
	return b.graphqlRepository.GetAccessToken(chatID)
}

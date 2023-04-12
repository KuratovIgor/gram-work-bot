package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) initAuthorizationProcess(message *tgbotapi.Message) error {
	authorizeLink, _ := b.users[message.Chat.ID].GetAuthorizationURL(message.Chat.ID)

	return b.displayAuthorizeMessage(authorizeLink, message)
}

func (b *Bot) handleLogout(message *tgbotapi.Message) error {
	resErr := b.logout(message.Chat.ID)
	if resErr != nil {
		return resErr
	}

	removeErr := b.graphqlRepository.RemoveSession(message.Chat.ID)
	if removeErr != nil {
		return removeErr
	}

	delete(b.users, message.Chat.ID)

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

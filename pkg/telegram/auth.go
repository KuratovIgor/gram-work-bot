package telegram

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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

	rmAccessTokenErr := b.tokenRepository.Delete(message.Chat.ID, repository.AccessTokens)
	if rmAccessTokenErr != nil {
		return rmAccessTokenErr
	}

	rmRefreshTokenErr := b.tokenRepository.Delete(message.Chat.ID, repository.RefreshToken)
	if rmRefreshTokenErr != nil {
		return rmRefreshTokenErr
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Работа завершена!\nВозвращайся скорее!")

	_, sendErr := b.bot.Send(msg)
	if sendErr != nil {
		return sendErr
	}

	return nil
}

func (b *Bot) getAccessToken(chatID int64) (string, error) {
	return b.tokenRepository.Get(chatID, repository.AccessTokens)
}

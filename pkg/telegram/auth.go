package telegram

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) initAuthorizationProcess(message *tgbotapi.Message) error {
	authorizeLink, _ := b.client.GetAuthorizationURL(message.Chat.ID)

	return b.displayAuthorizeMessage(authorizeLink, message)

}

func (b *Bot) getAccessToken(chatID int64) (string, error) {
	return b.tokenRepository.Get(chatID, repository.AccessTokens)
}

package telegram

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var params = api.NewParams()

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownCommand)

	switch message.Command() {
	case commandStart:
		msg.Text = ""
		params.ClearParams()
		b.openBaseKeyboard(message)
		_, error := b.bot.Send(msg)
		return error
	default:
		_, error := b.bot.Send(msg)
		return error
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	switch b.mode {
	case "search":
		params.SetSearch(msg.Text)
		error := b.handleGetVacancies(message)
		return error
	case "salary":
		params.SetSalary(msg.Text)
		error := b.handleGetVacancies(message)
		return error
	}

	return nil
}

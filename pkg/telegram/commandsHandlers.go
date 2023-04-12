package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownCommand)

	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		_, error := b.bot.Send(msg)
		return error
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	b.users[message.Chat.ID].UrlParams.ClearParams()

	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthorizationProcess(message)
	}

	b.openBaseKeyboard(message)

	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	switch b.mode {
	case "search":
		return b.handleSearch(message)
	case "salary":
		return b.handleFilterBySalary(message)
	case "area":
		return b.handleFilterByArea(message)
	case "schedule":
		return b.handleFilterBySchedule(message)
	case "experience":
		return b.handleFilterByExperience(message)
	case "message":
		b.applyMessage = message.Text
		applyErr := b.handleApplyToJob(message)
		if applyErr != nil {
			return applyErr
		}
	}

	return nil
}

func (b *Bot) handleInlineCommand(update tgbotapi.Update) error {
	switch b.mode {
	case "chooseResume":
		applyErr := b.handleApplyToJobByResume(update.CallbackQuery.Message, update.CallbackQuery.Data)
		if applyErr != nil {
			return applyErr
		}
	case "apply":
		b.handleSendApplyMessage(update.CallbackQuery.Message, update.CallbackQuery.Data)
	default:
		b.handleSendApplyMessage(update.CallbackQuery.Message, update.CallbackQuery.Data)
	}

	return nil
}

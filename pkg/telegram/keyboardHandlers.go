package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *Bot) handleBaseKeyboard(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownCommand)

	switch message.Text {
	case baseCommands[0]:
		params.SetPage(0)
		error := b.handleGetVacancies(message)
		b.openVacanciesKeyboard(message)
		return error
	case baseCommands[1]:
		msg.Text = "Создаем резюме..."
		_, error := b.bot.Send(msg)
		return error
	}

	return nil
}

func (b *Bot) handleVacanciesKeyboard(message *tgbotapi.Message) error {
	switch message.Text {
	case vacanciesCommands[0]:
		params.SetPage(params.Page + 1)
		error := b.handleGetVacancies(message)
		return error
	case vacanciesCommands[1]:
		b.openFilterKeyboard(message)
	case vacanciesCommands[2]:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Введите должность для поиска")
		b.bot.Send(msg)
		b.mode = "search"
	case vacanciesCommands[3]:
		b.mode = ""
		params.ClearParams()
		b.openBaseKeyboard(message)
	}

	return nil
}

func (b *Bot) handleFiltersKeyboard(message *tgbotapi.Message) error {
	switch message.Text {
	case filterCommands[0]:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Введите сумму в рублях")
		b.bot.Send(msg)
		b.mode = "salary"
	case filterCommands[4]:
		params.ClearFilters()
		b.mode = ""
		error := b.handleGetVacancies(message)
		return error
	case filterCommands[5]:
		b.mode = ""
		b.openVacanciesKeyboard(message)
	}

	return nil
}

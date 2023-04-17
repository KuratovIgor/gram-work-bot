package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart  = "start"
	commandHelp   = "help"
	commandLK     = "lk"
	commandLogout = "logout"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownCommand)

	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandHelp:
		return b.handleHelpCommand(message)
	case commandLK:
		return b.displayLKUrl(message)
	case commandLogout:
		return b.handleLogout(message)
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

func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	text := "*ВАКАНСИИ:*\n\n- *Поиск вакансий* — список из 5 вакансий.\n\n- *Больше вакансий* — список из следующих 5 вакансий.\n\n- *Поиск* — поиск вакансий по должности.\n\n\n*ФИЛЬТРЫ:*\n\n- *З/П* — фильтрация по зароботной плате.\n\n- *Город* — фильтрация по местоположению.\n\n- *График* — фильтрация по графику работы.\n\n- *Опыт* — фильтрация по опыту работы.\n\n- *По умолчанию* — фильтры будут установлены в соответствии с фильтрами из личного кабинета.\n\n- *Сбросить* — сброс всех фильтров и параметров поиска.\n\n\n*ПОЛЬЗОВАТЕЛЬСКИЕ:*\n\n- *Мои отклики* — список из последних 5 откликов.\n\n- *Мои резюме* — список ваших резюме.\n\n- *Личный кабинет* — вход в личный кабинет, бот пришлет ссылку и код доступа.\n\n\n*Завершить поиск* — выход из системы."

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ParseMode = "markdown"

	// "ВАКАНСИИ:\n\n<Поиск вакансий - бот пришлет список из 5 вакансий.\n\nБольше вакансий - бот пришлет ещё 5 вакансий.\n\nПоиск - после ввода должности бот пришлет найденные вакансии."

	if _, err := b.bot.Send(msg); err != nil {
		return err
	}

	return nil
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

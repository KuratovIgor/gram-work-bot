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
	b.client.UrlParams.ClearParams()

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
	}

	return nil
}

func (b *Bot) handleSearch(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	b.client.UrlParams.SetSearch(msg.Text)
	vacancies, _ := b.client.GetVacancies()

	return b.displayVacancies(vacancies, message)
}

func (b *Bot) handleFilterBySalary(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	b.client.UrlParams.SetSalary(msg.Text)
	vacancies, _ := b.client.GetVacancies()

	return b.displayVacancies(vacancies, message)
}

func (b *Bot) handleFilterByArea(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	areaId := searchAreaByName(msg.Text)
	b.client.UrlParams.SetArea(areaId)
	vacancies, _ := b.client.GetVacancies()

	return b.displayVacancies(vacancies, message)
}

func (b *Bot) handleFilterBySchedule(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	scheduleId := getScheduleIdByText(message.Text)

	if scheduleId != "unknown" {
		b.client.UrlParams.SetSchedule(scheduleId)
		vacancies, _ := b.client.GetVacancies()

		return b.displayVacancies(vacancies, message)
	}

	msg.Text = "Ты ввел неизвестный график :("
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleFilterByExperience(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	experienceId := getExperienceIdByText(message.Text)

	if experienceId != "unknown" {
		b.client.UrlParams.SetExperience(experienceId)
		vacancies, _ := b.client.GetVacancies()

		return b.displayVacancies(vacancies, message)
	}

	msg.Text = "Ты ввел неизвестный опыт работы :("
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleInlineCommand(update tgbotapi.Update) error {
	switch b.mode {
	case "chooseResume":
		return b.handleApplyToJobByResume(update.CallbackQuery.Message, update.CallbackQuery.Data)
	case "apply":
		return b.handleApplyToJob(update.CallbackQuery.Message, update.CallbackQuery.Data)
	default:
		return b.handleApplyToJob(update.CallbackQuery.Message, update.CallbackQuery.Data)
	}

	return nil
}

func (b *Bot) handleApplyToJob(message *tgbotapi.Message, vacancyId string) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Ваш отклик успешно отправлен!")

	b.chosenResumeId = vacancyId
	resumeIds, _ := b.client.GetResumesIds()

	// Если пользователь выбрал сначала резюме, а не вакансию, то вакансии с id резюме не существует
	if Contains(resumeIds, vacancyId) {
		return nil
	}

	if len(resumeIds) > 1 {
		return b.displayChoosingResume(message)
	}

	err := b.client.ApplyToJob(vacancyId, b.chosenResumeId, "")
	if err != nil {
		msg.Text = "Вы уже откликнулись на эту вакансию"
	}

	b.chosenResumeId = ""
	b.mode = ""

	_, sendErr := b.bot.Send(msg)
	if sendErr != nil {
		return sendErr
	}

	return nil
}

func (b *Bot) handleApplyToJobByResume(message *tgbotapi.Message, resumeId string) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Ваш отклик успешно отправлен!")

	err := b.client.ApplyToJob(b.chosenResumeId, resumeId, "")
	if err != nil {
		msg.Text = "Вы уже откликнулись на эту вакансию"
	}

	b.chosenResumeId = ""
	b.mode = ""

	_, sendErr := b.bot.Send(msg)
	if sendErr != nil {
		return sendErr
	}

	return nil
}

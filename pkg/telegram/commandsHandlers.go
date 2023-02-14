package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
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
	case "message":
		b.applyMessage = message.Text
		return b.handleApplyToJob(message)
	}

	return nil
}

func (b *Bot) handleSearch(message *tgbotapi.Message) error {
	token, tokenErr := b.getAccessToken(message.Chat.ID)
	if tokenErr != nil {
		return tokenErr
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	b.client.UrlParams.SetSearch(msg.Text)
	vacancies, _ := b.client.GetVacancies(token)

	return b.displayVacancies(vacancies, message)
}

func (b *Bot) handleFilterBySalary(message *tgbotapi.Message) error {
	token, tokenErr := b.getAccessToken(message.Chat.ID)
	if tokenErr != nil {
		return tokenErr
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	b.client.UrlParams.SetSalary(msg.Text)
	vacancies, _ := b.client.GetVacancies(token)

	return b.displayVacancies(vacancies, message)
}

func (b *Bot) handleFilterByArea(message *tgbotapi.Message) error {
	token, tokenErr := b.getAccessToken(message.Chat.ID)
	if tokenErr != nil {
		return tokenErr
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	areaId := searchAreaByName(msg.Text)
	b.client.UrlParams.SetArea(areaId)
	vacancies, _ := b.client.GetVacancies(token)

	return b.displayVacancies(vacancies, message)
}

func (b *Bot) handleFilterBySchedule(message *tgbotapi.Message) error {
	token, tokenErr := b.getAccessToken(message.Chat.ID)
	if tokenErr != nil {
		return tokenErr
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	scheduleId := getScheduleIdByText(message.Text)

	if scheduleId != "unknown" {
		b.client.UrlParams.SetSchedule(scheduleId)
		vacancies, _ := b.client.GetVacancies(token)

		return b.displayVacancies(vacancies, message)
	}

	msg.Text = "Ты ввел неизвестный график :("
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleFilterByExperience(message *tgbotapi.Message) error {
	token, tokenErr := b.getAccessToken(message.Chat.ID)
	if tokenErr != nil {
		return tokenErr
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	experienceId := getExperienceIdByText(message.Text)

	if experienceId != "unknown" {
		b.client.UrlParams.SetExperience(experienceId)
		vacancies, _ := b.client.GetVacancies(token)

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
		b.handleSendApplyMessage(update.CallbackQuery.Message, update.CallbackQuery.Data)
	default:
		b.handleSendApplyMessage(update.CallbackQuery.Message, update.CallbackQuery.Data)
	}

	return nil
}

func (b *Bot) handleSendApplyMessage(message *tgbotapi.Message, vacancyId string) {
	b.mode = "message"
	b.chosenVacancyId = vacancyId
	b.openCancelMessageKeyboard(message)
}

func (b *Bot) handleApplyToJob(message *tgbotapi.Message) error {
	token, tokenErr := b.getAccessToken(message.Chat.ID)
	if tokenErr != nil {
		return tokenErr
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Отклик успешно отправлен")

	resumeIds, _ := b.client.GetResumesIds(token)

	// Если пользователь выбрал сначала резюме, а не вакансию, то вакансии с id резюме не существует
	if Contains(resumeIds, b.chosenVacancyId) {
		return nil
	}

	if len(resumeIds) > 1 {
		return b.displayChoosingResume(message)
	}

	err := b.client.ApplyToJob(b.chosenVacancyId, resumeIds[0], b.applyMessage, token)
	if err != nil {
		msg.Text = "Вы уже откликнулись на эту вакансию"
	}

	b.chosenVacancyId = ""
	b.mode = ""

	_, sendErr := b.bot.Send(msg)
	if sendErr != nil {
		return sendErr
	}

	b.openVacanciesKeyboard(message)

	return nil
}

func (b *Bot) handleApplyToJobByResume(message *tgbotapi.Message, resumeId string) error {
	token, tokenErr := b.getAccessToken(message.Chat.ID)
	if tokenErr != nil {
		return tokenErr
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Отклик успешно отправлен")

	err := b.client.ApplyToJob(b.chosenVacancyId, resumeId, b.applyMessage, token)
	log.Println(123123123)
	if err != nil {
		msg.Text = "Вы уже откликнулись на эту вакансию"
	}

	b.chosenVacancyId = ""
	b.mode = ""

	_, sendErr := b.bot.Send(msg)
	if sendErr != nil {
		return sendErr
	}

	b.openVacanciesKeyboard(message)

	return nil
}

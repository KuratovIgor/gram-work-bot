package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleSendApplyMessage(message *tgbotapi.Message, vacancyId string) {
	b.mode = "message"
	b.chosenVacancyId = vacancyId
	b.openCancelMessageKeyboard(message)
}

func (b *Bot) handleApplyToJob(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Отклик успешно отправлен")

	resumeIds, resErr := b.getResumesIds(message.Chat.ID)
	if resErr != nil {
		return resErr
	}

	// Если пользователь выбрал сначала резюме, а не вакансию, то вакансии с id резюме не существует
	if Contains(resumeIds, b.chosenVacancyId) {
		return nil
	}

	if len(resumeIds) > 1 {
		return b.displayChoosingResume(message)
	}

	err := b.applyToJob(message.Chat.ID, b.chosenVacancyId, resumeIds[0], b.applyMessage)
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
	msg := tgbotapi.NewMessage(message.Chat.ID, "Отклик успешно отправлен")

	err := b.applyToJob(message.Chat.ID, b.chosenVacancyId, resumeId, b.applyMessage)
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

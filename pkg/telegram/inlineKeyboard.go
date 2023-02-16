package telegram

import (
	headhunter "github.com/KuratovIgor/head_hunter_sdk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) getAuthorizeButton(authorizeLink string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Авторизироваться", authorizeLink),
		),
	)
}

func (b *Bot) getVacancyKeyboard(vacancy headhunter.Vacancy) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Откликнуться", vacancy.Id),
			tgbotapi.NewInlineKeyboardButtonURL("Перейти", vacancy.AlternateUrl),
		),
	)
}

func (b *Bot) getOpeningResumeButton(resume headhunter.Resume) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Открыть", resume.URL),
		),
	)
}

func (b *Bot) getChoosingResumeKeyboard(resumes []headhunter.Resume) []tgbotapi.InlineKeyboardButton {
	var resumeKeyboard []tgbotapi.InlineKeyboardButton

	for _, resume := range resumes {
		resumeKeyboard = append(resumeKeyboard, tgbotapi.NewInlineKeyboardButtonData(resume.Title, resume.Id))
	}

	return resumeKeyboard
}

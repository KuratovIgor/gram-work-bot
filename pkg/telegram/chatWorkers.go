package telegram

import (
	"fmt"
	headhunter "github.com/KuratovIgor/head_hunter_sdk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

const (
	vacancyMessage = "ДОЛЖНОСТЬ:\n%s\n\nЗАРПЛАТА:\n от %s до %s %s\n\nГород:\n%s\n\nРАБОТОДАТЕЛЬ:\n%s\n\nОПИСАНИЕ:\n%s\n\nТРЕБОВАНИЯ:\n%s\n\nГРАФИК:\n%s\n\nОПУБЛИКОВАНО:\n%s"
	resumeMessage  = "%s, %s\n\nДОЛЖНОСТЬ\n%s    %s\n\nГОРОД\n%s\n\nОБРАЗОВНАИЕ\n%s"
)

func (b *Bot) displayAuthorizeMessage(authorizeLink string, message *tgbotapi.Message) error {
	var button = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Авторизироваться", authorizeLink),
		),
	)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Привет!\nДля начала поиска тебе нужно авторизироваться.\nПожалуйста, перейди по ссылке, нажав на кнопку ниже.\n\nПосле авторизации введи команду /start для начала работы.")
	msg.ReplyMarkup = button

	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) displayVacancies(vacancies []headhunter.Vacancy, message *tgbotapi.Message) error {
	for _, item := range vacancies {
		var buttons = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Откликнуться", item.Id),
				tgbotapi.NewInlineKeyboardButtonURL("Перейти", item.AlternateUrl),
			),
		)

		time, _ := time.Parse("2006-01-02T15:04:05-0700", item.PublishedAt)
		publishedDate := time.Format("02 January 2006")

		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(vacancyMessage, item.Name, item.Salary.From, item.Salary.To,
			item.Salary.Currency, item.Area, item.Employer, item.Responsibility, item.Requirement, item.Schedule, publishedDate))

		msg.ReplyMarkup = buttons

		b.mode = "apply"

		_, err := b.bot.Send(msg)
		if err != nil {
			return nil
		}
	}

	return nil
}

func (b *Bot) displayMyResumes(message *tgbotapi.Message) error {
	token, tokenErr := b.getAccessToken(message.Chat.ID)
	if tokenErr != nil {
		return tokenErr
	}

	resumes, err := b.client.GetResumes(token)
	if err != nil {
		return err
	}

	for _, resume := range resumes {
		var buttons = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("Открыть", resume.URL),
			),
		)

		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(resumeMessage, resume.Name, resume.Age, resume.Title, resume.Salary, resume.Area, resume.Education))
		msg.ReplyMarkup = buttons

		_, sendErr := b.bot.Send(msg)
		if sendErr != nil {
			return sendErr
		}
	}

	return nil
}

func (b *Bot) displayChoosingResume(message *tgbotapi.Message) error {
	token, tokenErr := b.getAccessToken(message.Chat.ID)
	if tokenErr != nil {
		return tokenErr
	}

	resumes, err := b.client.GetResumes(token)
	if err != nil {
		return err
	}

	var buttons []tgbotapi.InlineKeyboardButton
	for _, resume := range resumes {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(resume.Title, resume.Id))
	}

	var keyboard = tgbotapi.NewInlineKeyboardMarkup(buttons)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Выбери резюме, которое желаешь отправить")
	msg.ReplyMarkup = keyboard

	b.mode = "chooseResume"

	_, sendErr := b.bot.Send(msg)
	if sendErr != nil {
		return nil
	}

	return nil
}

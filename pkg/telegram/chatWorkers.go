package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

const (
	vacancyMessage = "ДОЛЖНОСТЬ:\n%s\n\nЗАРПЛАТА:\n от %s до %s %s\n\nГород:\n%s\n\nРАБОТОДАТЕЛЬ:\n%s\n\nОПИСАНИЕ:\n%s\n\nТРЕБОВАНИЯ:\n%s\n\nГРАФИК:\n%s\n\nОПУБЛИКОВАНО:\n%s"
	resumeMessage  = "%s, %s\n\nДОЛЖНОСТЬ\n%s    %s\n\nГОРОД\n%s\n\nОБРАЗОВНАИЕ\n%s"
)

func (b *Bot) displayAuthorizeMessage(authorizeLink string, message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Привет!\nДля начала поиска тебе нужно авторизироваться.\nПожалуйста, перейди по ссылке, нажав на кнопку ниже.\n\nПосле авторизации введи команду /start для начала работы.")

	authButton := b.getAuthorizeButton(authorizeLink)
	msg.ReplyMarkup = authButton

	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) displayVacancies(message *tgbotapi.Message) error {
	vacancies, err := b.getVacancies(message.Chat.ID)
	if err != nil {
		return err
	}

	for _, vacancy := range vacancies {
		time, _ := time.Parse("2006-01-02T15:04:05-0700", vacancy.PublishedAt)
		publishedDate := time.Format("02 January 2006")

		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(vacancyMessage, vacancy.Name, vacancy.Salary.From, vacancy.Salary.To,
			vacancy.Salary.Currency, vacancy.Area, vacancy.Employer, vacancy.Responsibility, vacancy.Requirement, vacancy.Schedule, publishedDate))

		vacancyKeyboard := b.getVacancyKeyboard(vacancy)
		msg.ReplyMarkup = vacancyKeyboard

		_, err := b.bot.Send(msg)
		if err != nil {
			return nil
		}
	}

	b.mode = "apply"

	return nil
}

func (b *Bot) displayMyResumes(message *tgbotapi.Message) error {
	resumes, err := b.getResumes(message.Chat.ID)
	if err != nil {
		return err
	}

	for _, resume := range resumes {
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(resumeMessage, resume.Name, resume.Age, resume.Title, resume.Salary, resume.Area, resume.Education))

		resumeButton := b.getOpeningResumeButton(resume)
		msg.ReplyMarkup = resumeButton

		_, sendErr := b.bot.Send(msg)
		if sendErr != nil {
			return sendErr
		}
	}

	return nil
}

func (b *Bot) displayChoosingResume(message *tgbotapi.Message) error {
	resumes, err := b.getResumes(message.Chat.ID)
	if err != nil {
		return err
	}

	choosingResumeKeyboard := b.getChoosingResumeKeyboard(resumes)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(choosingResumeKeyboard)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Выбери резюме, которое желаешь отправить")
	msg.ReplyMarkup = keyboard

	b.mode = "chooseResume"

	_, sendErr := b.bot.Send(msg)
	if sendErr != nil {
		return nil
	}

	return nil
}

package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

const (
	resumeMessage = "%s, %s\n\nДОЛЖНОСТЬ\n%s    %s\n\nГОРОД\n%s\n\nОБРАЗОВНАИЕ\n%s"
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
		messageTemplate := getVacancyMessageTemplate(vacancy)

		msg := tgbotapi.NewMessage(message.Chat.ID, messageTemplate)

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

func (b *Bot) displayMyResponses(message *tgbotapi.Message) error {
	responses, err := b.getResponses(message.Chat.ID)
	if err != nil {
		return err
	}

	for _, response := range responses {
		messageTemplate := getResponseMessageTemplate(response)

		msg := tgbotapi.NewMessage(message.Chat.ID, messageTemplate)

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

func (b *Bot) displayLKUrl(message *tgbotapi.Message) error {
	lkUrlButton := b.getLKUrlButton()

	msg := tgbotapi.NewMessage(message.Chat.ID, "Для входа в личный кабинет тебе необходимо перейти по ссылке ниже.\nПри авторизации необходимо использовать данный код:\n"+strconv.Itoa(int(message.Chat.ID)))
	msg.ReplyMarkup = lkUrlButton

	_, sendErr := b.bot.Send(msg)
	if sendErr != nil {
		return sendErr
	}

	return nil
}

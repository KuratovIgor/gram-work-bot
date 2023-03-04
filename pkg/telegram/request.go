package telegram

import (
	headhunter "github.com/KuratovIgor/head_hunter_sdk"
)

func (b *Bot) getInfoAboutMe(chatId int64) (*headhunter.MeType, error) {
	token, tokenErr := b.getAccessToken(chatId)
	if tokenErr != nil {
		return nil, tokenErr
	}

	info, resErr := b.client.GetInfoAboutMe(token)
	if resErr != nil {
		return nil, resErr
	}

	return info, resErr
}

func (b *Bot) getVacancies(chatId int64) ([]headhunter.Vacancy, error) {
	token, tokenErr := b.getAccessToken(chatId)
	if tokenErr != nil {
		return nil, tokenErr
	}

	vacancies, resErr := b.client.GetVacancies(token)
	if resErr != nil {
		return nil, resErr
	}

	return vacancies, nil
}

func (b *Bot) getVacancy(chatId int64, vacancyId string) (*headhunter.Vacancy, error) {
	token, tokenErr := b.getAccessToken(chatId)
	if tokenErr != nil {
		return nil, tokenErr
	}

	vacancy, resErr := b.client.GetVacancy(vacancyId, token)
	if resErr != nil {
		return nil, resErr
	}

	return vacancy, nil
}

func (b *Bot) getResumes(chatId int64) ([]headhunter.Resume, error) {
	token, tokenErr := b.getAccessToken(chatId)
	if tokenErr != nil {
		return nil, tokenErr
	}

	resumes, err := b.client.GetResumes(token)
	if err != nil {
		return nil, err
	}

	return resumes, nil
}

func (b *Bot) getResponses(chatId int64) ([]headhunter.Response, error) {
	token, tokenErr := b.getAccessToken(chatId)
	if tokenErr != nil {
		return nil, tokenErr
	}

	responses, err := b.client.GetResponseList(token)
	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (b *Bot) getResumesIds(chatId int64) ([]string, error) {
	token, tokenErr := b.getAccessToken(chatId)
	if tokenErr != nil {
		return nil, tokenErr
	}

	resumesIds, err := b.client.GetResumesIds(token)
	if err != nil {
		return nil, err
	}

	return resumesIds, nil
}

func (b *Bot) applyToJob(chatId int64, vacancyId string, resumeId string, message string) error {
	token, tokenErr := b.getAccessToken(chatId)
	if tokenErr != nil {
		return tokenErr
	}

	return b.client.ApplyToJob(vacancyId, resumeId, message, token)
}

func (b *Bot) logout(chatId int64) error {
	token, tokenErr := b.getAccessToken(chatId)
	if tokenErr != nil {
		return tokenErr
	}

	return b.client.Logout(token)
}

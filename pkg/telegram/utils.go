package telegram

import (
	"fmt"
	headhunter "github.com/KuratovIgor/head_hunter_sdk"
	"strings"
	"time"
)

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func searchAreaByName(name string) string {
	for _, area := range AllAreas {
		if strings.Contains(area.Name, name) {
			return area.Id
		}
	}

	return "113"
}

func getScheduleIdByText(schedule string) string {
	switch schedule {
	case scheduleCommands[0]:
		return "fullDay"
	case scheduleCommands[1]:
		return "shift"
	case scheduleCommands[2]:
		return "flexible"
	case scheduleCommands[3]:
		return "remote"
	case scheduleCommands[4]:
		return "flyInFlyOut"
	}

	return "unknown"
}

func getExperienceIdByText(experience string) string {
	switch experience {
	case experienceCommands[0]:
		return "noExperience"
	case experienceCommands[1]:
		return "between1And3"
	case experienceCommands[2]:
		return "between3And6"
	case experienceCommands[3]:
		return "moreThan6"

	}

	return "unknown"
}

func getVacancyMessageTemplate(vacancy headhunter.Vacancy) string {
	var messageTemplate string
	var salaryString string
	var responsibilityString string

	if len(vacancy.Responsibility) > 0 {
		responsibilityString = "*ОПИСАНИЕ*:\n%s\n\n"
	}

	if len(vacancy.Salary.From) > 0 && len(vacancy.Salary.To) > 0 {
		salaryString = "*ЗАРПЛАТА*:\n от %s до %s %s\n\n"
	} else {
		salaryString = "*ЗАРПЛАТА*:\n %s %s\n\n"
	}

	time, _ := time.Parse("2006-01-02T15:04:05-0700", vacancy.PublishedAt)
	publishedDate := time.Format("02 January 2006")

	messageTemplate = "*ДОЛЖНОСТЬ*:\n%s\n\n" + salaryString + "*ГОРОД*:\n%s\n\n*РАБОТОДАТЕЛЬ*:\n%s\n\n" + responsibilityString + "*ТРЕБОВАНИЯ*:\n%s\n\n*ГРАФИК*:\n%s\n\n*ОПУБЛИКОВАНО*:\n%s"

	if len(vacancy.Salary.From) > 0 && len(vacancy.Salary.To) > 0 {
		if len(responsibilityString) > 0 {
			return fmt.Sprintf(messageTemplate, vacancy.Name, vacancy.Salary.From, vacancy.Salary.To,
				vacancy.Salary.Currency, vacancy.Area, vacancy.Employer, vacancy.Responsibility, vacancy.Requirement, vacancy.Schedule, publishedDate)
		} else {
			return fmt.Sprintf(messageTemplate, vacancy.Name, vacancy.Salary.From, vacancy.Salary.To,
				vacancy.Salary.Currency, vacancy.Area, vacancy.Employer, vacancy.Requirement, vacancy.Schedule, publishedDate)
		}
	} else if len(vacancy.Salary.From) > 0 {
		if len(responsibilityString) > 0 {
			return fmt.Sprintf(messageTemplate, vacancy.Name, vacancy.Salary.From,
				vacancy.Salary.Currency, vacancy.Area, vacancy.Employer, vacancy.Responsibility, vacancy.Requirement, vacancy.Schedule, publishedDate)
		} else {
			return fmt.Sprintf(messageTemplate, vacancy.Name, vacancy.Salary.From,
				vacancy.Salary.Currency, vacancy.Area, vacancy.Employer, vacancy.Requirement, vacancy.Schedule, publishedDate)
		}
	} else if len(vacancy.Salary.To) > 0 {
		if len(responsibilityString) > 0 {
			return fmt.Sprintf(messageTemplate, vacancy.Name, vacancy.Salary.To,
				vacancy.Salary.Currency, vacancy.Area, vacancy.Employer, vacancy.Responsibility, vacancy.Requirement, vacancy.Schedule, publishedDate)
		} else {
			return fmt.Sprintf(messageTemplate, vacancy.Name, vacancy.Salary.To,
				vacancy.Salary.Currency, vacancy.Area, vacancy.Employer, vacancy.Requirement, vacancy.Schedule, publishedDate)
		}
	} else {
		if len(responsibilityString) > 0 {
			return fmt.Sprintf(messageTemplate, vacancy.Name, "Не указана",
				vacancy.Salary.Currency, vacancy.Area, vacancy.Employer, vacancy.Responsibility, vacancy.Requirement, vacancy.Schedule, publishedDate)
		} else {
			return fmt.Sprintf(messageTemplate, vacancy.Name, "Не указана",
				vacancy.Salary.Currency, vacancy.Area, vacancy.Employer, vacancy.Requirement, vacancy.Schedule, publishedDate)
		}
	}
}

func getResponseMessageTemplate(response headhunter.Response) string {
	var messageTemplate string
	var salaryString string

	if len(response.Vacancy.Salary.From) > 0 && len(response.Vacancy.Salary.To) > 0 {
		salaryString = "*ЗАРПЛАТА*:\n от %s до %s %s\n\n"
	} else {
		salaryString = "*ЗАРПЛАТА*:\n %s %s\n\n"
	}

	messageTemplate = "*ДОЛЖНОСТЬ*:\n%s\n\n" + salaryString + "*ГОРОД*:\n%s\n\n*РАБОТОДАТЕЛЬ*:\n%s\n\n%s\n\n*СТАТУС*:    %s"

	if len(response.Vacancy.Salary.From) > 0 && len(response.Vacancy.Salary.To) > 0 {
		return fmt.Sprintf(messageTemplate, response.Vacancy.Name, response.Vacancy.Salary.From, response.Vacancy.Salary.To,
			response.Vacancy.Salary.Currency, response.Vacancy.Area, response.Vacancy.Employer, response.Vacancy.AlternateUrl, response.State)
	} else if len(response.Vacancy.Salary.From) > 0 {
		return fmt.Sprintf(messageTemplate, response.Vacancy.Name, response.Vacancy.Salary.From,
			response.Vacancy.Salary.Currency, response.Vacancy.Area, response.Vacancy.Employer, response.Vacancy.AlternateUrl, response.State)
	} else if len(response.Vacancy.Salary.To) > 0 {
		return fmt.Sprintf(messageTemplate, response.Vacancy.Name, response.Vacancy.Salary.To,
			response.Vacancy.Salary.Currency, response.Vacancy.Area, response.Vacancy.Employer, response.Vacancy.AlternateUrl, response.State)
	} else {
		return fmt.Sprintf(messageTemplate, response.Vacancy.Name, "Не указана",
			response.Vacancy.Salary.Currency, response.Vacancy.Area, response.Vacancy.Employer, response.Vacancy.AlternateUrl, response.State)
	}
}

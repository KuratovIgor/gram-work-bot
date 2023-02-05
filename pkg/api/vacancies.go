package api

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/types"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
)

const baseURL = "https://api.hh.ru"
const vacanciesURL = "/vacancies"
const areaURL = "/areas"

type VacanciesApiType struct {
	baseURL      string
	vacanciesURL string
	areaURL      string
}

func newVacanciesApi() *VacanciesApiType {
	return &VacanciesApiType{baseURL: baseURL, vacanciesURL: vacanciesURL, areaURL: areaURL}
}

func (v *VacanciesApiType) GetVacancies(params *Params) (types.Vacancies, error) {
	log.Println(baseURL + vacanciesURL + params.GetQueryString())
	res, _ := http.Get(baseURL + vacanciesURL + params.GetQueryString())
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var vacancies types.Vacancies

	value := gjson.Get(string(body), "items")
	for _, item := range value.Array() {
		var vacancy types.Vacancy
		vacancy.Id = gjson.Get(item.String(), "id").String()
		vacancy.Name = gjson.Get(item.String(), "name").String()
		vacancy.Salary.From = gjson.Get(item.String(), "salary.from").String()
		vacancy.Salary.To = gjson.Get(item.String(), "salary.to").String()
		vacancy.Salary.Currency = gjson.Get(item.String(), "salary.currency").String()
		vacancy.Address.City = gjson.Get(item.String(), "address.city").String()
		vacancy.Address.Street = gjson.Get(item.String(), "address.street").String()
		vacancy.Address.Building = gjson.Get(item.String(), "address.building").String()
		vacancy.PublishedAt = gjson.Get(item.String(), "published_at").String()
		vacancy.Employer = gjson.Get(item.String(), "employer.name").String()
		vacancy.Requirement = gjson.Get(item.String(), "snippet.requirement").String()
		vacancy.Responsibility = gjson.Get(item.String(), "snippet.responsibility").String()
		vacancy.Schedule = gjson.Get(item.String(), "schedule.name").String()
		vacancy.AlternateUrl = gjson.Get(item.String(), "alternate_url").String()
		vacancy.Area = gjson.Get(item.String(), "area.name").String()

		vacancies.Items = append(vacancies.Items, vacancy)
	}

	return vacancies, nil
}

var VacanciesApi = newVacanciesApi()

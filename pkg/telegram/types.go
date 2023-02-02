package telegram

type Vacancies struct {
	perPage int `json:"per_page"`
	items   []Vacancy
}

type Vacancy struct {
	salary           Salary
	name             string
	insiderInterview InsiderInterview `json:"insider_interview"`
	area             Area
	url              string
	publishedAt      string `json:"published_at"`
	relations        []string
	employer         Employer
}

type Salary struct {
	to       string
	from     string
	currency int
	gross    bool
}

type InsiderInterview struct {
	id  string
	url string
}

type Area struct {
	url  string
	id   string
	name string
}

type Employer struct {
	logoUrls     LogoUrls `json:"logo_urls"`
	name         string
	url          string
	alternateUrl string `json:"alternate_url"`
	id           string
	trusted      bool
}

type LogoUrls struct {
}
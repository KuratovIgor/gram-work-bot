package telegram

type Vacancies struct {
	items []Vacancy
	pages int
	page  int
}

type Vacancy struct {
	id             string
	name           string
	salary         Salary
	address        Address
	publishedAt    string
	employer       string
	requirement    string
	responsibility string
	schedule       string
	alternateUrl   string
}

type Address struct {
	city     string
	street   string
	building string
}

type Salary struct {
	from     string
	to       string
	currency string
}

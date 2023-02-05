package types

type Vacancies struct {
	Items []Vacancy
	Pages int
	Page  int
}

type Vacancy struct {
	Id             string
	Name           string
	Salary         Salary
	Address        Address
	PublishedAt    string
	Employer       string
	Requirement    string
	Responsibility string
	Schedule       string
	AlternateUrl   string
	Area           string
}

type Address struct {
	City     string
	Street   string
	Building string
}

type Salary struct {
	From     string
	To       string
	Currency string
}

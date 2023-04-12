package repository

type DefaultFilterType struct {
	Search     string `json:"search"`
	Salary     int    `json:"salary"`
	Schedule   string `json:"schedule"`
	Experience string `json:"experience"`
	AreaId     string `json:"area_id"`
}

type GraphqlRepository interface {
	GetAccessToken(chatId int64) (string, error)
	GetSessions() ([]int64, error)
	CreateSession(chatId int64, accessToken string, refreshToken string, userId string) error
	RemoveSession(chatId int64) error
	CreateUser(chatId int64, name string, lastname string, middlename string, email string, phone string, userId string) error
	SaveApplyToJob(userId string, vacancyId string, vacancyName string, employer string, alternateUrl string, area string, status string, salaryFrom int, salaryTo int, date string) error
	UpdateResponseStatus(vacancyId string, userId string, status string) error
	CreateDefaultFilters(chatId int64) error
	GetDefaultFilters(chatId int64) (*DefaultFilterType, error)
}

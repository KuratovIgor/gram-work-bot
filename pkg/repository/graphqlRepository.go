package repository

type GraphqlRepository interface {
	GetAccessToken(chatId int64) (string, error)
	CreateSession(chatId int64, accessToken string, refreshToken string) error
	RemoveSession(chatId int64) error
	CreateUser(chatId int64, name string, lastname string, middlename string, email string, phone string, userId string) error
	SaveApplyToJob(userId string, vacancyId string, vacancyName string, employer string, alternateUrl string, area string, status string, salaryFrom int, salaryTo int, date string) error
	UpdateResponseStatus(vacancyId string, userId string, status string) error
}

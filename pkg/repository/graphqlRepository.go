package repository

type GraphqlRepository interface {
	GetAccessToken(chatId int64) (string, error)
	CreateSession(chatId int64, accessToken string, refreshToken string) error
	RemoveSession(chatId int64) error
}
